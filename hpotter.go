package main

import (
	"context"
	"flag"
	"log"
	"log/slog"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/hpotter/pkg/container"
	"github.com/hpotter/pkg/database"
	"github.com/hpotter/pkg/dockerclient"
	"github.com/hpotter/pkg/listener"
	"github.com/hpotter/pkg/logger"
	"github.com/hpotter/pkg/parser"
	"github.com/hpotter/pkg/terminator"
)

var debugMode bool
var logFilePath, configFilePath string

// init initializes the cli flags for hpotter
func init() {
	const (
		defaultLogFile    = "./hpotter.log"
		logFileUsage      = "the path to the log file to write to"
		defaultConfigFile = "./config.yml"
		configFileUsage   = "the path to the config file to write to"
		debugUsage        = "turn on debug logging"
	)

	flag.StringVar(&logFilePath, "log-file", defaultLogFile, logFileUsage)
	flag.StringVar(&logFilePath, "l", defaultLogFile, logFileUsage+" (shorthand)")

	flag.StringVar(&configFilePath, "config-file", defaultConfigFile, configFileUsage)
	flag.StringVar(&configFilePath, "c", defaultConfigFile, configFileUsage+" (shorthand)")

	flag.BoolVar(&debugMode, "debug", false, debugUsage)
	flag.BoolVar(&debugMode, "d", false, debugUsage+" (shorthand)")
}

// main is the main function
func main() {
	flag.Parse()

	err := logger.InitLogger(logFilePath, debugMode)
	if err != nil {
		log.Fatal(err)
	}

	slog.Info("hpotter started...")

	confParser := parser.NewYAMLParser()
	config, err := confParser.Parse(configFilePath)
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	dockerClient, err := dockerclient.NewDockerClient()
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	d := &database.Database{
		DbImage:      config.Database.Image,
		DbPort:       config.Database.Port,
		Username:     config.Database.Username,
		Password:     config.Database.Password,
		DockerClient: dockerClient,
	}

	ctx, cancelFunc := context.WithCancel(context.Background())
	go handleShutdown(cancelFunc, d)

	handleMissingImages(config, dockerClient)

	err = d.Create()
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	terminator := terminator.NewTlsTerminator(ctx)

	var wg sync.WaitGroup
	l := &listener.Listener{
		Connections:     []net.Conn{},
		Wg:              &wg,
		Ctx:             ctx,
		ContainerThread: &container.ContainerThread{DockerClient: dockerClient},
		Terminator:      terminator,
		Db:              d,
	}

	for _, svc := range config.Services {
		wg.Add(1)
		go l.Listen(svc)
	}
	slog.Info("listening...")
	wg.Wait()
}

// handleShutdown is a routine that listens for a SIGTERM or SIGINT signal
// and attempts to orchestrate a graceful shutdown of all connections and
// running containers.
func handleShutdown(cancelFunc context.CancelFunc, db *database.Database) {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGTERM, syscall.SIGINT)

	slog.Info("waiting for cancellation signal...")
	<-sigChan

	slog.Info("cancellation signal received, shutting down hpotter")
	signal.Stop(sigChan)

	db.HandleShutdown()
	cancelFunc()
}

// handleMissingImages iterates through the configuration and checks if
// the container images need to be pulled.
func handleMissingImages(config *parser.Config, dc *dockerclient.Docker) {
	for _, svc := range config.Services {
		checkImage(dc, svc.Image)
	}

	checkImage(dc, config.Database.Image)
}

// checkImage checks if the given container image already exists on the host,
// and if not, calls the Docker API client to pull the image.
func checkImage(d *dockerclient.Docker, i string) {
	hasImage, err := d.HasImage(i)
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
	if !hasImage {
		slog.Info("image not found on host", "image", i)
		slog.Info("attempting to download image")
		err := d.PullImage(i)
		if err != nil {
			slog.Error(err.Error())
			os.Exit(1)
		}
		slog.Info("successfully downloaded image", "image", i)
	}
}
