package logparser

import (
	"fmt"
	"slices"
	"strings"

	"github.com/Jeff-Rowell/hpotter/internal/services"
	"github.com/Jeff-Rowell/hpotter/types"
)

type DefaultLogParserFactory struct{}

func NewLogParserFactory() LogParserFactory {
	return &DefaultLogParserFactory{}
}

func (f *DefaultLogParserFactory) CreateParser(service types.Service) (LogParser, error) {
	serviceRegistry := services.NewServiceRegistry()
	serviceName := strings.ToLower(serviceRegistry.GetServiceNameByConfig(service))

	switch serviceName {
	case "ssh":
		return NewSSHLogParser(service), nil
	case "telnet":
		return NewTelnetLogParser(service), nil
	case "httpd":
		return NewHttpdLogParser(service), nil
	default:
		return nil, fmt.Errorf("no log parser available for service: %s", serviceName)
	}
}

func (f *DefaultLogParserFactory) SupportedProtocols() []string {
	return []string{"ssh", "telnet", "httpd"}
}

func (f *DefaultLogParserFactory) IsSupported(service types.Service) bool {
	serviceRegistry := services.NewServiceRegistry()
	supported := f.SupportedProtocols()
	serviceName := strings.ToLower(serviceRegistry.GetServiceNameByConfig(service))

	return slices.Contains(supported, serviceName)
}
