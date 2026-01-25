package listener

import (
	"context"
	"crypto/tls"
	"fmt"
	"log/slog"
	"net"
	"os"
	"strings"
	"time"

	"github.com/hpotter/pkg/container"
	"github.com/hpotter/pkg/parser"
)

const TIMEOUT = 5 * time.Second

// Listen starts a new net.TCPListener or net.UDPListener based on
// the ListenProto specified by svc.
func (l *Listener) Listen(svc parser.Service) {
	slog.Info("starting listener", "service", fmt.Sprintf("%+v", svc))

	if strings.ToLower(svc.ListenProto) == "tcp" {
		l.listenTCP(svc)
	}
}

// listenTCP starts a new net.TCPListener given a TCP service specified
// by svc. listenTCP defers calling Done() on l.wg which only occurs if
// a cancellation signal is sent through l.ctx.
func (l *Listener) listenTCP(svc parser.Service) {
	defer l.Wg.Done()

	var err error
	var tlsConfig *tls.Config
	if svc.GenerateCerts {
		tlsConfig, err = l.Terminator.GenerateTlsKeyPair()
		if err != nil {
			slog.Error("failed to generate TLS config", "error", err)
			os.Exit(1)
		}
	}

	errChan := make(chan error)
	ln := l.createTCPListener(svc)
	for {
		select {
		case err := <-errChan:
			if slog.Default().Enabled(context.Background(), slog.LevelDebug) {
				slog.Error("container failure", "error", err)
			}
			return
		case <-l.Ctx.Done():
			l.ContainerThread.DockerClient.HandleCancel()
			l.handleShutdown()

			slog.Info("shutting down the TCP listener...")
			err := ln.Close()
			if err != nil {
				slog.Error("failed to close listener", "error", err)
			}
			return
		default:
			err := ln.SetDeadline(time.Now().Add(TIMEOUT))
			if err != nil {
				slog.Error("failed to set deadline", "error", err)
			}

			conn, err := ln.Accept()

			var finalConn net.Conn = conn
			if err, ok := err.(*net.OpError); ok && err.Timeout() {
				continue
			} else if err != nil {
				slog.Error("error creating connection", "error", err)
			} else {
				slog.Debug("recieved connection", "conn", conn.RemoteAddr())

				if svc.GenerateCerts {
					tlsConn := tls.Server(conn, tlsConfig)
					tlsErr := tlsConn.Handshake()
					if tlsErr != nil {
						if slog.Default().Enabled(context.Background(),
							slog.LevelDebug) {
							slog.Error("tls handshake failed", "error", tlsErr)
						}
						conn.Close()
						continue
					}

					finalConn = tlsConn
				}

				l.Connections = append(l.Connections, finalConn)
				slog.Debug("total connections", "size", len(l.Connections))
			}

			containerJob := container.ContainerJob{
				Svc:     svc,
				Conn:    finalConn,
				Started: time.Now(),
				ErrChan: errChan,
				Ctx:     l.Ctx,
				Db:      l.Db,
			}
			go l.ContainerThread.Start(&containerJob)
		}
	}
}

// createTCPListener creates and returns a new net.TCPListener using the
// address and port of svc.
func (l *Listener) createTCPListener(svc parser.Service) *net.TCPListener {
	a := fmt.Sprintf("%s:%s", svc.ListenAddress, svc.ListenPort)
	listenAddr, err := net.ResolveTCPAddr("tcp", a)
	if err != nil {
		slog.Error("failed to create TCPAddr", "error", err)
		os.Exit(1)
	}

	ln, err := net.ListenTCP("tcp", listenAddr)
	if err != nil {
		slog.Error("failed to start ListenTCP", "error", err)
		os.Exit(1)
	}

	return ln
}

// handleTCPShutdown closes all of the connections associated with l
// and closes the top level listener.
func (l *Listener) handleShutdown() {
	slog.Info("shutting down all open connections...")
	slog.Debug("total connections to shutdown", "size", len(l.Connections))
	for _, c := range l.Connections {
		c.Close()
	}
}
