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

		// Get service details from registry by name
		var serviceName string
		if svc.Service != "" {
			// New format: service name specified
			serviceName = svc.Service
			if !serviceRegistry.IsSupported(serviceName) {
				log.Fatalf("error: unsupported service '%s'. %s", serviceName, serviceRegistry.GetSupportedServicesString())
			}

			// Populate port/protocol from registry
			registryService, _ := serviceRegistry.FindService(serviceName)
			svc.ListenPort = registryService.Port
			svc.ListenProto = registryService.Protocol
		} else {
			// Legacy format: port/protocol specified
			serviceName = serviceRegistry.GetServiceNameByConfig(*svc)
			if serviceName == "" {
				log.Fatalf("error: unsupported service configuration (port: %d, protocol: %s). %s", svc.ListenPort, svc.ListenProto, serviceRegistry.GetSupportedServicesString())
			}
		}

		if svc.CollectCredentials {
			if !serviceRegistry.IsSupportedByConfig(*svc) {
				log.Fatalf("error: service '%s' listening on '%d/%s' is not supported for credential collection: %s", serviceName, svc.ListenPort, svc.ListenProto, serviceRegistry.GetSupportedServicesString())
			}
		}

		// Auto-add required environment variables for telnet
		if serviceName == "telnet" {
			requiredEnvVars := serviceRegistry.GetRequiredEnvVars(*svc)
			for key, value := range requiredEnvVars {
				// Check if env var already exists
				found := false
				for j := range svc.EnvVars {
					if svc.EnvVars[j].Key == key {
						found = true
						break
					}
				}
				// Add if not found
				if !found {
					svc.EnvVars = append(svc.EnvVars, types.EnvVar{Key: key, Value: value})
				}
			}
		}

		if svc.CommandLimit > 0 {
			if serviceName != "ssh" && serviceName != "telnet" {
				log.Fatalf("error: command_limit is only allowed for ssh or telnet services, found in service '%s'", serviceName)
			}
		} else if serviceName == "ssh" || serviceName == "telnet" {
			svc.CommandLimit = 10
		}

		// Validate TLS options
		if svc.UseTLS || svc.CertificatePath != "" || svc.KeyPath != "" || svc.GenerateCerts {
			if serviceName != "httpd" {
				log.Fatalf("error: TLS options (use_tls, certificate_path, key_path, generate_certs) are only allowed for httpd service, found in service '%s'", serviceName)
			}

			if svc.UseTLS {
				if svc.GenerateCerts {
					if svc.CertificatePath != "" || svc.KeyPath != "" {
						log.Fatalf("error: certificate_path and key_path should not be set when generate_certs is true for service '%s'", serviceName)
					}
				} else {
					if svc.CertificatePath == "" {
						log.Fatalf("error: certificate_path is required when use_tls is true and generate_certs is false for service '%s'", serviceName)
					}
					if svc.KeyPath == "" {
						log.Fatalf("error: key_path is required when use_tls is true and generate_certs is false for service '%s'", serviceName)
					}
				}
			}

			if svc.GenerateCerts && !svc.UseTLS {
				log.Fatalf("error: generate_certs can only be true when use_tls is also true for service '%s'", serviceName)
			}
		}
	}
}
