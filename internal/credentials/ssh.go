package credentials

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/Jeff-Rowell/hpotter/internal/database"
	"golang.org/x/crypto/ssh"
)

type SSHCredentialCollector struct {
	conn          net.Conn
	db            *database.Database
	dbConn        *database.Connections
	hostKey       ssh.Signer
	config        *ssh.ServerConfig
	capturedCreds *database.Credentials
}

// NewSSHCredentialCollector creates a new SSH credential collector
func NewSSHCredentialCollector(conn net.Conn, db *database.Database, dbConn *database.Connections) (*SSHCredentialCollector, error) {
	collector := &SSHCredentialCollector{
		conn:   conn,
		db:     db,
		dbConn: dbConn,
	}

	if err := collector.generateHostKey(); err != nil {
		return nil, fmt.Errorf("failed to generate host key: %w", err)
	}

	collector.config = &ssh.ServerConfig{
		PasswordCallback:  collector.passwordCallback,
		PublicKeyCallback: collector.pubKeyCallback,
		AuthLogCallback:   collector.authLogCallback,
	}
	collector.config.AddHostKey(collector.hostKey)

	return collector, nil
}

func (s *SSHCredentialCollector) generateHostKey() error {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return fmt.Errorf("failed to generate RSA key: %w", err)
	}

	privateKeyPEM := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	}

	privateKeyBytes := pem.EncodeToMemory(privateKeyPEM)

	hostKey, err := ssh.ParsePrivateKey(privateKeyBytes)
	if err != nil {
		return fmt.Errorf("failed to parse private key: %w", err)
	}

	s.hostKey = hostKey
	return nil
}

func (s *SSHCredentialCollector) passwordCallback(conn ssh.ConnMetadata, password []byte) (*ssh.Permissions, error) {
	username := conn.User()
	passwordStr := string(password)

	log.Printf("SSH login attempt: user=%s from=%s", username, conn.RemoteAddr())

	creds := &database.Credentials{
		Username:      username,
		Password:      passwordStr,
		ConnectionsID: s.dbConn.ID,
	}

	s.capturedCreds = creds

	if err := s.db.Write(creds); err != nil {
		log.Printf("failed to save SSH credentials: %v", err)
		return nil, fmt.Errorf("authentication failed")
	}

	return nil, nil
}

func (s *SSHCredentialCollector) pubKeyCallback(conn ssh.ConnMetadata, key ssh.PublicKey) (*ssh.Permissions, error) {
	username := conn.User()
	keyType := key.Type()
	fingerprint := ssh.FingerprintSHA256(key)
	log.Printf("SSH public key attempt: user=%s, type=%s, fingerprint=%s from=%s", username, keyType, fingerprint, conn.RemoteAddr())
	return nil, fmt.Errorf("public key authentication not accepted")
}

func (s *SSHCredentialCollector) authLogCallback(conn ssh.ConnMetadata, method string, err error) {
	log.Printf("SSH auth attempt: user=%s, method=%s, success=%v from=%s", conn.User(), method, err == nil, conn.RemoteAddr())
}

// CollectCredentials performs SSH handshake and collects credentials
func (s *SSHCredentialCollector) CollectCredentials() (*database.Credentials, error) {
	s.conn.SetDeadline(time.Now().Add(CredentialsTimeout))
	defer s.conn.SetDeadline(time.Time{})

	serverConn, _, _, err := ssh.NewServerConn(s.conn, s.config)
	if err != nil {
		return nil, fmt.Errorf("SSH handshake failed: %w", err)
	}
	defer serverConn.Close()

	if s.capturedCreds == nil {
		return nil, fmt.Errorf("no credentials were captured during authentication")
	}

	return s.capturedCreds, nil
}
