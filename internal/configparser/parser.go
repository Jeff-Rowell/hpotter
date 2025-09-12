package configparser

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"

	"github.com/Jeff-Rowell/hpotter/internal/services"
	"github.com/Jeff-Rowell/hpotter/types"
)

type Parser struct {
	DBConfig types.DBConfig  `json:"db_config"`
	Services []types.Service `json:"services"`
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

	serviceRegistry := services.NewServiceRegistry()

	for _, svc := range p.Services {
		if svc.CollectCredentials {
			if svc.ServiceName == "" {
				log.Fatalf("error: service_name is required")
			}
			if !serviceRegistry.IsSupported(svc.ServiceName) {
				log.Fatalf("error: service '%s' listening on '%d/%s' is not supported for credential collection: %s", svc.ServiceName, svc.ListenPort, svc.ListenProto, serviceRegistry.GetSupportedServicesString())
			}
		}
	}
}
