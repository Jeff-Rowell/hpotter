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
	"github.com/Jeff-Rowell/hpotter/internal/logparser"
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

		if strings.EqualFold(oneway.Direction, "request") {
			factory := logparser.NewLogParserFactory()
			if !factory.IsSupported(oneway.Container.Svc) {
				// Service not supported for log parsing, save totalData as string
				if len(totalData) > 0 && oneway.Container.Svc.RequestSave {
					record := &database.Data{
						Direction:     oneway.Direction,
						Data:          fmt.Sprintf("%q", totalData),
						ConnectionsID: oneway.DBConn.ID,
					}
					if err := oneway.Database.Write(record); err != nil {
						log.Fatalf("error writing request data to database: %+v: %v", record, err)
					}
					log.Println("successfully wrote request data to db")
				}
				return
			}

			logData, err := oneway.Container.ReadLogs()
			if err != nil {
				log.Fatalf("error reading container logs: %v", err)
			}

			parser, err := factory.CreateParser(oneway.Container.Svc)
			if err != nil {
				log.Fatalf("error creating log parser: %v", err)
			}

			cred := parser.ParseCredentials(logData)
			if cred != nil {
				cred.ConnectionsID = oneway.DBConn.ID
				if err := oneway.Database.Write(cred); err != nil {
					log.Fatalf("error writing credential to database: %+v: %v", cred, err)
				}
			}

			data := parser.ParseSessionData(logData)
			if data != nil {
				data.ConnectionsID = oneway.DBConn.ID
				if err := oneway.Database.Write(data); err != nil {
					log.Fatalf("error writing session data to database: %+v: %v", data, err)
				}
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
