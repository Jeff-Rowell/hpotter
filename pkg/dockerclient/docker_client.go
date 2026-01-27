package dockerclient

import (
	"context"
	"fmt"
	"log/slog"
	"net/netip"
	"slices"
	"strconv"

	"github.com/moby/moby/api/types/container"
	"github.com/moby/moby/api/types/mount"
	"github.com/moby/moby/api/types/network"
	"github.com/moby/moby/client"
)

const (
	DB_NAME    = "hpotter-db"
	DB_NETWORK = "hpotter-net"
	DB_VOLUME  = "hpotter-data"
)

// NewDockerClient creates a new Docker API client from the environment
// and returns the initialized Docker API client or an error if the client
// was not able to be initialized from the environment
func NewDockerClient() (*Docker, error) {
	dockerClient, err := client.New(client.FromEnv)
	if err != nil {
		return nil, err
	}

	d := &Docker{
		Client:   dockerClient,
		Ctx:      context.Background(),
		Messages: make(chan ActorMessage),
	}

	go manageContainers(d.Messages)

	return d, nil
}

// CreateDbVolume creates a docker volume DB_VOLUME and returns an error
// if any is encountered during creation.
func (d *Docker) CreateDbVolume() error {
	exists, err := d.checkVolume()
	if err != nil {
		return err
	}

	if exists {
		return nil
	}

	vOpts := client.VolumeCreateOptions{Name: DB_VOLUME}
	_, err = d.Client.VolumeCreate(d.Ctx, vOpts)
	if err != nil {
		return err
	}
	return nil
}

// GetDbNetworkId calls the Docker API to search for docker networks with
// a name matching DB_NETWORK. The network id is returned if found along
// with a nil error. If no networks are found, the empty string is returned
// with a nil error.
func (d *Docker) GetDbNetworkId() (string, error) {
	filters := make(client.Filters).Add("name", DB_NETWORK)
	listOpts := client.NetworkListOptions{Filters: filters}

	networks, err := d.Client.NetworkList(d.Ctx, listOpts)
	if err != nil {
		return "", err
	}

	if len(networks.Items) == 0 {
		return "", nil
	}

	return networks.Items[0].ID, nil
}

// CreateDbNetwork creates a network DB_NETWORK and returns the network id
// and a nil error if successful. If an error is encountered, the empty
// string is returned along with an error.
func (d *Docker) CreateDbNetwork() (string, error) {
	netOpts := client.NetworkCreateOptions{Driver: "bridge"}

	network, err := d.Client.NetworkCreate(d.Ctx, DB_NETWORK, netOpts)
	if err != nil {
		return "", err
	}

	return network.ID, nil
}

// CreateDbContainer creates a container specifically for running the
// database. CreateDbContainer returns the container id and a nil error
// if the operation is successful. Otherwise the empty string and the error
// are returned.
func (d *Docker) CreateDbContainer(u, p, i string, pt int) (string, error) {
	hostConfig, err := d.getDbHostConfig(pt)
	if err != nil {
		return "", err
	}

	containerConfig, err := d.getDbContainerConfig(u, p, i, pt)
	if err != nil {
		return "", err
	}

	containerOpts := client.ContainerCreateOptions{
		Config:           containerConfig,
		HostConfig:       hostConfig,
		NetworkingConfig: d.getDbNetworkConfig(),
		Name:             DB_NAME,
	}

	container, err := d.Client.ContainerCreate(d.Ctx, containerOpts)
	if err != nil {
		return "", err
	}

	return container.ID, nil
}

// StartDbContainer starts the given container id and returns an error.
func (d *Docker) StartDbContainer(containerId string) error {
	return d.StartContainer(containerId)
}

// getDbHostConfig builds and returns the host configuration for the
// database container. The host configuration is returns with a nil error
// if successful, otherwise nil is returned with a non-nil error.
func (d *Docker) getDbHostConfig(pt int) (*container.HostConfig, error) {
	listenOn, err := netip.ParseAddr("127.0.0.1")
	if err != nil {
		return nil, err
	}

	port, err := network.ParsePort(fmt.Sprintf("%d/tcp", pt))
	if err != nil {
		return nil, err
	}

	portBinding := network.PortBinding{
		HostIP:   listenOn,
		HostPort: strconv.Itoa(pt),
	}

	portMap := make(map[network.Port][]network.PortBinding)
	portMap[port] = []network.PortBinding{portBinding}

	hostConfig := container.HostConfig{
		NetworkMode: "bridge",
		Privileged:  false,
		Mounts: []mount.Mount{
			{
				Type:   mount.TypeVolume,
				Source: DB_VOLUME,
				Target: "/var/lib/postgresql/data",
			},
		},
		RestartPolicy: container.RestartPolicy{Name: "unless-stopped"},
	}

	return &hostConfig, nil
}

// getDbContainerConfig builds and returns the container configuration
// for the database container. If an error is encountered, the config returned
// will be nil along with a non-nil error.
func (d *Docker) getDbContainerConfig(u, p, i string, pt int) (*container.Config, error) {
	port, err := network.ParsePort(fmt.Sprintf("%d/tcp", pt))
	if err != nil {
		return nil, err
	}
	portSet := make(map[network.Port]struct{}, 1)
	portSet[port] = struct{}{}

	conf := container.Config{
		Hostname:     "diagon-alley",
		Domainname:   "hogwarts",
		Image:        i,
		ExposedPorts: portSet,
		Env: []string{
			fmt.Sprintf("POSTGRES_DB=%s", DB_NAME),
			fmt.Sprintf("POSTGRES_USER=%s", u),
			fmt.Sprintf("POSTGRES_PASSWORD=%s", p),
		},
	}
	return &conf, nil
}

