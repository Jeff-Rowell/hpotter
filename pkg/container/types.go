package container

import (
	"context"
	"net"
	"time"

	"github.com/hpotter/pkg/database"
	"github.com/hpotter/pkg/dockerclient"
	"github.com/hpotter/pkg/parser"
)

// ContainerJob is a job type for running a new honey pot container.
// It has a parser.Service that specifies the honey pot container and a
// net.Conn for communicating back to the incoming connection.
// It also has a write-only channel for communicating any errors back
// to the main context.
type ContainerJob struct {
	// Svc is the honey pot container service specification parsed from
	// the configuration file.
	Svc parser.Service

	// Conn is the accepted inbound connection on the host that needs
	// to be forwarded to the honey pot container for processing.
	Conn net.Conn

	// Started is the time at which the container job transaction started.
	Started time.Time

	// ErrChan is a write-only channel passed in from the main context
	// and is used to communicate container-based errors up to the main
	// context.
	ErrChan chan<- error

	// Ctx is the main context passed in from the main context and is used
	// to listen for cancellation signals and terminate in-flight requests.
	Ctx context.Context

	// Db is the handle to the database.
	Db *database.Database
}

// Docker is the docker client for running honey pot containers. It has a
// [Client] that is used to check if honey pot images exist, pull images,
// create containers, start containers, and remove containers.
type ContainerThread struct {
	// DockerClient is the Docker API client
	DockerClient *dockerclient.Docker
}
