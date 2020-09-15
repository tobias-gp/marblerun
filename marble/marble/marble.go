package marble

import (
	"context"
	"crypto/ed25519"
	"crypto/rand"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"math"
	"math/big"
	"net"
	"os"
	"path/filepath"
	"time"

	"github.com/edgelesssys/coordinator/coordinator/quote"
	"github.com/edgelesssys/coordinator/coordinator/rpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// EdgCoordinatorAddr is a required env variable with Coordinator addr
const EdgCoordinatorAddr string = "EDG_COORDINATOR_ADDR"

// EdgMarbleType is a required env variable with type of this marble
const EdgMarbleType string = "EDG_MARBLE_TYPE"

// EdgMarbleCert is the env variable used to store the cert signed by the Coordinator
const EdgMarbleCert string = "EDG_MARBLE_CERT"

// EdgRootCA is the env variable used to store the Coordinator's RootCA
const EdgRootCA string = "EDG_ROOT_CA"

// EdgMarblePrivKey is the env variable used to store the private key for the cert
const EdgMarblePrivKey string = "EDG_MARBLE_PRIV_KEY"

// TODO: Create a central place where all certificate information is managed
// TLS Cert orgName
const orgName string = "Edgeless Systems GmbH"

// Signature for main function
type mainFunc func(int, []string, []string) int

// Authenticator holds the information for authenticating with the Coordinator
type Authenticator struct {
	commonName string
	orgName    string
	privk      ed25519.PrivateKey
	pubk       ed25519.PublicKey
	initCert   *x509.Certificate
	csr        *x509.CertificateRequest
	quote      []byte
	qi         quote.Issuer
	marbleCert *x509.Certificate
	rootCA     *x509.Certificate
	params     *rpc.Parameters
}

// NewAuthenticator creates a new Authenticator instance
func NewAuthenticator(orgName string, commonName string, qi quote.Issuer) (*Authenticator, error) {
	a := &Authenticator{
		commonName: commonName,
		orgName:    orgName,
		qi:         qi,
	}
	if err := a.generateCert(); err != nil {
		return nil, err
	}
	return a, nil
}

// loadTLSCreddentials builds a TLS config from the Authenticator's self-signed certificate and the Coordinator's RootCA
func loadTLSCredentials(a *Authenticator) (credentials.TransportCredentials, error) {
	clientCert, err := a.getTLSCertificate()
	if err != nil {
		return nil, fmt.Errorf("failed to get Marble self-signed x509 certificate")
	}
	tlsConfig := &tls.Config{
		Certificates:       []tls.Certificate{*clientCert},
		InsecureSkipVerify: true,
	}
	return credentials.NewTLS(tlsConfig), nil
}

// getTLSCertificate creates a TLS certificate for the Marbles self-signed x509 certificate
func (a *Authenticator) getTLSCertificate() (*tls.Certificate, error) {
	return tlsCertFromDER(a.initCert.Raw, a.privk), nil
}

// generateSerial returns a random serialNumber
func generateSerial() (*big.Int, error) {
	serialNumberLimit := new(big.Int).Lsh(big.NewInt(1), 128)
	return rand.Int(rand.Reader, serialNumberLimit)
}

// generateCert generates a new self-signed certificate associated key-pair
func (a *Authenticator) generateCert() error {

	// code (including generateSerial()) adapted from golang.org/src/crypto/tls/generate_cert.go
	pubk, privk, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		return err
	}
	notBefore := time.Now()
	notAfter := notBefore.Add(math.MaxInt64)

	serialNumber, err := generateSerial()
	if err != nil {
		return err
	}

	// TODO: what else do we need to set here?
	// Do we need x509.KeyUsageKeyEncipherment?
	template := x509.Certificate{
		SerialNumber: serialNumber,
		Subject: pkix.Name{
			Organization: []string{a.orgName},
			CommonName:   a.commonName,
		},
		NotBefore: notBefore,
		NotAfter:  notAfter,

		KeyUsage:              x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth, x509.ExtKeyUsageClientAuth},
		BasicConstraintsValid: false,
		IsCA:                  true,
	}

	certRaw, err := x509.CreateCertificate(rand.Reader, &template, &template, pubk, privk)
	if err != nil {
		return err
	}
	cert, err := x509.ParseCertificate(certRaw)
	if err != nil {
		return err
	}
	quote, err := a.qi.Issue(certRaw)
	if err != nil {
		return err
	}
	a.pubk = pubk
	a.privk = privk
	a.quote = quote
	a.initCert = cert
	return nil
}

