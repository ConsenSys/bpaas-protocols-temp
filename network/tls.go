package network

import (
	"crypto/tls"
)

//TODO: Utilise the new Confguration Management system
//TLSConfig - Get TLS configuration object to use for secure comms
func TLSConfig() (*tls.Config, error) {
	certificates := []tls.Certificate{}
	certificate, err := tls.LoadX509KeyPair("./tls/microservice.crt", "./tls/microservice.key")
	rootCertificate, err := tls.LoadX509KeyPair("./tls/root-ca.pem", "./tls/root-ca.key")

	if err != nil {
		return nil, err
	}

	certificates = append(certificates, certificate)
	certificates = append(certificates, rootCertificate)

	tlsConfig := tls.Config{
		Certificates:       certificates,
		ClientAuth:         tls.RequireAnyClientCert,
		InsecureSkipVerify: true,
	}
	return &tlsConfig, nil
}
