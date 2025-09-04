package services

import (
	"fmt"
	"strings"
)

type ServiceType string

type CredentialService struct {
	Name        string
	Protocol    string
	Port        int
	Description string
}

var SupportedServices = []CredentialService{
	{
		Name:        "telnet",
		Description: "Telnet service with username/password collection",
	},
	{
		Name:        "ssh",
		Description: "SSH service with username/password collection",
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
