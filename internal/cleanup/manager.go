package cleanup

import (
	"context"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/moby/moby/api/types/container"
	"github.com/moby/moby/api/types/filters"
	"github.com/moby/moby/client"
)

type GlobalContainerManager struct {
	dockerClient *client.Client
}

var globalManager *GlobalContainerManager
var once sync.Once

func GetGlobalContainerManager() *GlobalContainerManager {
	once.Do(func() {
		dockerClient, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
		if err != nil {
			log.Fatalf("failed to create docker client for cleanup: %v", err)
		}

		globalManager = &GlobalContainerManager{
			dockerClient: dockerClient,
		}
	})
	return globalManager
}

func (gcm *GlobalContainerManager) CleanupAllHPotterContainers() {
	log.Printf("cleaning up all HPotter containers...")

	listCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	filters := filters.NewArgs()
	filters.Add("label", "hpotter=container")
	containers, err := gcm.dockerClient.ContainerList(listCtx, container.ListOptions{
		Filters: filters,
		All:     true,
	})
	if err != nil {
		log.Printf("error listing hpotter containers: %v", err)
		return
	}

	if len(containers) == 0 {
		log.Printf("no HPotter containers found to cleanup")
		return
	}

	log.Printf("found %d HPotter containers to cleanup", len(containers))
	for _, cont := range containers {
		gcm.removeContainer(cont.ID, cont.Image)
	}

	gcm.cleanupTempConfigFiles()

	gcm.dockerClient.Close()
	log.Printf("HPotter container cleanup complete")
}

func (gcm *GlobalContainerManager) removeContainer(containerID, imageName string) {
	cleanupCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	log.Printf("removing container %s (image: %s)", containerID[:12], imageName)

	stopCtx, stopCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer stopCancel()

	err := gcm.dockerClient.ContainerStop(stopCtx, containerID, container.StopOptions{})
	if err != nil {
		log.Printf("error stopping container %s: %v", containerID[:12], err)
	}

	removeOpts := container.RemoveOptions{
		Force: true,
	}
	if err := gcm.dockerClient.ContainerRemove(cleanupCtx, containerID, removeOpts); err != nil {
		log.Printf("error removing container %s: %v", containerID[:12], err)
	} else {
		log.Printf("successfully removed container %s", containerID[:12])
	}
}

func (gcm *GlobalContainerManager) cleanupTempConfigFiles() {
	tempDir := os.TempDir()
	
	// Clean up httpd config files
	configPattern := filepath.Join(tempDir, "httpd-*.conf")
	configMatches, err := filepath.Glob(configPattern)
	if err != nil {
		log.Printf("error finding temp config files: %v", err)
	} else if len(configMatches) > 0 {
		log.Printf("cleaning up %d temporary config files", len(configMatches))
		for _, file := range configMatches {
			if err := os.Remove(file); err != nil {
				log.Printf("warning: failed to remove temporary config file %s: %v", file, err)
			} else {
				log.Printf("removed temporary config file: %s", file)
			}
		}
	}
	
	// Clean up generated certificate files
	certPattern := filepath.Join(tempDir, "hpotter-cert-*.crt")
	certMatches, err := filepath.Glob(certPattern)
	if err != nil {
		log.Printf("error finding temp certificate files: %v", err)
	} else if len(certMatches) > 0 {
		log.Printf("cleaning up %d temporary certificate files", len(certMatches))
		for _, file := range certMatches {
			if err := os.Remove(file); err != nil {
				log.Printf("warning: failed to remove temporary certificate file %s: %v", file, err)
			} else {
				log.Printf("removed temporary certificate file: %s", file)
			}
		}
	}
	
	// Clean up generated key files
	keyPattern := filepath.Join(tempDir, "hpotter-key-*.key")
	keyMatches, err := filepath.Glob(keyPattern)
	if err != nil {
		log.Printf("error finding temp key files: %v", err)
	} else if len(keyMatches) > 0 {
		log.Printf("cleaning up %d temporary key files", len(keyMatches))
		for _, file := range keyMatches {
			if err := os.Remove(file); err != nil {
				log.Printf("warning: failed to remove temporary key file %s: %v", file, err)
			} else {
				log.Printf("removed temporary key file: %s", file)
			}
		}
	}
	
	totalFiles := len(configMatches) + len(certMatches) + len(keyMatches)
	if totalFiles == 0 {
		log.Printf("no temporary files found to cleanup")
	}
}
