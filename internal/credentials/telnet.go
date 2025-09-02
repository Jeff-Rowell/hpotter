package credentials

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"time"

	"github.com/Jeff-Rowell/hpotter/internal/database"
)

const (
	TelnetLoginPrompt    = "login: "
	TelnetPasswordPrompt = "Password: "
	CredentialsTimeout   = 30 * time.Second
	MaxCredentialLength  = 256

	// Telnet protocol constants
	IAC  = 255 // Interpret As Command
	WILL = 251 // Will option
	WONT = 252 // Won't option
	DO   = 253 // Do option
	DONT = 254 // Don't option
	ECHO = 1   // Echo option
)

type TelnetCredentialCollector struct {
	conn   net.Conn
	db     *database.Database
	dbConn *database.Connections
}

// NewTelnetCredentialCollector creates a new telnet credential collector
func NewTelnetCredentialCollector(conn net.Conn, db *database.Database, dbConn *database.Connections) *TelnetCredentialCollector {
	return &TelnetCredentialCollector{
		conn:   conn,
		db:     db,
		dbConn: dbConn,
	}
}

// CollectCredentials presents a fake telnet login and collects credentials
func (t *TelnetCredentialCollector) CollectCredentials() (*database.Credentials, error) {
	t.conn.SetDeadline(time.Now().Add(CredentialsTimeout))
	defer t.conn.SetDeadline(time.Time{})

	username, err := t.promptForCredential(TelnetLoginPrompt, false)
	if err != nil {
		return nil, fmt.Errorf("failed to get username: %w", err)
	}

	password, err := t.promptForCredential(TelnetPasswordPrompt, true)
	if err != nil {
		return nil, fmt.Errorf("failed to get password: %w", err)
	}

	creds := &database.Credentials{
		Username:      username,
		Password:      password,
		ConnectionsID: t.dbConn.ID,
	}

	if err := t.db.Write(creds); err != nil {
		return nil, fmt.Errorf("failed to save credentials: %w", err)
	}

	return creds, nil
}

func (t *TelnetCredentialCollector) promptForCredential(prompt string, isPassword bool) (string, error) {
	maxAttempts := 3

	for attempt := range maxAttempts {
		if _, err := t.conn.Write([]byte(prompt)); err != nil {
			return "", fmt.Errorf("failed to send prompt: %w", err)
		}

		var response string
		var err error
		if isPassword {
			if err := t.disableEcho(); err != nil {
				return "", fmt.Errorf("failed to disable echo: %w", err)
			}
			response, err = t.readLine()
			t.enableEcho()
		} else {
			response, err = t.readLine()
		}

		if err != nil {
			return "", fmt.Errorf("failed to read response: %w", err)
		}

		// Remove the null byte added by telnet
		response = strings.TrimRight(response, "\x00")

		if len(response) > 0 && len(response) <= MaxCredentialLength {
			if isPassword {
				t.conn.Write([]byte("\r\n"))
			}
			return response, nil
		}

		if attempt == maxAttempts-1 {
			return "", fmt.Errorf("no valid response after %d attempts", maxAttempts)
		}
	}

	return "", fmt.Errorf("failed to get credential")
}

func (t *TelnetCredentialCollector) readLine() (string, error) {
	reader := bufio.NewReader(t.conn)
	var line strings.Builder

	for {
		b, err := reader.ReadByte()
		if err != nil {
			return "", err
		}

		// Handle telnet protocol bytes
		if b == 255 { // IAC (Interpret As Command)
			// Read and ignore next two bytes for telnet negotiation
			reader.ReadByte()
			reader.ReadByte()
			continue
		}

		// Check for line endings
		switch b {
		case '\r':
			next, err := reader.ReadByte()
			if err != nil {
				return line.String(), nil
			}
			if next != '\n' {
				line.WriteByte(next)
			}
			return line.String(), nil
		case '\n':
			return line.String(), nil
		}

		line.WriteByte(b)

		if line.Len() > MaxCredentialLength {
			return line.String(), nil
		}
	}
}

func (t *TelnetCredentialCollector) disableEcho() error {
	cmd := []byte{IAC, WILL, ECHO}
	_, err := t.conn.Write(cmd)
	return err
}

func (t *TelnetCredentialCollector) enableEcho() error {
	cmd := []byte{IAC, WONT, ECHO}
	_, err := t.conn.Write(cmd)
	return err
}
