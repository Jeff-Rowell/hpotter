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
		service:          service,
		credentialRegex:  regexp.MustCompile(``),
		sessionDataRegex: regexp.MustCompile(`(\d{1,3}\.){3}\d{1,3} - (.*)`),
	}

	return parser
}

func (p *HTTPDLogParser) ParseCredentials(allLogData string) *database.Credentials {
	// TODO
	return nil
}

func (p *HTTPDLogParser) ParseSessionData(allLogData string) *database.Data {
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
