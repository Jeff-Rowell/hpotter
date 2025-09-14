package logparser

import (
	"log"
	"regexp"
	"strings"

	"encoding/base64"

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
		credentialRegex:  regexp.MustCompile(`AUTH: Basic ([^"]{0,})`),
		sessionDataRegex: regexp.MustCompile(`(\d{1,3}\.){3}\d{1,3} - (.*)`),
	}

	return parser
}

func (p *HTTPDLogParser) ParseCredentials(allLogData string) *database.Credentials {
	matches := p.credentialRegex.FindAllStringSubmatch(allLogData, -1)

	for _, match := range matches {
		if len(match) > 1 {
			base64Decoded, err := base64.StdEncoding.DecodeString(match[1])
			if err != nil {
				log.Fatalf("error base64 decoding authorization header: %v", err)
			}

			decodedParts := strings.Split(string(base64Decoded), ":")

			if len(decodedParts) < 2 {
				log.Fatalln("error base64 decoding authorization header")
			}

			return &database.Credentials{
				Username: decodedParts[0],
				Password: decodedParts[1],
			}
		}
	}

	return nil
}

func (p *HTTPDLogParser) ParseSessionData(allLogData string) *database.Data {
	var sessionData []string
	matches := p.sessionDataRegex.FindAllString(allLogData, -1)

	for _, match := range matches {
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
