package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/Jeff-Rowell/hpotter/internal/cleanup"
	"github.com/Jeff-Rowell/hpotter/internal/configparser"
	"github.com/Jeff-Rowell/hpotter/internal/database"
	"github.com/Jeff-Rowell/hpotter/internal/frontend"
	"github.com/Jeff-Rowell/hpotter/internal/threads"
)

type flags struct {
	configFile string
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigChan
		log.Printf("main: received cancellation signal")
		cancel()
	}()

	flags := parseFlags()
	config := configparser.NewParser()
	config.Parse(flags.configFile)

	dbContainer, err := database.NewDatabaseContainer(ctx, config.DBConfig)
	if err != nil {
		log.Fatalf("failed to create database container manager: %v", err)
	}

	if err := dbContainer.Setup(); err != nil {
		log.Fatalf("failed to setup database: %v", err)
	}

	defer func() {
		log.Printf("cleaning up resources...")
		containerManager := cleanup.GetGlobalContainerManager()
		containerManager.CleanupAllHPotterContainers()

		log.Printf("cleaning up database resources...")
		cleanupCtx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
		defer cancel()
		if err := dbContainer.Cleanup(cleanupCtx); err != nil {
			log.Printf("failed to cleanup database resources: %v", err)
		}
	}()

	db, err := database.NewDatabase(config.DBConfig)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer db.Close()

	log.Printf("database initialized successfully")

	var wg sync.WaitGroup

	if config.FrontendConfig.Enabled {
		wg.Add(1)
		go func() {
			defer wg.Done()
			frontendServer := frontend.NewServer(":80")
			if err := frontendServer.Start(); err != nil {
				log.Printf("frontend server error: %v", err)
			}
		}()
	}

	log.Printf("starting %d socket listeners", len(config.Services))
	for _, serviceCfg := range config.Services {
		wg.Add(1)
		go threads.StartListener(serviceCfg, &wg, ctx, db)
	}
	wg.Wait()
}

func parseFlags() flags {
	configFile := flag.String("config", "config.yaml", "the absolute or relative path to yaml config")
	flag.Parse()
	return flags{
		configFile: *configFile,
	}
}
