package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"os"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/staking"
)

func main() {
	certificateBytes, privateKeyBytes, err := staking.NewCertAndKeyBytes()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to generate certificate and private key bytes: %v\n", err)
		os.Exit(1)
	}

	x509KeyPair, err := tls.X509KeyPair(certificateBytes, privateKeyBytes)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to parse X509 key pair: %v\n", err)
		os.Exit(1)
	}

	x509KeyPair.Leaf, err = x509.ParseCertificate(x509KeyPair.Certificate[0])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to parse x509 certificate: %v\n", err)
		os.Exit(1)
	}

	parsedCertificate := staking.CertificateFromX509(x509KeyPair.Leaf)
	nodeIDDerivedFromCertificate := ids.NodeIDFromCert(parsedCertificate)
	fmt.Println(nodeIDDerivedFromCertificate)
}
