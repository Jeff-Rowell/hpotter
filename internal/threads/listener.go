package threads

import (
	"fmt"
	"log"
	"net"
	"strings"
	"sync"

	"github.com/Jeff-Rowell/hpotter/types"
)

func StartListener(service types.Service, wg *sync.WaitGroup) {
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
		log.Fatalf("error creating listener on %s %d/%s", service.ListenAddress, service.ListenPort, service.ListenProto)
	}
	defer listenSocket.Close()
	log.Printf("created socket listener on %s", service.ListenAddress)

	if err != nil {
		log.Fatalf("error: failed to create listener: %v", err)
	}

	for {
		conn, err := listenSocket.Accept()
		if err != nil {
			log.Fatalf("error: failed to accept connection: %v", err)
		}
		log.Printf("connection received: (src=%s, dst=%s, proto=%s)", conn.RemoteAddr(), conn.LocalAddr(), conn.LocalAddr().Network())
		containerThread := NewContainerThread(service, conn)
		containerThread.LaunchContainer()
		containerThread.Connect()
		containerThread.Communicate(wg)
	}
}
