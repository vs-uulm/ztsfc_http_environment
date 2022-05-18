package init

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"

	gct "github.com/leobrada/golang_convenience_tools"
	"github.com/vs-uulm/ztsfc_http_pip/internal/app/config"
)

func initPip() error {
	fields := ""
	var err error

	if config.Config.Pip.ListenAddr == "" {
		fields += "listen_addr"
	}

	if config.Config.Pip.CertsPipAcceptsWhenShownByPdp == nil {
		fields += "certs_pip_accepts_when_shown_by_pdp"
	}

	if config.Config.Pip.CertShownByPipToPdp == "" {
		fields += "cert_shown_by_pip_to_pdp"
	}

	if config.Config.Pip.PrivkeyForCertShownByPipToPdp == "" {
		fields += "privkey_for_certs_shown_by_pip_to_pdp"
	}

	// Read CA certs used for signing client certs and are accepted by the PEP
	config.Config.Pip.CaCertPoolPipAcceptsFromPdp = x509.NewCertPool()
	for _, acceptedPdpCert := range config.Config.Pip.CertsPipAcceptsWhenShownByPdp {
		if err = gct.LoadCACertificate(acceptedPdpCert, config.Config.Pip.CaCertPoolPipAcceptsFromPdp); err != nil {
			return fmt.Errorf("initPipParams(): error loading certificates PIP accepts from PDP: %s", err.Error())
		}
	}

	config.Config.Pip.X509KeyPairShownByPipToPdp, err = gct.LoadX509KeyPair(config.Config.Pip.CertShownByPipToPdp,
		config.Config.Pip.PrivkeyForCertShownByPipToPdp)

	return err
}

// function unifies the loading of CA certificates for different components
func loadCACertificate(certfile string, componentName string, certPool *x509.CertPool) error {
	caRoot, err := ioutil.ReadFile(certfile)
	if err != nil {
		return fmt.Errorf("loadCACertificate(): Loading %s CA certificate from %s error: %v", componentName, certfile, err)
	}

	certPool.AppendCertsFromPEM(caRoot)
	return nil
}

// function unifies the loading of X509 key pairs for different components
func loadX509KeyPair(certfile, keyfile, componentName, certAttr string) (tls.Certificate, error) {
	keyPair, err := tls.LoadX509KeyPair(certfile, keyfile)
	if err != nil {
		return keyPair, fmt.Errorf("loadX509KeyPair(): critical error when loading %s X509KeyPair for %s from %s and %s: %v", certAttr, componentName, certfile, keyfile, err)
	}

	return keyPair, nil
}
