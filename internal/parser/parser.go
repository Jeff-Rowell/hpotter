package parser

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"

	"github.com/Jeff-Rowell/hpotter/types"
)

type Parser struct {
	DBConfig dbConfig        `json:"db_config"`
	Services []types.Service `json:"services"`
}

type dbConfig struct {
	DBType   string `json:"db_type"`
	Name     string `json:"name"`
	User     string `json:"user"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     int16  `json:"port"`
}

func NewParser() Parser {
	return Parser{}
}

func (p *Parser) Parse(configJson string) {
	cleanConfigJson := filepath.Clean(configJson)
	log.Printf("creating new parser from '%s'", cleanConfigJson)
	data, err := os.ReadFile(cleanConfigJson)
	if err != nil {
		log.Fatalf("error: failed to read data from file '%s': %v", cleanConfigJson, err)
	}

	if err := json.Unmarshal(data, &p); err != nil {
		log.Println(err)
		log.Fatalf("error: failed to unmarshal data from file '%s': %v", cleanConfigJson, err)
	}
}
