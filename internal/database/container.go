package database

import (
	"context"
	"fmt"
	"log"
	"slices"
	"strconv"

	"github.com/Jeff-Rowell/hpotter/types"
	"github.com/docker/go-connections/nat"
	"github.com/moby/moby/api/types/container"
	"github.com/moby/moby/api/types/filters"
	"github.com/moby/moby/api/types/mount"
	"github.com/moby/moby/api/types/network"
	"github.com/moby/moby/api/types/volume"
	"github.com/moby/moby/client"
)

const (
	DatabaseNetworkName   = "hpotter-db-network"
	DatabaseVolumeName    = "hpotter-db-data"
	DatabaseContainerName = "hpotter-database"
)

type DatabaseContainer struct {
	Client      *client.Client
	Config      types.DBConfig
	ContainerID string
	NetworkID   string
	VolumeName  string
}

func NewDatabaseContainer(config types.DBConfig) (*DatabaseContainer, error) {
	dockerClient, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return nil, fmt.Errorf("failed to create docker client: %w", err)
	}

	return &DatabaseContainer{
		Client: dockerClient,
		Config: config,
	}, nil
}

func (dc *DatabaseContainer) CheckExistingVolume(ctx context.Context) (bool, error) {
	volumes, err := dc.Client.VolumeList(ctx, volume.ListOptions{})
	if err != nil {
		return false, fmt.Errorf("failed to list volumes: %w", err)
	}

	for _, vol := range volumes.Volumes {
		if vol.Name == DatabaseVolumeName {
			dc.VolumeName = vol.Name
			log.Printf("found existing database volume: %s", vol.Name)
			return true, nil
		}
	}
	return false, nil
}

func (dc *DatabaseContainer) CreateVolume(ctx context.Context) error {
	volumeResponse, err := dc.Client.VolumeCreate(ctx, volume.CreateOptions{
		Name: DatabaseVolumeName,
		Labels: map[string]string{
			"hpotter": "database-volume",
		},
	})
	if err != nil {
		return fmt.Errorf("failed to create volume: %w", err)
	}

	dc.VolumeName = volumeResponse.Name
	log.Printf("created database volume: %s", volumeResponse.Name)
	return nil
}

func (dc *DatabaseContainer) CreateNetwork(ctx context.Context) error {
	networkResponse, err := dc.Client.NetworkCreate(ctx, DatabaseNetworkName, network.CreateOptions{
		Driver: "bridge",
		Labels: map[string]string{
			"hpotter": "database-network",
		},
		Internal: true, // Make it isolated from external networks
	})
	if err != nil {
		return fmt.Errorf("failed to create network: %w", err)
	}

	dc.NetworkID = networkResponse.ID
	log.Printf("created isolated database network: %s", DatabaseNetworkName)
	return nil
}

func (dc *DatabaseContainer) StartContainer(ctx context.Context) error {
	var envVars []string
	var volumeMount string

	switch dc.Config.DBType {
	case "postgres", "postgresql":
		envVars = []string{
			fmt.Sprintf("POSTGRES_DB=%s", DatabaseContainerName),
			fmt.Sprintf("POSTGRES_USER=%s", dc.Config.User),
			fmt.Sprintf("POSTGRES_PASSWORD=%s", dc.Config.Password),
		}
		volumeMount = "/var/lib/postgresql/data"
	case "sqlite":
		envVars = []string{}
		volumeMount = "/data"
	default:
		return fmt.Errorf("unsupported database type: %s", dc.Config.DBType)
	}

	if dc.containerExists(ctx) {
		log.Printf("database container already exists: %s", DatabaseContainerName)
		return nil
	}

	var containerConfig *container.Config
	var hostConfig *container.HostConfig

	containerConfig = &container.Config{
		Image: dc.Config.DBImageName,
		Env:   envVars,
		Labels: map[string]string{
			"hpotter": "database",
		},
	}

	hostConfig = &container.HostConfig{
		Mounts: []mount.Mount{
			{
				Type:   mount.TypeVolume,
				Source: DatabaseVolumeName,
				Target: volumeMount,
			},
		},
		RestartPolicy: container.RestartPolicy{
			Name: "unless-stopped",
		},
	}

	if dc.Config.DBType == "postgres" || dc.Config.DBType == "postgresql" {
		port, err := nat.NewPort("tcp", "5432")
		if err != nil {
			return fmt.Errorf("failed to create port: %w", err)
		}

		portInt, _ := strconv.Atoi("5432")
		containerPortBinding := nat.PortBinding{
			HostIP:   "127.0.0.1",
			HostPort: strconv.Itoa(portInt),
		}

		containerConfig.ExposedPorts = nat.PortSet{
			port: struct{}{},
		}
		hostConfig.PortBindings = nat.PortMap{
			port: []nat.PortBinding{containerPortBinding},
		}
	}

	createdContainer, err := dc.Client.ContainerCreate(
		ctx,
		containerConfig,
		hostConfig,
		&network.NetworkingConfig{
			EndpointsConfig: map[string]*network.EndpointSettings{
				DatabaseNetworkName: {},
			},
		},
		nil,
		DatabaseContainerName,
	)
	if err != nil {
		return fmt.Errorf("failed to create container: %w", err)
	}

	dc.ContainerID = createdContainer.ID

	err = dc.Client.ContainerStart(ctx, dc.ContainerID, container.StartOptions{})
	if err != nil {
		return fmt.Errorf("failed to start container: %w", err)
	}

	log.Printf("started database container: %s (%s)", DatabaseContainerName, dc.ContainerID)
	return nil
}

