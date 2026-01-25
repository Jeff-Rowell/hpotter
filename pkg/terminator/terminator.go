package terminator

import (
	"context"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"io"
	"log/slog"
	"math/big"
	"net"
	"os"
	"time"

	"github.com/hashicorp/go-retryablehttp"
)

func NewTlsTerminator(ctx context.Context) *Terminator {
	terminatorActorChan := make(chan TerminatorMessage)
	go manageTlsKeyPairs(terminatorActorChan)
	t := Terminator{
		ctx:                 ctx,
		Certs:               []TlsKeyPair{},
		TerminatorActorChan: terminatorActorChan,
	}
	go t.handleCancel()
	return &t
}

func (t *Terminator) GenerateTlsKeyPair() (*tls.Config, error) {
	priv, err := ecdsa.GenerateKey(elliptic.P384(), rand.Reader)
	if err != nil {
		slog.Error("failed to generate private key", "error", err)
		os.Exit(1)
	}

	notBefore := time.Now().AddDate(0, -6, 0)
	notAfter := notBefore.AddDate(1, 0, 0)

	serialNumber := new(big.Int)
	serialNumber.SetBytes([]byte{
		0x00,
		0xAA,
		0xEA,
		0x6A,
		0x58,
		0x2A,
		0x83,
		0xA5,
		0x73,
		0x20,
		0xB5,
		0xAD,
		0x43,
		0x29,
		0xB3,
		0xDB,
		0xBA,
	})

	template := x509.Certificate{
		SerialNumber: serialNumber,
		NotBefore:    notBefore,
		NotAfter:     notAfter,
		KeyUsage:     x509.KeyUsageDigitalSignature,
		ExtKeyUsage:  []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		Subject: pkix.Name{
			Country:            []string{"UK"},
			Organization:       []string{"Hogwarts"},
			OrganizationalUnit: []string{"Gryffindor"},
			Locality:           []string{"Gordic's Hallow"},
			Province:           []string{"Edinburgh"},
			StreetAddress:      []string{"4 Privet Drive"},
			PostalCode:         []string{"62442"},
			CommonName:         "HPotter",
			SerialNumber:       serialNumber.String(),
		},
		IPAddresses: []net.IP{
			net.ParseIP("127.0.0.1"),
			t.getPublicIP(),
		},
		DNSNames:              []string{"localhost"},
		BasicConstraintsValid: true,
	}

	derBytes, err := x509.CreateCertificate(rand.Reader, &template, &template, &priv.PublicKey, priv)

	certOut, err := os.CreateTemp("", "cert_*.pem")
	if err != nil {
		slog.Error("failed to create temp cert file", "error", err)
		os.Exit(1)
	}

	err = pem.Encode(certOut, &pem.Block{Type: "CERTIFICATE", Bytes: derBytes})
	if err != nil {
		slog.Error("failed to pem encode cert", "error", err)
		os.Exit(1)
	}

	if err := certOut.Close(); err != nil {
		slog.Error("failed to close cert file", "error", err)
		os.Exit(1)
	}

	keyOut, err := os.CreateTemp("", "key_*.pem")
	if err != nil {
		slog.Error("failed to create temp key file", "error", err)
		os.Exit(1)
	}

	privBytes, err := x509.MarshalPKCS8PrivateKey(priv)
	if err != nil {
		slog.Error("unable to marshal ECDSA private key", "error", err)
		os.Exit(1)
	}

	err = pem.Encode(keyOut, &pem.Block{Type: "PRIVATE KEY", Bytes: privBytes})
	if err != nil {
		slog.Error("failed to pem encode key", "error", err)
		os.Exit(1)
	}

	if err := keyOut.Close(); err != nil {
		slog.Error("failed to close key file", "error", err)
		os.Exit(1)
	}

	cert, err := tls.LoadX509KeyPair(certOut.Name(), keyOut.Name())
	if err != nil {
		slog.Error("failed to load x509 key pair", "error", err)
		os.Exit(1)
	}

	tlsConfig := &tls.Config{
		Certificates:       []tls.Certificate{cert},
		InsecureSkipVerify: true,
	}

	t.TerminatorActorChan <- TerminatorMessage{
		Type: AddCert,
		Value: TlsKeyPair{
			CertPath: certOut.Name(),
			KeyPath:  keyOut.Name(),
		},
	}
	return tlsConfig, nil
}

func (t *Terminator) handleCancel() {
	for range t.ctx.Done() {
		slog.Info("cleaning up generated TLS key pairs...")

		toDelete := make(chan []TlsKeyPair)
		t.TerminatorActorChan <- TerminatorMessage{
			Type:     GetCerts,
			Response: toDelete,
		}
		certs := <-toDelete

		for _, c := range certs {
			t.TerminatorActorChan <- TerminatorMessage{
				Type:  RemoveCert,
				Value: TlsKeyPair{CertPath: c.CertPath, KeyPath: c.KeyPath},
			}
		}
	}
}

func (t *Terminator) getPublicIP() net.IP {
	client := retryablehttp.NewClient()
	client.RetryMax = 3
	client.RetryWaitMin = 1 * time.Second
	client.RetryWaitMax = 15 * time.Second
	client.Logger = nil
	req, err := retryablehttp.NewRequest("GET", "https://ifconfig.me/ip", nil)
	if err != nil {
		slog.Error("failed to get public ip", "error", err)
		return net.ParseIP("127.0.0.1")
	}

	resp, err := client.Do(req)
	if err != nil {
		slog.Error("failed to send request", "error", err)
		return net.ParseIP("127.0.0.1")
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		slog.Error("failed to read response", "error", err)
		return net.ParseIP("127.0.0.1")
	}

	return net.ParseIP(string(data))
}
