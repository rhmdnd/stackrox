package mtls

import (
	"crypto"
	"crypto/tls"
	"crypto/x509"

	"github.com/cloudflare/cfssl/helpers"
	cfsslSigner "github.com/cloudflare/cfssl/signer"
	"github.com/cloudflare/cfssl/signer/local"
	"github.com/pkg/errors"
	"github.com/stackrox/rox/pkg/sliceutils"
)

// CA represents a StackRox service certificate authority.
//go:generate mockgen-wrapper
type CA interface {
	// CheckProperties checks that the underlying CA certificate is indeed one
	// generated by StackRox (e.g., that it has the expected Common Name and other fields set
	// as expected).
	CheckProperties() error
	// Certificate returns the (public) CA certificate.
	Certificate() *x509.Certificate
	// PrivateKey returns the private key of the CA certificate.
	PrivateKey() crypto.PrivateKey
	// CertPEM returns the PEM-encoded form of the CA certificate.
	CertPEM() []byte
	// KeyPEM returns the PEM-encopded form of the CA key.
	KeyPEM() []byte
	// CertPool returns a certificate pool containing the single CA certificate.
	CertPool() *x509.CertPool
	// IssueCertForSubject issues a new certificate for the given subject from this CA.
	IssueCertForSubject(subj Subject, opts ...IssueCertOption) (*IssuedCert, error)
	// ValidateAndExtractSubject validates that the given certificate is a service certificate issued by this CA,
	// and extracts the subject information.
	ValidateAndExtractSubject(cert *x509.Certificate) (Subject, error)
}

type ca struct {
	certPEM, keyPEM []byte
	tlsCert         tls.Certificate
	signer          cfsslSigner.Signer
}

// LoadCAForSigning loads and instantiates a CA that can be used for signing
// from the given certificate and key.
// Note: this function does not verify that the given certificate is actually a valid
// StackRox service CA. To check for this, call `Validate()`.
func LoadCAForSigning(certPEM, keyPEM []byte) (CA, error) {
	return loadCA(certPEM, keyPEM, true)
}

// LoadCAForValidation loads a CA certificate to be used for certificate validation
// only. It can not be used to sign/issue new certificates, and invoking the respective
// methods will result in an error.
func LoadCAForValidation(certPEM []byte) (CA, error) {
	return loadCA(certPEM, nil, false)
}

func loadCA(certPEM, keyPEM []byte, forSigning bool) (CA, error) {
	ca := &ca{
		certPEM: sliceutils.ByteClone(certPEM),
		keyPEM:  sliceutils.ByteClone(keyPEM),
	}

	if forSigning {
		// certificate issuance and validation mode
		var err error
		ca.tlsCert, err = tls.X509KeyPair(certPEM, keyPEM)
		if err != nil {
			return nil, err
		}

		ca.tlsCert.Leaf, _ = x509.ParseCertificate(ca.tlsCert.Certificate[0])

		priv, err := helpers.ParsePrivateKeyPEM(keyPEM)
		if err != nil {
			return nil, err
		}

		ca.signer, err = local.NewSigner(priv, ca.tlsCert.Leaf, cfsslSigner.DefaultSigAlgo(priv), createSigningPolicy())
		if err != nil {
			return nil, err
		}
	} else {
		// CA without key, can be used for validation only, not cert issuing.
		var err error
		ca.tlsCert.Leaf, err = helpers.ParseCertificatePEM(certPEM)
		if err != nil {
			return nil, err
		}
		ca.tlsCert.Certificate = [][]byte{ca.tlsCert.Leaf.Raw}
	}

	return ca, nil
}

func (c *ca) CheckProperties() error {
	if !c.tlsCert.Leaf.IsCA {
		return errors.New("certificate is not valid as CA")
	}
	if cn := c.tlsCert.Leaf.Subject.CommonName; cn != ServiceCACommonName {
		return errors.Errorf("invalid certificate common name %q", cn)
	}
	return nil
}

func (c *ca) Certificate() *x509.Certificate {
	cert, _ := x509.ParseCertificate(c.tlsCert.Certificate[0])
	return cert
}

func (c *ca) PrivateKey() crypto.PrivateKey {
	return c.tlsCert.PrivateKey
}

func (c *ca) CertPEM() []byte {
	return sliceutils.ByteClone(c.certPEM)
}

func (c *ca) KeyPEM() []byte {
	return sliceutils.ByteClone(c.keyPEM)
}

func (c *ca) CertPool() *x509.CertPool {
	pool := x509.NewCertPool()
	pool.AddCert(c.tlsCert.Leaf)
	return pool
}

func (c *ca) IssueCertForSubject(subj Subject, opts ...IssueCertOption) (*IssuedCert, error) {
	if c.signer == nil {
		return nil, errors.New("CA is not set up for signing")
	}
	return issueNewCertFromSigner(subj, c.signer, opts)
}

func (c *ca) ValidateAndExtractSubject(cert *x509.Certificate) (Subject, error) {
	if _, err := cert.Verify(x509.VerifyOptions{Roots: c.CertPool()}); err != nil {
		return Subject{}, err
	}
	return SubjectFromCommonName(cert.Subject.CommonName), nil
}

// LoadDefaultCA loads the default StackRox Service CA from the default file paths.
// This function will read the files directly from the filesystem; if this does not
// meet your performance needs, you must implement your own caching.
func LoadDefaultCA() (CA, error) {
	_, caCertFileContents, _, err := readCA()
	if err != nil {
		return nil, err
	}
	caKeyFileContents, err := readCAKey()
	if err != nil {
		return nil, err
	}
	return LoadCAForSigning(caCertFileContents, caKeyFileContents)
}
