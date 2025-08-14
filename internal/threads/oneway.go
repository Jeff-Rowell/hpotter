package threads

import (
	"log"
	"net"
	"sync"
	"time"
)

type OneWayThread struct {
	Direction   string
	Source      net.Conn
	Destination net.Conn
	Container   Container
}

func NewOneWayThread(direction string, container *Container) OneWayThread {
	if direction == "request" {
		return OneWayThread{
			Direction:   direction,
			Source:      container.Source,
			Destination: container.Destination,
			Container:   *container,
		}
	} else {
		return OneWayThread{
			Direction:   direction,
			Source:      container.Destination,
			Destination: container.Source,
			Container:   *container,
		}
	}
}

func (oneway *OneWayThread) StartOneWayThread(wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-oneway.Container.Ctx.Done():
			return
		default:
		}
		oneway.Source.SetReadDeadline(time.Now().Add(100 * time.Millisecond))

		bytes := make([]byte, 4096)
		numBytesRead, err := oneway.Source.Read(bytes)
		if err != nil {
			if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
				continue
			}
			return
		}
		if numBytesRead == 0 {
			break
		}
		_, err = oneway.Destination.Write(bytes[:numBytesRead])
		if err != nil {
			log.Printf("error writing bytes in %s thread to container %s: %v", oneway.Direction, oneway.Container.CreateResponse.ID, err)
			return
		}
	}
}
