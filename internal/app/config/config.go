package config

import (
    "crypto/x509"
    "crypto/tls"

    logger "github.com/vs-uulm/ztsfc_http_logger"
)

var (
    Config ConfigT
    SysLogger *logger.Logger
)

type ConfigT struct {
    SysLogger sysLoggerT `yaml:"system_logger"`
    Pip PipT `yaml:"pip"`
    ThreatIntelligence ThreatIntelligenceT `yaml:"threat_intelligence"`
}

type sysLoggerT struct {
    LogLevel string `yaml:"system_logger_logging_level"`
    LogFilePath string `yaml:"system_logger_destination"`
    IfTextFormatter string `yaml:"system_logger_format"`
}

type PipT struct {
    ListenAddr string `yaml:"listen_addr"`
    CertsPipAcceptsWhenShownByPdp []string `yaml:"certs_pip_accepts_when_shown_by_pdp"`
    CertShownByPipToPdp string  `yaml:"cert_shown_by_pip_to_pdp"`
    PrivkeyForCertShownByPipToPdp  string  `yaml:"privkey_for_cert_shown_by_pip_to_pdp"`

    CaCertPoolPipAcceptsFromPdp *x509.CertPool
    X509KeyPairShownByPipToPdp  tls.Certificate
}

type ThreatIntelligenceT struct {
    ListenAddr string `yaml:"listen_addr"`
}
