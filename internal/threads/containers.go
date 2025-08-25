package threads

import (
	"context"
	"fmt"
	"log"
	"net"
	"path/filepath"
	"strconv"
	"sync"
	"time"

	"github.com/Jeff-Rowell/hpotter/internal/database"
	"github.com/Jeff-Rowell/hpotter/types"
	"github.com/docker/go-connections/nat"
	"github.com/moby/moby/api/types/container"
	"github.com/moby/moby/api/types/filters"
	"github.com/moby/moby/client"
)

type Container struct {
	CreateResponse container.CreateResponse
	ContainerIP    string
	ContainerProto string
	Source         net.Conn
	Destination    net.Conn
	Svc            types.Service
	Ctx            context.Context
	DockerClient   *client.Client
	Labels         map[string]string
}

func NewContainerThread(service types.Service, source net.Conn, ctx context.Context) Container {
	dockerClient, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		log.Fatalf("error creating docker client: %v", err)
	}
	labels := map[string]string{
		"hpotter": "container",
	}
	return Container{
		Ctx:          ctx,
		DockerClient: dockerClient,
		Svc:          service,
		Source:       source,
		Labels:       labels,
	}
}

func (c *Container) LaunchContainer() {
	log.Printf("creating container: %s", filepath.Clean(c.Svc.ImageName))
	port, err := nat.NewPort(c.Svc.ListenProto, strconv.Itoa(c.Svc.ListenPort))
	if err != nil {
		log.Fatalf("error creating nat port %d/%s: %v", c.Svc.ListenPort, c.Svc.ListenProto, err)
	}

	var portBinding nat.PortBinding
	if c.Svc.ListenAddress == "" {
		portBinding = nat.PortBinding{
			HostIP: "0.0.0.0",
		}
	} else {
		portBinding = nat.PortBinding{
			HostIP: c.Svc.ListenAddress,
		}
	}
	portSet := container.PortMap{port: []nat.PortBinding{portBinding}}

	// Build environment variables slice from service config
	var envVars []string
	for _, envVar := range c.Svc.EnvVars {
		envVars = append(envVars, fmt.Sprintf("%s=%s", envVar.Key, envVar.Value))
	}

	createdContainer, err := c.DockerClient.ContainerCreate(
		c.Ctx,
		&container.Config{
			Image:  c.Svc.ImageName,
			Labels: c.Labels,
			Env:    envVars,
		},
		&container.HostConfig{
			PortBindings: portSet,
		},
		nil,
		nil,
		"")
	if err != nil {
		log.Fatalf("error creating container using image '%s': %v", c.Svc.ImageName, err)
	}

	log.Printf("create container response: %+v", createdContainer)
	c.CreateResponse = createdContainer

	log.Printf("starting container: %s", filepath.Clean(c.Svc.ImageName))
	err = c.DockerClient.ContainerStart(c.Ctx, c.CreateResponse.ID, container.StartOptions{})
	if err != nil {
		log.Fatalf("error: failed to start container %s running image %s: %v", c.CreateResponse.ID, c.Svc.ImageName, err)
	}
}

func (c *Container) Connect() {
	log.Printf("connecting to container %s running image %s", c.CreateResponse.ID, c.Svc.ImageName)
	inspectResponse, err := c.DockerClient.ContainerInspect(c.Ctx, c.CreateResponse.ID)
	if err != nil {
		log.Fatalf("error inspecting container %s running image %s", c.CreateResponse.ID, c.Svc.ImageName)
	}
	c.ContainerIP = inspectResponse.NetworkSettings.Networks["bridge"].IPAddress

	var errSlice []string
	for range 10 {
		select {
		case <-c.Ctx.Done():
			return
		default:
		}

		dest, err := net.DialTimeout(c.Svc.ListenProto, fmt.Sprintf("%s:%d", c.ContainerIP, c.Svc.ListenPort), 5*time.Second)
		if err != nil {
			errSlice = append(errSlice, fmt.Sprintf("network error: %v\n", err))
			select {
			case <-c.Ctx.Done():
				return
			case <-time.After(2 * time.Second):
			}
		} else {
			c.Destination = dest
			break
		}
	}
	if len(errSlice) == 10 {
		log.Fatalf("error attempting connection 10 times: %v\n", errSlice)
	}
	log.Printf("successfully connected to container %s running image %s on %s", c.CreateResponse.ID, c.Svc.ImageName, c.ContainerIP)
}

func (c *Container) Communicate(wg *sync.WaitGroup, db *database.Database, dbConn *database.Connections) {
	wg.Add(1)
	requestThread := NewOneWayThread("request", c, db, *dbConn)
	go requestThread.StartOneWayThread(wg)

	wg.Add(1)
	responseThread := NewOneWayThread("response", c, db, *dbConn)
	go responseThread.StartOneWayThread(wg)
}

func (c *Container) GetOurContainers() []container.Summary {
	listCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	filters := filters.NewArgs()
	filters.Add("label", "hpotter=container")
	containers, err := c.DockerClient.ContainerList(listCtx, container.ListOptions{Filters: filters})
	if err != nil {
		log.Fatalf("error listing containers by hpotter=container label: %v", err)
	}
	return containers
}

func (c *Container) RemoveContainer(containerID, imageName string) {
	cleanupCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	removeOps := container.RemoveOptions{
		Force: true,
	}
	if err := c.DockerClient.ContainerRemove(cleanupCtx, containerID, removeOps); err != nil {
		log.Printf("error removing container %s running image %s: %v", containerID, imageName, err)
		return
	}
}

func (c *Container) RemoveAllContainers() {
	ourContainers := c.GetOurContainers()
	for _, container := range ourContainers {
		c.RemoveContainer(container.ID, container.Image)
	}
}
