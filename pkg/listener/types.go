package listener

import (
	"context"
	"net"
	"sync"

	"github.com/hpotter/pkg/container"
	"github.com/hpotter/pkg/database"
	"github.com/hpotter/pkg/terminator"
)

// Listener is the honey pot listener and keeps track of open connections.
// The Listener is responsible for brokering new connections to the
// downstream Container routines. The Listener is also responsible
// for listening for cancellation signals and closing the listener and
// all connections
type Listener struct {
	// connections is the slice of all open connections
	Connections []net.Conn

	// wg is the sync.WaitGroup responsible for synchronizing the routines
	Wg *sync.WaitGroup

	// ctx is the main context.Context to listen for cancellation signals on
	Ctx context.Context

	// ContainerThread is responsible for spinning up containers
	ContainerThread *container.ContainerThread

	// Terminator is responsible for terminating TLS
	Terminator *terminator.Terminator

	// Db is the handle to our database
	Db *database.Database
}
