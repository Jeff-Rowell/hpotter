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

	listenSocket, err := net.Listen(lowerProto, fmt.Sprintf("0.0.0.0:%d", service.ListenPort))
	if err != nil {
		log.Fatalf("error: failed to create listener: %v", err)
	}

	for {
		conn, err := listenSocket.Accept()
		fmt.Printf("service: %+v", service)
		if err != nil {
			log.Fatalf("error: failed to accept connection: %v", err)
		}
		log.Printf("connection received: (src=%s, dst=%s, proto=%s)", conn.RemoteAddr(), conn.LocalAddr(), conn.LocalAddr().Network())
		LaunchContainer(service)
	}
}
