package logparser

import (
	"log"
	"regexp"
	"strings"

	"github.com/Jeff-Rowell/hpotter/internal/database"
	"github.com/Jeff-Rowell/hpotter/types"
)

type SSHLogParser struct {
	service          types.Service
	credentialRegex  *regexp.Regexp
	sessionDataRegex *regexp.Regexp
}

func NewSSHLogParser(service types.Service) *SSHLogParser {
	parser := &SSHLogParser{
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

func (p *SSHLogParser) ParseCredentials(allLogData string) *database.Credentials {
	if p.credentialRegex == nil {
		return nil
	}

	matches := p.credentialRegex.FindAllStringSubmatch(allLogData, -1)

	for _, match := range matches {
		if len(match) < 3 {
			continue
		}

		username := strings.Trim(match[1], "b'\"")
		password := strings.Trim(match[2], "b'\"")

		return &database.Credentials{
			Username: username,
			Password: password,
		}
	}

	return nil
}

func (p *SSHLogParser) ParseSessionData(allLogData string) *database.Data {
	if p.sessionDataRegex == nil {
		return nil
	}

	var sessionData []string
	matches := p.sessionDataRegex.FindAllStringSubmatch(allLogData, -1)

	for _, match := range matches {
		if len(match) < 2 {
			continue
		}

		sessionData = append(sessionData, match[1])
	}

	return &database.Data{
		Direction: "request",
		Data:      strings.Join(sessionData, "\n"),
	}
}
