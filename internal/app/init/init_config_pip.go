package init

import (
    "fmt"
    "crypto/x509"

    "github.com/vs-uulm/ztsfc_http_pip/internal/app/config"
    gct "github.com/leobrada/golang_convenience_tools"
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
            return fmt.Errorf("initPipParams(): error loading certificates PIP accepts from PDP: %w", err)
        }
    }

    config.Config.Pip.X509KeyPairShownByPipToPdp, err = gct.LoadX509KeyPair(config.Config.Pip.CertShownByPipToPdp,
        config.Config.Pip.PrivkeyForCertShownByPipToPdp)

    return err
}
