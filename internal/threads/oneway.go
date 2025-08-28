package threads

import (
	"fmt"
	"io"
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
	var consecutiveEOFs int

	defer func() {
		log.Printf("terminating oneway '%s' thread...", oneway.Direction)
		if len(totalData) > 0 {
			dataString := fmt.Sprintf("%q", totalData)
			record := &database.Data{
				Direction:     oneway.Direction,
				Data:          dataString,
				ConnectionsID: oneway.DBConn.ID,
			}
			if strings.ToLower(oneway.Direction) == "request" && oneway.Container.Svc.RequestSave {
				if err := oneway.Database.Write(record); err != nil {
					log.Printf("error writing 'request' record to database: %+v: %v", record, err)
				}
				log.Println("successfully wrote 'request' data to db")
			}
			if strings.ToLower(oneway.Direction) == "response" && oneway.Container.Svc.ResponseSave {
				if err := oneway.Database.Write(record); err != nil {
					log.Printf("error writing 'response' record to database: %+v: %v", record, err)
				}
				log.Println("successfully wrote 'response' data to db")
			}
		}
	}()
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
			if err == io.EOF {
				consecutiveEOFs++
				// If we get multiple consecutive EOFs, the connection is likely closed
				if consecutiveEOFs >= 2 {
					return
				}
				continue
			}
			return
		}
		if numBytesRead == 0 {
			break
		}
		consecutiveEOFs = 0
		totalData = append(totalData, bytes[:numBytesRead]...)

		oneway.Destination.SetWriteDeadline(time.Now().Add(1 * time.Second))
		_, err = oneway.Destination.Write(bytes[:numBytesRead])
		if err != nil {
			if isConnectionClosed(err) {
				return
			}
			break
		}
	}
}

func isConnectionClosed(err error) bool {
	if err == nil {
		return false
	}

	errStr := err.Error()
	return strings.Contains(errStr, "broken pipe") ||
		strings.Contains(errStr, "connection reset by peer") ||
		strings.Contains(errStr, "use of closed network connection") ||
		strings.Contains(errStr, "connection refused")
}
