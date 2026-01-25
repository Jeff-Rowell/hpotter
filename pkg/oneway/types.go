package oneway

import (
	"context"
	"net"
	"sync"
)

// OnewayJob contains the data required to start a one-way communication
// thread with a honeypot container.
type OnewayJob struct {
	// Source is the local originating connection
	Source net.Conn

	// Destination is the remote desination connection
	Destination net.Conn

	// Direction indicates which direction the one-way job is for
	// (i.e, "request" or "response")
	Direction string

	// Wg is the sync.WaitGroup used to keep the one-way threads alive until
	// the complete or are cancelled
	Wg *sync.WaitGroup

	// Ctx is the context.Context used to listen for cancellation signals
	Ctx context.Context

	// ErrChan is the channel used to communicate errors upstream with the
	// container thread so that the container can be gracefully removed
	// in the case of a failure.
	ErrChan chan error
}
