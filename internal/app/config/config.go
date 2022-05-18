package config

import (
	"crypto/tls"
	"crypto/x509"

	logger "github.com/vs-uulm/ztsfc_http_logger"
)

var (
	Config    ConfigT
	SysLogger *logger.Logger
)

type ConfigT struct {
	SysLogger          sysLoggerT          `yaml:"system_logger"`
	Pip                PipT                `yaml:"pip"`
	ThreatIntelligence ThreatIntelligenceT `yaml:"threat_intelligence"`
}

type sysLoggerT struct {
	LogLevel        string `yaml:"system_logger_logging_level"`
	LogFilePath     string `yaml:"system_logger_destination"`
	IfTextFormatter string `yaml:"system_logger_format"`
}

type LoggingHookT struct {
	HookURL                        string   `yaml:"hook_url"`
	CertShownByPipToHook           string   `yaml:"cert_shown_by_pip_to_hook"`
	PrivkeyForCertShownByPipToHook string   `yaml:"privkey_for_cert_shown_by_pip_to_hook"`
	CertsPipAcceptsWhenShownByHook []string `yaml:"certs_pip_accepts_when_shown_by_hook"`
}

type PipT struct {
	ListenAddr                    string       `yaml:"listen_addr"`
	CertsPipAcceptsWhenShownByPdp []string     `yaml:"certs_pip_accepts_when_shown_by_pdp"`
	CertShownByPipToPdp           string       `yaml:"cert_shown_by_pip_to_pdp"`
	PrivkeyForCertShownByPipToPdp string       `yaml:"privkey_for_cert_shown_by_pip_to_pdp"`
	LoggingHook                   LoggingHookT `yaml:"logging_hook"`

	CaCertPoolPipAcceptsFromPdp  *x509.CertPool
	X509KeyPairShownByPipToPdp   tls.Certificate
	CaCertPoolPipAcceptsFromHook *x509.CertPool
	X509KeyPairShownByPipToHook  tls.Certificate
}

type ThreatIntelligenceT struct {
	ListenAddr string `yaml:"listen_addr"`
}
