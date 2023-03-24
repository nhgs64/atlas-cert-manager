package signer

import (
	"context"
	"crypto/x509"
	"encoding/pem"
	"math/big"
	"time"

	sampleissuerapi "github.com/cert-manager/sample-external-issuer/api/v1alpha1"
	"github.com/globalsign/hvclient"
)

var err error

type HealthChecker interface {
	Check() error
}

type HealthCheckerBuilder func(*sampleissuerapi.IssuerSpec, map[string][]byte) (HealthChecker, error)

type Signer interface {
	Sign([]byte) ([]byte, error)
}

type SignerBuilder func(*sampleissuerapi.IssuerSpec, map[string][]byte) (Signer, error)

func HVCAHealthCheckerFromIssuerAndSecretData(*sampleissuerapi.IssuerSpec, map[string][]byte) (HealthChecker, error) {
	return &hvcaSigner{}, nil
}

func HVCASignerFromIssuerAndSecretData(spec *sampleissuerapi.IssuerSpec, secret map[string][]byte) (Signer, error) {
	hvconfig := new(hvclient.Config)
	hvconfig.APIKey = string(secret["apikey"])
	hvconfig.APISecret = string(secret["apisecret"])
	hvconfig.URL = string(spec.URL)
	// decode pem to der expected by HVCA signer
	certDER, _ := pem.Decode(secret["cert"])
	keyDER, _ := pem.Decode(secret["certkey"])
	if hvconfig.TLSCert, err = x509.ParseCertificate(certDER.Bytes); err != nil {
		return nil, err
	}
	if hvconfig.TLSKey, err = x509.ParsePKCS1PrivateKey(keyDER.Bytes); err != nil {
		return nil, err
	}
	if err = hvconfig.Validate(); err != nil {
		return nil, err
	}
	return &hvcaSigner{config: hvconfig}, nil
}

type hvcaSigner struct {
	config *hvclient.Config
}

func (o *hvcaSigner) Check() error {
	return nil
}

func (o *hvcaSigner) Sign(csrBytes []byte) ([]byte, error) {
	ctx, cancel := context.WithCancel(context.Background())
	var clnt *hvclient.Client
	var serial *big.Int
	var info *hvclient.CertInfo
	defer cancel()
	if clnt, err = hvclient.NewClient(ctx, o.config); err != nil {
		return nil, err
	}
	// Parse the csr
	csr, err := parseCSR(csrBytes)
	if err != nil {
		return nil, err
	}

	var req = hvclient.Request{
		CSR: csr,
		Subject: &hvclient.DN{ // TODO: make this more flexible
			CommonName: csr.Subject.CommonName,
		},
		Validity: &hvclient.Validity{NotBefore: time.Now(), NotAfter: time.Unix(0, 0)}, //hardcoded max validity TODO
	}
	// Request cert
	if serial, err = clnt.CertificateRequest(ctx, &req); err != nil {
		return nil, err
	}
	// Retrieve cert
	if info, err = clnt.CertificateRetrieve(ctx, serial); err != nil {
		return nil, err
	}
	return pem.EncodeToMemory(&pem.Block{
		Type:  "CERTIFICATE",
		Bytes: info.X509.Raw,
	}), nil

	/*
		csr, err := parseCSR(csrBytes)
		if err != nil {
			return nil, err
		}
		key, err := parseKey(keyPEM)
		if err != nil {
			return nil, err
		}
		cert, err := parseCert(certPEM)
		if err != nil {
			return nil, err
		}
		ca := &CertificateAuthority{
			Certificate: cert,
			PrivateKey:  key,
			Backdate:    5 * time.Minute,
		}
		crtDER, err := ca.Sign(csr.Raw, PermissiveSigningPolicy{
			TTL: duration,
			Usages: []capi.KeyUsage{
				capi.UsageServerAuth,
			},
		})
		if err != nil {
			return nil, err
		}
		return pem.EncodeToMemory(&pem.Block{
			Type:  "CERTIFICATE",
			Bytes: crtDER,
		}), nil
		return nil, nil
	*/
}