// getDbNetworkConfig builds and returns the network configuration for the
// database container.
func (d *Docker) getDbNetworkConfig() *network.NetworkingConfig {
	endpointConf := make(map[string]*network.EndpointSettings)
	endpointConf[DB_NETWORK] = &network.EndpointSettings{}
	netConf := network.NetworkingConfig{
		EndpointsConfig: endpointConf,
	}
	return &netConf
}

// GetDbName returns the DB_NAME for external reference.
func (d *Docker) GetDbName() string {
	return DB_NAME
}

// checkVolume checks if the volume DB_VOLUME exists and returns a bool
// indicating the existense along with an error if any.
func (d *Docker) checkVolume() (bool, error) {
	filters := make(client.Filters).Add("name", DB_VOLUME)
	opts := client.VolumeListOptions{Filters: filters}

	volumes, err := d.Client.VolumeList(d.Ctx, opts)
	if err != nil {
		return false, err
	}

	return len(volumes.Items) > 0, nil
}

// StartContainer starts a container specified by containerId which should
// be obtained from calling [CreateContainer]. StartContainer will return
// nil if the container was successfully started, otherwise the error
// will be returned.
func (d *Docker) StartContainer(containerId string) error {
	startOpts := client.ContainerStartOptions{}
	_, err := d.Client.ContainerStart(d.Ctx, containerId, startOpts)
	if err != nil {
		return err
	}
	return nil
}

// GetDBContainerIP calls the Docker API to inspect the container specified
// by id, and returns the IPv4 address of the bridge network. If an error
// is encounted "" is returned instead of the IPv4 address, otherwise the
// error will be nil.
func (d *Docker) GetDBContainerIP(id string) (string, error) {
	inspectOpts := client.ContainerInspectOptions{}
	inspectResult, err := d.Client.ContainerInspect(d.Ctx, id, inspectOpts)
	if err != nil {
		return "", err
	}

	ip := inspectResult.Container.NetworkSettings.Networks[DB_NETWORK].IPAddress
	return ip.String(), nil
}

// ContainerStop stops the container specified by id with default Docker API
// timeout settings. If the remove flag is passed in, the container will also
// be removed. If an error is encountered, the error is returned, otherwise
// ContainerStop will return nil.
func (d *Docker) ContainerStop(id string, remove bool) error {
	_, err := d.Client.ContainerStop(d.Ctx, id, client.ContainerStopOptions{})
	if err != nil {
		return err
	}

	if remove {
		opts := client.ContainerRemoveOptions{
			Force:         true,
			RemoveVolumes: false,
			RemoveLinks:   false,
		}
		_, err := d.Client.ContainerRemove(d.Ctx, id, opts)
		if err != nil {
			return err
		}
	}
	return nil
}

// GetContainerIP calls the Docker API to inspect the container specified
// by id, and returns the IPv4 address of the bridge network. If an error
// is encounted "" is returned instead of the IPv4 address, otherwise the
// error will be nil.
func (d *Docker) GetContainerIP(id string) (string, error) {
	inspectOpts := client.ContainerInspectOptions{}
	inspectResult, err := d.Client.ContainerInspect(d.Ctx, id, inspectOpts)
	if err != nil {
		return "", err
	}

	ip := inspectResult.Container.NetworkSettings.Networks["bridge"].IPAddress
	return ip.String(), nil
}

// shutdown stops and removes the container associated with the given
// container id, id.
func (d *Docker) Shutdown(id string) error {
	slog.Debug("shutting down container", "container", id)
	err := d.ContainerStop(id, false)

	if err != nil {
		return err
	}

	d.Messages <- ActorMessage{Type: Delete, Value: id}

	return nil
}

// HasImage calls the Docker API to list all images on the host and returns
// true if i is already present on the host, otherwise false. If an
// error is encountered with the Docker API, the error is returned with
// false indicating the failure.
func (d *Docker) HasImage(i string) (bool, error) {
	listOpts := client.ImageListOptions{All: true}
	images, err := d.Client.ImageList(d.Ctx, listOpts)
	if err != nil {
		return false, err
	}

	for _, img := range images.Items {
		if slices.Contains(img.RepoTags, i) {
			return true, nil
		}
	}

	return false, nil
}

// PullImage attempts to download image i by calling the Docker API. If an
// error is encountered while attempting to pull the image, the error is
// returned, otherwise nil is returned.
func (d *Docker) PullImage(i string) error {
	pullOpts := client.ImagePullOptions{}
	imageReadCloser, err := d.Client.ImagePull(d.Ctx, i, pullOpts)
	if err != nil {
		return err
	}

	err = imageReadCloser.Wait(d.Ctx)
	if err != nil {
		return err
	}

	imageReadCloser.Close()

	return nil
}

// HandleCancel attempts to stop all running honey pot containers and provides
// an external handle into cleaning up all running container threads.
// HandleCancel only needs to call [ContainerStop] because the containers
// are created with AutoRemove set to true.
func (d *Docker) HandleCancel() {
	slog.Info("received cancellation signal...")
	resultChan := make(chan []string)
	d.Messages <- ActorMessage{Type: Get, Response: resultChan}
	runningContainers := <-resultChan
	for _, cId := range runningContainers {
		slog.Info("removing container", "container id", cId)
		err := d.ContainerStop(cId, false)
		if err != nil {
			slog.Error("failed to stop container", "error", err)
		}
		d.Messages <- ActorMessage{Type: Delete, Value: cId}
	}
}
