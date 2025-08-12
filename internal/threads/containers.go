package threads

import (
	"context"
	"log"
	"path/filepath"

	"github.com/Jeff-Rowell/hpotter/types"
	"github.com/moby/moby/api/types/container"
	"github.com/moby/moby/client"
)

func LaunchContainer(service types.Service) {
	ctx := context.Background()

	apiClient, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		log.Fatalf("error creating docker client: %v", err)
	}
	defer apiClient.Close()

	log.Printf("starting container: %s", filepath.Clean(service.ImageName))
	createdContainer, err := apiClient.ContainerCreate(ctx,
		&container.Config{
			Image: service.ImageName,
		},
		nil, nil, nil, "")
	if err != nil {
		log.Fatalf("error running image '%s': %v", service.ImageName, err)
	}

	log.Printf("create container response: %+v", createdContainer)
	defer func() {
		log.Printf("removing container %s running image %s", createdContainer.ID, service.ImageName)
		removeOps := container.RemoveOptions{
			Force: true,
		}
		if err := apiClient.ContainerRemove(ctx, createdContainer.ID, removeOps); err != nil {
			log.Fatalf("error removing container %s running image %s: %v", createdContainer.ID, service.ImageName, err)
		}
		log.Printf("successfully removed container %s running image %s", createdContainer.ID, service.ImageName)
	}()
}
