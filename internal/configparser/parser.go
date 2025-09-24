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

	for i := range p.Services {
		svc := &p.Services[i]
		if svc.CollectCredentials {
			if svc.ServiceName == "" {
				log.Fatalf("error: service_name is required")
			}
			if !serviceRegistry.IsSupported(svc.ServiceName) {
				log.Fatalf("error: service '%s' listening on '%d/%s' is not supported for credential collection: %s", svc.ServiceName, svc.ListenPort, svc.ListenProto, serviceRegistry.GetSupportedServicesString())
			}
		}

		if svc.CommandLimit > 0 {
			if svc.ServiceName != "ssh" && svc.ServiceName != "telnet" {
				log.Fatalf("error: command_limit is only allowed for ssh or telnet services, found in service '%s'", svc.ServiceName)
			}
		} else if svc.ServiceName == "ssh" || svc.ServiceName == "telnet" {
			svc.CommandLimit = 10
		}

		// Validate TLS options
		if svc.UseTLS || svc.CertificatePath != "" || svc.KeyPath != "" || svc.GenerateCerts {
			if svc.ServiceName != "httpd" {
				log.Fatalf("error: TLS options (use_tls, certificate_path, key_path, generate_certs) are only allowed for httpd service, found in service '%s'", svc.ServiceName)
			}

			if svc.UseTLS {
				if svc.GenerateCerts {
					if svc.CertificatePath != "" || svc.KeyPath != "" {
						log.Fatalf("error: certificate_path and key_path should not be set when generate_certs is true for service '%s'", svc.ServiceName)
					}
				} else {
					if svc.CertificatePath == "" {
						log.Fatalf("error: certificate_path is required when use_tls is true and generate_certs is false for service '%s'", svc.ServiceName)
					}
					if svc.KeyPath == "" {
						log.Fatalf("error: key_path is required when use_tls is true and generate_certs is false for service '%s'", svc.ServiceName)
					}
				}
			}

			if svc.GenerateCerts && !svc.UseTLS {
				log.Fatalf("error: generate_certs can only be true when use_tls is also true for service '%s'", svc.ServiceName)
			}
		}
	}
}
