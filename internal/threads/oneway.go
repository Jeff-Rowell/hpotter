package threads

import (
	"log"
	"net"
	"strings"
	"sync"
	"time"

	"github.com/Jeff-Rowell/hpotter/internal/database"
)

type OneWayThread struct {
	Direction   string
	Source      net.Conn
	Destination net.Conn
	Container   *Container
	Database    *database.Database
	DBConn      database.Connections
}

func NewOneWayThread(direction string, container *Container, db *database.Database, dbConn database.Connections) OneWayThread {
	if direction == "request" {
		return OneWayThread{
			Direction:   direction,
			Source:      container.Source,
			Destination: container.Destination,
			Container:   container,
			Database:    db,
			DBConn:      dbConn,
		}
	} else {
		return OneWayThread{
			Direction:   direction,
			Source:      container.Destination,
			Destination: container.Source,
			Container:   container,
			Database:    db,
			DBConn:      dbConn,
		}
	}
}

func (oneway *OneWayThread) StartOneWayThread(wg *sync.WaitGroup) {
	defer wg.Done()
	var totalData []byte
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
		totalData = append(totalData, bytes[:numBytesRead]...)
	}
	record := database.Data{
		Direction:     oneway.Direction,
		Data:          string(totalData),
		ConnectionsID: oneway.DBConn.ID,
	}
	if strings.ToLower(oneway.Direction) == "request" && oneway.Container.Svc.RequestSave {
		log.Printf("Request data: %s", string(totalData))
		if err := oneway.Database.Write(record); err != nil {
			log.Fatalf("error writing record to database: %+v: %v", record, err)
		}
	}
	if strings.ToLower(oneway.Direction) == "response" && oneway.Container.Svc.ResponseSave {
		log.Printf("Response data: %s", string(totalData))
		if err := oneway.Database.Write(record); err != nil {
			log.Fatalf("error writing record to database: %+v: %v", record, err)
		}
	}
}
