package logparser

import (
	"slices"
	"fmt"
	"strings"

	"github.com/Jeff-Rowell/hpotter/types"
)

type DefaultLogParserFactory struct{}

func NewLogParserFactory() LogParserFactory {
	return &DefaultLogParserFactory{}
}

func (f *DefaultLogParserFactory) CreateParser(service types.Service) (LogParser, error) {
	switch strings.ToLower(service.ServiceName) {
	case "ssh":
		return NewSSHLogParser(service), nil
	case "telnet":
		return NewTelnetLogParser(service), nil
	default:
		return nil, fmt.Errorf("no log parser available for service: %s", service.ServiceName)
	}
}

func (f *DefaultLogParserFactory) SupportedProtocols() []string {
	return []string{"ssh", "telnet"}
}

func (f *DefaultLogParserFactory) IsSupported(service types.Service) bool {
	supported := f.SupportedProtocols()
	serviceName := strings.ToLower(service.ServiceName)

	return slices.Contains(supported, serviceName)
}
