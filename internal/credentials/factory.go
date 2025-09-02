package credentials

import (
	"fmt"
	"net"
	"strings"

	"github.com/Jeff-Rowell/hpotter/internal/database"
	"github.com/Jeff-Rowell/hpotter/types"
)

// CredentialCollector interface for different protocol implementations
type CredentialCollector interface {
	CollectCredentials() (*database.Credentials, error)
}

// NewCredentialCollector creates an appropriate credential collector based on service configuration
func NewCredentialCollector(service types.Service, conn net.Conn, db *database.Database, dbConn *database.Connections) (CredentialCollector, error) {
	protocol := strings.ToLower(service.ListenProto)

	switch {
	case protocol == "tcp" && service.ListenPort == 23:
		return NewTelnetCredentialCollector(conn, db, dbConn), nil
	default:
		return nil, fmt.Errorf("no credential collector available for service on port %d", service.ListenPort)
	}
}
