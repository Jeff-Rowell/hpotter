package threads

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"strings"
	"sync"
	"time"

	retryablehttp "github.com/hashicorp/go-retryablehttp"

	"github.com/Jeff-Rowell/hpotter/internal/credentials"
	"github.com/Jeff-Rowell/hpotter/internal/database"
	"github.com/Jeff-Rowell/hpotter/types"
)

func StartListener(service types.Service, wg *sync.WaitGroup, ctx context.Context, db *database.Database) {
	defer wg.Done()
	lowerProto := strings.ToLower(service.ListenProto)
	log.Printf("starting listener on %s port %d", lowerProto, service.ListenPort)

	var err error
	var listenSocket net.Listener
	if service.ListenAddress == "" {
		service.ListenAddress = "127.0.0.1"
	}

	listenSocket, err = net.Listen(lowerProto, fmt.Sprintf("%s:%d", service.ListenAddress, service.ListenPort))
	if err != nil {
		log.Fatalf("error creating listener on %s %d/%s: %v", service.ListenAddress, service.ListenPort, service.ListenProto, err)
	}
	defer listenSocket.Close()
	log.Printf("created socket listener on %s", service.ListenAddress)

	if err != nil {
		log.Fatalf("error: failed to create listener: %v", err)
	}

	connChan := make(chan net.Conn)
	errChan := make(chan error)

	go func() {
		for {
			conn, err := listenSocket.Accept()
			if err != nil {
				errChan <- err
				return
			}
			connChan <- conn
		}
	}()

	for {
		select {
		case <-ctx.Done():
			log.Printf("listener: received cancellation signal")
			return
		case conn := <-connChan:
			log.Printf("connection received: (src=%s, dst=%s, proto=%s)", conn.RemoteAddr(), conn.LocalAddr(), conn.LocalAddr().Network())

			dbConn := buildConnection(service.ImageName, conn)
			go writeConnection(dbConn, db)

			if service.CollectCredentials {
				go handleCredentialCollection(service, conn, ctx, wg, db, dbConn)
			} else {
				containerThread := NewContainerThread(service, conn, ctx)
				go handleConnection(containerThread, wg, db, dbConn)
			}
		case err := <-errChan:
			log.Printf("error: failed to accept connection: %v", err)
			return
		}
	}
}

func handleConnection(containerThread Container, wg *sync.WaitGroup, db *database.Database, dbConn *database.Connections) {
	containerThread.LaunchContainer()
	containerThread.Connect()
	containerThread.Communicate(wg, db, dbConn)
	wg.Wait()
}

func buildConnection(imageName string, conn net.Conn) *database.Connections {
	var dbConn database.Connections

	switch addr := conn.LocalAddr().(type) {
	case *net.UDPAddr:
		remoteAddr := conn.RemoteAddr().(*net.UDPAddr)
		sourceLat, sourceLong, err := getLatitudeLongitude(remoteAddr.IP.String())
		if err != nil {
			log.Fatalf("error getting latitude and longitude: %v", err)
		}
		dbConn = database.Connections{
			CreatedAt:          time.Now().UTC(),
			SourceAddress:      remoteAddr.IP.String(),
			SourcePort:         remoteAddr.Port,
			DestinationAddress: addr.IP.String(),
			DestinationPort:    addr.Port,
			Latitude:           sourceLat,
			Longitude:          sourceLong,
			Container:          imageName,
			Proto:              6,
		}
	case *net.TCPAddr:
		remoteAddr := conn.RemoteAddr().(*net.TCPAddr)
		sourceLat, sourceLong, err := getLatitudeLongitude(remoteAddr.IP.String())
		if err != nil {
			log.Fatalf("error getting latitude and longitude: %v", err)
		}
		dbConn = database.Connections{
			CreatedAt:          time.Now().UTC(),
			SourceAddress:      remoteAddr.IP.String(),
			SourcePort:         remoteAddr.Port,
			DestinationAddress: addr.IP.String(),
			DestinationPort:    addr.Port,
			Latitude:           sourceLat,
			Longitude:          sourceLong,
			Container:          imageName,
			Proto:              17,
		}
	}
	return &dbConn
}

func getLatitudeLongitude(ipAddr string) (float32, float32, error) {
	client := retryablehttp.NewClient()
	client.RetryMax = 5
	client.RetryWaitMin = 1 * time.Second
	client.RetryWaitMax = 30 * time.Second
	client.Logger = nil

	httpClient := client.StandardClient()
	req, err := http.NewRequest("GET", fmt.Sprintf("http://ip-api.com/json/%s", ipAddr), nil)
	if err != nil {
		return -1, -1, err
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		return -1, -1, err
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return -1, -1, err
	}

	var results types.IPInfo
	err = json.Unmarshal(bodyBytes, &results)
	if err != nil {
		return -1, -1, err
	}
	return results.Latitude, results.Longitude, nil
}

func writeConnection(dbConn *database.Connections, db *database.Database) {
	if err := db.Write(dbConn); err != nil {
		log.Fatalf("error writing connection to database: %+v: %v", dbConn, err)
	}
}

func handleCredentialCollection(service types.Service, conn net.Conn, _ context.Context, _ *sync.WaitGroup, db *database.Database, dbConn *database.Connections) {
	defer conn.Close()

	log.Printf("starting credential collection for service on port %d", service.ListenPort)

	collector, err := credentials.NewCredentialCollector(service, conn, db, dbConn)
	if err != nil {
		log.Printf("error creating credential collector: %v", err)
		return
	}

	creds, err := collector.CollectCredentials()
	if err != nil {
		log.Printf("error collecting credentials: %v", err)
		return
	}

	log.Printf("successfully collected credentials: user=%s, connection_id=%d", creds.Username, creds.ConnectionsID)
	log.Printf("credential collection complete for connection from %s", conn.RemoteAddr())
}
