package threads

import (
	"context"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"

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
		service.ListenAddress = "0.0.0.0"
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
			containerThread := NewContainerThread(service, conn, ctx)
			go handleConnection(containerThread, ctx, wg, db)
			defer containerThread.RemoveAllContainers()
		case err := <-errChan:
			log.Printf("error: failed to accept connection: %v", err)
			return
		}
	}
}

func handleConnection(containerThread Container, ctx context.Context, wg *sync.WaitGroup, db *database.Database) {
	containerThread.LaunchContainer()
	containerThread.Connect()
	containerThread.Communicate(wg, db)
	<-ctx.Done()
}
