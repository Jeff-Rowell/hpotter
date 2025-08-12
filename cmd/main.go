package main

import (
	"flag"
	"log"
	"sync"

	"github.com/Jeff-Rowell/hpotter/internal/parser"
	"github.com/Jeff-Rowell/hpotter/internal/threads"
)

type flags struct {
	configJson string
}

func main() {
	flags := parseFlags()
	config := parser.NewParser()
	config.Parse(flags.configJson)

	var wg sync.WaitGroup
	log.Printf("starting %d socket listeners", len(config.Services))
	for _, serviceCfg := range config.Services {
		wg.Add(1)
		go threads.StartListener(serviceCfg, &wg)
	}
	wg.Wait()
}

func parseFlags() flags {
	configJson := flag.String("config", "config.json", "the absolute or relative path to json config")
	flag.Parse()
	return flags{
		configJson: *configJson,
	}
}
