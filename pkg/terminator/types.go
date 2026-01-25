package terminator

import "context"

// Terminator is responsible for keepings track of generated TLS key pairs
// and cleaning them up upon program termination.
type Terminator struct {
	// Certs is a map of generated TLS key pairs for tracking purposes.
	Certs []TlsKeyPair

	// TerminatorActorChan is the channel used to communicate with the
	// Terminator actor for managing TLS key pairs.
	TerminatorActorChan chan TerminatorMessage

	// ctx is the main context.Context to listen for cancellation signals on
	ctx context.Context
}

// MessageType is the type of message being sent to the TerminatorActor
type MessageType int

const (
	AddCert MessageType = iota
	RemoveCert
	GetCerts
)

// TlsKeyPair keeps reference to a TLS certificate and key file path
type TlsKeyPair struct {
	// CertPath is the file path to the TLS certificate
	CertPath string

	// KeyPath is the file path to the TLS key
	KeyPath string
}

// TerminatorMessage represents a job for the TerminatorActor
type TerminatorMessage struct {
	Type     MessageType
	Value    TlsKeyPair
	Response chan []TlsKeyPair
}
