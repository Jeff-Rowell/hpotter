package logparser

import (
	"log"
	"regexp"
	"strings"

	"github.com/Jeff-Rowell/hpotter/internal/database"
	"github.com/Jeff-Rowell/hpotter/types"
)

type HTTPDLogParser struct {
	service          types.Service
	credentialRegex  *regexp.Regexp
	sessionDataRegex *regexp.Regexp
}

func NewHttpdLogParser(service types.Service) *HTTPDLogParser {
	parser := &HTTPDLogParser{
		service: service,
	}

	if service.CredentialLogPattern != "" {
		parser.credentialRegex = regexp.MustCompile(service.CredentialLogPattern)
	} else {
		log.Fatalf("error: 'credential_log_pattern' is required.")
	}

	if service.SessionDataLogPattern != "" {
		parser.sessionDataRegex = regexp.MustCompile(service.SessionDataLogPattern)
	} else {
		log.Fatalf("error: 'session_data_log_pattern' is required.")
	}

	return parser
}

func (p *HTTPDLogParser) ParseCredentials(allLogData string) *database.Credentials {
	// TODO
	return nil
}

func (p *HTTPDLogParser) ParseSessionData(allLogData string) *database.Data {
	if p.sessionDataRegex == nil {
		return nil
	}

	var sessionData []string
	matches := p.sessionDataRegex.FindAllString(allLogData, -1)

	for _, match := range matches {
		log.Printf("DEBUG: match: %s", match)
		if len(match) < 1 {
			continue
		}

		sessionData = append(sessionData, match)
	}

	return &database.Data{
		Direction: "request",
		Data:      strings.Join(sessionData, "\n"),
	}
}
