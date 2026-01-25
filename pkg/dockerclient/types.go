package dockerclient

import (
	"context"

	"github.com/moby/moby/client"
)

// Docker is the docker client for running honey pot containers. It has a
// [Client] that is used to check if honey pot images exist, pull images,
// create containers, start containers, and remove containers.
type Docker struct {
	// Client is the reference to the Docker API client
	Client *client.Client

	// ctx is the context to pass to the Docker API client and is different
	// than the main honey pot context
	Ctx context.Context

	// messages is a communication channel that follows the actor pattern
	// to control concurrent access to the global containers slice that keeps
	// track of running container ids
	Messages chan ActorMessage
}
