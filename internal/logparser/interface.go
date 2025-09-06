package logparser

import (
	"github.com/Jeff-Rowell/hpotter/internal/database"
	"github.com/Jeff-Rowell/hpotter/types"
)

// LogParser defines the interface for protocol-specific log parsing
type LogParser interface {
	ParseCredentials(allLogData string) *database.Credentials
	ParseSessionData(allLogData string) *database.Data
}

// LogParserFactory creates appropriate log parsers based on service configuration
type LogParserFactory interface {
	CreateParser(service types.Service) (LogParser, error)
	SupportedProtocols() []string
	IsSupported(service types.Service) bool
}
