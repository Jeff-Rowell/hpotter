package configparser

import (
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"

	"github.com/Jeff-Rowell/hpotter/internal/services"
	"github.com/Jeff-Rowell/hpotter/types"
)

type Parser struct {
	DBConfig types.DBConfig  `yaml:"db_config"`
	Services []types.Service `yaml:"services"`
}

func NewParser() Parser {
	return Parser{}
}

func (p *Parser) Parse(configFile string) {
	cleanConfigFile := filepath.Clean(configFile)
	log.Printf("creating new parser from '%s'", cleanConfigFile)
	data, err := os.ReadFile(cleanConfigFile)
	if err != nil {
		log.Fatalf("error: failed to read data from file '%s': %v", cleanConfigFile, err)
	}

	if err := yaml.Unmarshal(data, &p); err != nil {
		log.Println(err)
		log.Fatalf("error: failed to unmarshal data from file '%s': %v", cleanConfigFile, err)
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

		// Validate TLS options
		if svc.UseTLS || svc.CertificatePath != "" || svc.KeyPath != "" {
			if svc.ServiceName != "httpd" {
				log.Fatalf("error: TLS options (use_tls, certificate_path, key_path) are only allowed for httpd service, found in service '%s'", svc.ServiceName)
			}
			
			if svc.UseTLS {
				if svc.CertificatePath == "" {
					log.Fatalf("error: certificate_path is required when use_tls is true for service '%s'", svc.ServiceName)
				}
				if svc.KeyPath == "" {
					log.Fatalf("error: key_path is required when use_tls is true for service '%s'", svc.ServiceName)
				}
			}
		}
	}
}
