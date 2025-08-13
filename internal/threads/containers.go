package threads

import (
	"context"
	"fmt"
	"log"
	"net"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/Jeff-Rowell/hpotter/types"
	"github.com/docker/go-connections/nat"
	"github.com/moby/moby/api/types/container"
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
}

func NewContainerThread(service types.Service, source net.Conn) Container {
	ctx := context.Background()
	dockerClient, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		log.Fatalf("error creating docker client: %v", err)
	}
	return Container{
		Ctx:          ctx,
		DockerClient: dockerClient,
		Svc:          service,
		Source:       source,
	}
}

func (c *Container) LaunchContainer() {
	log.Printf("starting container: %s", filepath.Clean(c.Svc.ImageName))
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

	createdContainer, err := c.DockerClient.ContainerCreate(
		c.Ctx,
		&container.Config{
			Image: c.Svc.ImageName,
		},
		&container.HostConfig{
			PortBindings: portSet,
		},
		nil,
		nil,
		"")
	if err != nil {
		log.Fatalf("error running image '%s': %v", c.Svc.ImageName, err)
	}

	log.Printf("create container response: %+v", createdContainer)
	c.CreateResponse = createdContainer

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

	errSlice := make([]string, 10)
	if strings.EqualFold(c.Svc.ListenProto, "tcp") {
		for range 10 {
			dest, err := net.DialTCP("tcp", nil, &net.TCPAddr{IP: net.ParseIP(c.ContainerIP), Port: c.Svc.ListenPort})
			if err != nil {
				errSlice = append(errSlice, fmt.Sprintf("network error: %v\n", err))
				time.Sleep(2 * time.Second)
			} else {
				c.Destination = dest
				break
			}
		}
	} else {
		for range 10 {
			dest, err := net.DialUDP("udp", nil, &net.UDPAddr{IP: net.ParseIP(c.ContainerIP), Port: c.Svc.ListenPort})
			if err != nil {
				errSlice = append(errSlice, fmt.Sprintf("network error: %v\n", err))
				time.Sleep(2 * time.Second)
			} else {
				c.Destination = dest
				break
			}
		}
	}
	if len(errSlice) == 10 {
		log.Fatalf("error attempting connection 10 times: %v\n", errSlice)
	}
	log.Printf("successfully connected to container %s running image %s on %s", c.CreateResponse.ID, c.Svc.ImageName, c.ContainerIP)
}

func (c *Container) Communicate(wg *sync.WaitGroup) {
	wg.Add(1)
	requestThread := NewOneWayThread("request", c)
	go requestThread.StartOneWayThread(wg)

	wg.Add(1)
	responseThread := NewOneWayThread("response", c)
	go responseThread.StartOneWayThread(wg)
}