func (a *Authenticator) generateCSR() error {
	template := x509.CertificateRequest{
		Subject: pkix.Name{
			Organization: []string{a.orgName},
			CommonName:   a.commonName,
		},
		PublicKey: a.pubk,
		// TODO: Add proper AltNames here: AB #172
		DNSNames:    []string{"localhost"},
		IPAddresses: []net.IP{net.IPv4(127, 0, 0, 1), net.IPv6loopback},
	}
	csrRaw, err := x509.CreateCertificateRequest(rand.Reader, &template, a.privk)
	if err != nil {
		return err
	}
	csr, err := x509.ParseCertificateRequest(csrRaw)
	if err != nil {
		return err
	}
	a.csr = csr
	return nil
}

// tlsCertFromDER converts a certificate from raw DER representation to a tls.Certificate
func tlsCertFromDER(certDER []byte, privk interface{}) *tls.Certificate {
	return &tls.Certificate{Certificate: [][]byte{certDER}, PrivateKey: privk}
}

// PreMain is supposed to run before the App's actual main and authenticate with the Coordinator
func PreMain(a *Authenticator, main mainFunc) (*x509.Certificate, *rpc.Parameters, error) {
	// get env variables
	coordAddr := os.Getenv(EdgCoordinatorAddr)
	if len(coordAddr) == 0 {
		return nil, nil, fmt.Errorf("environment variable not set: %v", EdgCoordinatorAddr)
	}

	marbleType := os.Getenv(EdgMarbleType)
	if len(marbleType) == 0 {
		return nil, nil, fmt.Errorf("environment variable not set: %v", EdgMarbleType)
	}

	// load TLS Credentials
	tlsCredentials, err := loadTLSCredentials(a)
	if err != nil {
		return nil, nil, err
	}

	// initiate grpc connection to Coordinator
	cc, err := grpc.Dial(coordAddr, grpc.WithTransportCredentials(tlsCredentials))

	if err != nil {
		return nil, nil, err
	}

	defer cc.Close()

	// generate CSR
	if err := a.generateCSR(); err != nil {
		return nil, nil, err
	}

	// authenticate with Coordinator
	c := rpc.NewMarbleClient(cc)
	req := &rpc.ActivationReq{
		CSR:        a.csr.Raw,
		MarbleType: marbleType,
		Quote:      a.quote,
	}

	activationResp, err := c.Activate(context.Background(), req)
	if err != nil {
		return nil, nil, err
	}
	newCert, err := x509.ParseCertificate(activationResp.GetCertificate())
	if err != nil {
		return nil, nil, err
	}
	a.marbleCert = newCert
	rootCA, err := x509.ParseCertificate(activationResp.GetRootCA())
	if err != nil {
		return nil, nil, err
	}
	a.rootCA = rootCA
	a.params = activationResp.GetParameters()

	// Store certificate in environment and file system
	pemCert := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: a.marbleCert.Raw})
	os.Setenv(EdgMarbleCert, string(pemCert))
	certFile, err := ioutil.TempFile("", "*.pem")
	if err != nil {
		return nil, nil, err
	}
	certFilename := certFile.Name()
	_, err = certFile.Write(pemCert)
	if err != nil {
		return nil, nil, err
	}
	certFile.Close()
	defer os.Remove(certFilename)

	// Store RootCA in environment
	pemRootCA := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: a.rootCA.Raw})
	os.Setenv(EdgRootCA, string(pemRootCA))

	// Store private key in environment
	privKeyPKCS8, err := x509.MarshalPKCS8PrivateKey(a.privk)
	if err != nil {
		return nil, nil, err
	}
	pemKey := pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: privKeyPKCS8})
	os.Setenv(EdgMarblePrivKey, string(pemKey))

	// Store files in file system
	for path, content := range a.params.Files {
		os.MkdirAll(filepath.Dir(path), os.ModePerm)
		err := ioutil.WriteFile(path, content, 0600)
		if err != nil {
			return nil, nil, err
		}
	}

	// // Set environment variables
	for key, value := range a.params.Env {
		os.Setenv(key, value)
	}

	// call main with args
	argv := a.params.Argv
	argc := len(argv)
	env := os.Environ()
	status := main(argc, argv, env)
	if status != 0 {
		return nil, nil, fmt.Errorf("main function returned error code: %v", status)
	}
	return a.marbleCert, a.params, nil

}