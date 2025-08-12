package threads

import (
	"context"
	"log"
	"net"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/Jeff-Rowell/hpotter/types"
	"github.com/docker/go-connections/nat"
	"github.com/moby/moby/api/types/container"
	"github.com/moby/moby/client"
)

type Container struct {
	Container      container.CreateResponse
	ContainerIP    string
	ContainerProto string
	Destination    net.Conn
	Svc            types.Service
	Ctx            context.Context
	DockerClient   *client.Client
	RequestThread  struct{} // TODO
	ResponseThread struct{} // TODO
}

func NewContainerThread(service types.Service) Container {
	ctx := context.Background()
	dockerClient, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		log.Fatalf("error creating docker client: %v", err)
	}
	return Container{
		Ctx:          ctx,
		DockerClient: dockerClient,
		Svc:          service,
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
	c.Container = createdContainer

	err = c.DockerClient.ContainerStart(c.Ctx, c.Container.ID, container.StartOptions{})
	if err != nil {
		log.Fatalf("error: failed to start container %s running image %s: %v", c.Container.ID, c.Svc.ImageName, err)
	}
}

func (c *Container) Connect() {
	inspectResponse, err := c.DockerClient.ContainerInspect(c.Ctx, c.Container.ID)
	if err != nil {
		log.Fatalf("error inspecting container %s running image %s", c.Container.ID, c.Svc.ImageName)
	}
	c.ContainerIP = inspectResponse.NetworkSettings.Networks["bridge"].IPAddress

	if strings.EqualFold(c.Svc.ListenProto, "tcp") {
		dest, err := net.DialTCP("tcp", nil, &net.TCPAddr{IP: net.ParseIP(c.ContainerIP), Port: c.Svc.ListenPort})
		if err != nil {
			log.Fatalf("network error: %v", err)
		}
		c.Destination = dest
	} else {
		dest, err := net.DialUDP("udp", nil, &net.UDPAddr{IP: net.ParseIP(c.ContainerIP), Port: c.Svc.ListenPort})
		if err != nil {
			log.Fatalf("network error: %v", err)
		}
		c.Destination = dest
	}
}
