package terminator

// manageTlsKeyPairs is the actor function responsible for managing TLS
// key pairs. It takes a channel of TerminatorMessage as input and processes
// messages to add, remove, or get TLS key pairs.
func manageTlsKeyPairs(messages chan TerminatorMessage) {
	tlsKeyPairs := []TlsKeyPair{}

	for msg := range messages {
		switch msg.Type {
		case AddCert:
			tlsKeyPairs = append(tlsKeyPairs, msg.Value)
		case RemoveCert:
			for i, tls := range tlsKeyPairs {
				if tls.CertPath == msg.Value.CertPath &&
					tls.KeyPath == msg.Value.KeyPath {
					tlsKeyPairs = append(tlsKeyPairs[:i], tlsKeyPairs[i+1:]...)
					break
				}
			}
		case GetCerts:
			msg.Response <- tlsKeyPairs
		}
	}
}
