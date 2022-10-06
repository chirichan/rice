package rice

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
)

var (
	// Cert is a self signed certificate
	Cert tls.Certificate
	// CertPool contains the self signed certificate
	CertPool *x509.CertPool
)

func CertInit(certPEM, keyPEM []byte) error {
	var err error
	Cert, err = tls.X509KeyPair(certPEM, keyPEM)
	if err != nil {
		return fmt.Errorf("failed to parse key pair: %w", err)
	}
	Cert.Leaf, err = x509.ParseCertificate(Cert.Certificate[0])
	if err != nil {
		return fmt.Errorf("failed to parse certificate: %w", err)
	}

	CertPool = x509.NewCertPool()
	CertPool.AddCert(Cert.Leaf)
	return nil
}
