package threads

import (
	"context"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
	"time"

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

			containerThread := NewContainerThread(service, conn, ctx)
			go handleConnection(containerThread, wg, db, dbConn)
			defer containerThread.RemoveAllContainers()
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
	var srcAddr string
	var srcPort int
	var destAddr string
	var destPort int
	var proto int
	switch addr := conn.LocalAddr().(type) {
	case *net.UDPAddr:
		srcAddr = addr.IP.String()
		srcPort = addr.Port
		remoteAddr := conn.RemoteAddr().(*net.UDPAddr)
		destAddr = remoteAddr.IP.String()
		destPort = remoteAddr.Port
		proto = 6
	case *net.TCPAddr:
		srcAddr = addr.IP.String()
		srcPort = addr.Port
		remoteAddr := conn.RemoteAddr().(*net.TCPAddr)
		destAddr = remoteAddr.IP.String()
		destPort = remoteAddr.Port
		proto = 17
	}
	dbConn := &database.Connections{
		CreatedAt:          time.Now().UTC(),
		SourceAddress:      srcAddr,
		SourcePort:         srcPort,
		DestinationAddress: destAddr,
		DestinationPort:    destPort,
		Container:          imageName,
		Proto:              proto,
	}
	return dbConn
}

func writeConnection(dbConn *database.Connections, db *database.Database) {
	if err := db.Write(dbConn); err != nil {
		log.Fatalf("error writing connection to database: %+v: %v", dbConn, err)
	}
}
