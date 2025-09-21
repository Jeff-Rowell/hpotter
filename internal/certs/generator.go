package certs

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"math/big"
	"net"
	"os"
	"path/filepath"
	"time"
)

type CertificatePaths struct {
	CertPath string
	KeyPath  string
}

func GenerateSelfSignedCert() (*CertificatePaths, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, fmt.Errorf("failed to generate private key: %v", err)
	}

	template := x509.Certificate{
		SerialNumber: big.NewInt(1),
		Subject: pkix.Name{
			Organization:  []string{"Apache Webserver"},
			Country:       []string{"US"},
			Province:      []string{""},
			Locality:      []string{""},
			StreetAddress: []string{""},
			PostalCode:    []string{""},
		},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().Add(365 * 24 * time.Hour),
		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
	}

	template.IPAddresses = append(template.IPAddresses, net.ParseIP("127.0.0.1"))
	template.IPAddresses = append(template.IPAddresses, net.ParseIP("::1"))
	template.DNSNames = append(template.DNSNames, "localhost")

	certDER, err := x509.CreateCertificate(rand.Reader, &template, &template, &privateKey.PublicKey, privateKey)
	if err != nil {
		return nil, fmt.Errorf("failed to create certificate: %v", err)
	}

	tmpDir := os.TempDir()

	certFile, err := os.CreateTemp(tmpDir, "hpotter-cert-*.crt")
	if err != nil {
		return nil, fmt.Errorf("failed to create temporary certificate file: %v", err)
	}
	defer certFile.Close()

	keyFile, err := os.CreateTemp(tmpDir, "hpotter-key-*.key")
	if err != nil {
		return nil, fmt.Errorf("failed to create temporary key file: %v", err)
	}
	defer keyFile.Close()

	certPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "CERTIFICATE",
		Bytes: certDER,
	})
	if _, err := certFile.Write(certPEM); err != nil {
		return nil, fmt.Errorf("failed to write certificate: %v", err)
	}

	privateKeyDER, err := x509.MarshalPKCS8PrivateKey(privateKey)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal private key: %v", err)
	}

	keyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: privateKeyDER,
	})
	if _, err := keyFile.Write(keyPEM); err != nil {
		return nil, fmt.Errorf("failed to write private key: %v", err)
	}

	return &CertificatePaths{
		CertPath: filepath.Clean(certFile.Name()),
		KeyPath:  filepath.Clean(keyFile.Name()),
	}, nil
}

func CleanupCertificateFiles(certPath, keyPath string) {
	if certPath != "" {
		if err := os.Remove(certPath); err != nil {
			fmt.Printf("warning: failed to remove certificate file %s: %v\n", certPath, err)
		}
	}
	if keyPath != "" {
		if err := os.Remove(keyPath); err != nil {
			fmt.Printf("warning: failed to remove key file %s: %v\n", keyPath, err)
		}
	}
}