func (dc *DatabaseContainer) containerExists(ctx context.Context) bool {
	filters := filters.NewArgs()
	filters.Add("name", DatabaseContainerName)

	containers, err := dc.Client.ContainerList(ctx, container.ListOptions{
		All:     true,
		Filters: filters,
	})
	if err != nil {
		return false
	}

	for _, cont := range containers {
		if slices.Contains(cont.Names, "/"+DatabaseContainerName) {
			dc.ContainerID = cont.ID
			// Start the container if it's not running
			if cont.State != "running" {
				dc.Client.ContainerStart(ctx, cont.ID, container.StartOptions{})
			}
			return true
		}
	}
	return false
}

func (dc *DatabaseContainer) Setup(ctx context.Context) error {
	volumeExists, err := dc.CheckExistingVolume(ctx)
	if err != nil {
		return fmt.Errorf("failed to check existing volume: %w", err)
	}

	if volumeExists {
		log.Printf("using existing database volume")
		return dc.ensureContainerRunning(ctx)
	} else {
		log.Printf("creating new database infrastructure")
		return dc.createFreshDatabase(ctx)
	}
}

func (dc *DatabaseContainer) ensureContainerRunning(ctx context.Context) error {
	if dc.containerExists(ctx) {
		log.Printf("database container is already running")
		return nil
	}

	log.Printf("database volume exists but container is not running, starting container")

	if err := dc.CreateNetwork(ctx); err != nil {
		return fmt.Errorf("failed to create network: %w", err)
	}

	if err := dc.StartContainer(ctx); err != nil {
		return fmt.Errorf("failed to start container: %w", err)
	}

	return nil
}

func (dc *DatabaseContainer) createFreshDatabase(ctx context.Context) error {
	if err := dc.CreateVolume(ctx); err != nil {
		return fmt.Errorf("failed to create volume: %w", err)
	}

	if err := dc.CreateNetwork(ctx); err != nil {
		return fmt.Errorf("failed to create network: %w", err)
	}

	if err := dc.StartContainer(ctx); err != nil {
		return fmt.Errorf("failed to start container: %w", err)
	}

	return nil
}

func (dc *DatabaseContainer) Cleanup(ctx context.Context) error {
	if dc.ContainerID != "" {
		if err := dc.Client.ContainerRemove(ctx, dc.ContainerID, container.RemoveOptions{
			Force: true,
		}); err != nil {
			log.Printf("failed to remove database container: %v", err)
		} else {
			log.Printf("removed database container: %s", DatabaseContainerName)
		}
	}

	if dc.NetworkID != "" {
		if err := dc.Client.NetworkRemove(ctx, dc.NetworkID); err != nil {
			log.Printf("failed to remove database network: %v", err)
		} else {
			log.Printf("removed database network: %s", DatabaseNetworkName)
		}
	}

	// don't remove the volume to preserve data
	return nil
}
