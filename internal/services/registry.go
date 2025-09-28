package services

import (
	"fmt"
	"strings"

	"github.com/Jeff-Rowell/hpotter/types"
)

type ServiceType string

type CredentialService struct {
	Name            string
	Protocol        string
	Port            int
	Description     string
	ImageName       string
	RequiredEnvVars map[string]string // Environment variables required for this service
}

var SupportedServices = []CredentialService{
	{
		Name:        "telnet",
		Protocol:    "tcp",
		Port:        2223,
		Description: "Telnet service with username/password collection",
		ImageName:   "cowrie/cowrie:latest",
		RequiredEnvVars: map[string]string{
			"COWRIE_TELNET_ENABLED": "yes",
		},
	},
	{
		Name:        "ssh",
		Protocol:    "tcp",
		Port:        2222,
		Description: "SSH service with username/password collection",
		ImageName:   "cowrie/cowrie:latest",
	},
	{
		Name:        "httpd",
		Protocol:    "tcp",
		Port:        8080,
		Description: "HTTPd service with username/password collection",
		ImageName:   "httpd:2.4.65",
	},
}

type ServiceRegistry struct{}

func NewServiceRegistry() *ServiceRegistry {
	return &ServiceRegistry{}
}

func (sr *ServiceRegistry) IsSupported(name string) bool {
	_, found := sr.FindService(name)
	return found
}

func (sr *ServiceRegistry) FindService(name string) (*CredentialService, bool) {
	normalizedName := strings.ToLower(strings.TrimSpace(name))

	for _, service := range SupportedServices {
		serviceNameMatches := strings.EqualFold(service.Name, normalizedName)
		if serviceNameMatches {
			return &service, true
		}
	}
	return nil, false
}

func (sr *ServiceRegistry) GetSupportedServicesString() string {
	output := "supported services:"
	for _, svc := range SupportedServices {
		output += fmt.Sprintf(" service name: %s", svc.Name)
	}
	return output
}

func (sr *ServiceRegistry) GetImageName(serviceName string) string {
	service, found := sr.FindService(serviceName)
	if found {
		return service.ImageName
	}
	return ""
}

// FindServiceByPort finds a service by matching port and protocol
func (sr *ServiceRegistry) FindServiceByPort(port int, protocol string) (*CredentialService, bool) {
	normalizedProto := strings.ToLower(strings.TrimSpace(protocol))

	for _, service := range SupportedServices {
		if service.Port == port && strings.EqualFold(service.Protocol, normalizedProto) {
			return &service, true
		}
	}
	return nil, false
}

// GetServiceByConfig determines service type from Service config (port/protocol)
func (sr *ServiceRegistry) GetServiceByConfig(service types.Service) (*CredentialService, bool) {
	return sr.FindServiceByPort(service.ListenPort, service.ListenProto)
}

// GetImageNameByConfig returns image name for a Service config
func (sr *ServiceRegistry) GetImageNameByConfig(service types.Service) string {
	detectedService, found := sr.GetServiceByConfig(service)
	if found {
		return detectedService.ImageName
	}
	return ""
}

// GetServiceNameByConfig returns service name for a Service config
func (sr *ServiceRegistry) GetServiceNameByConfig(service types.Service) string {
	detectedService, found := sr.GetServiceByConfig(service)
	if found {
		return detectedService.Name
	}
	return ""
}

// IsSupportedByConfig checks if a Service config is supported
func (sr *ServiceRegistry) IsSupportedByConfig(service types.Service) bool {
	_, found := sr.GetServiceByConfig(service)
	return found
}

// GetRequiredEnvVars returns required environment variables for a Service config
func (sr *ServiceRegistry) GetRequiredEnvVars(service types.Service) map[string]string {
	detectedService, found := sr.GetServiceByConfig(service)
	if found && detectedService.RequiredEnvVars != nil {
		return detectedService.RequiredEnvVars
	}
	return make(map[string]string)
}

const DefaultDBImageName = "postgres:17.6"
