package init

import (
	"fmt"
	"strings"

	"github.com/vs-uulm/ztsfc_http_pip/internal/app/config"
)

func InitGUILoggerParams() error {
	var err error
	fields := ""

	// if (config.Config.Pip.LoggingHook == config.LoggingHookT{}) {
	// 	return fmt.Errorf("init: InitGUILoggerParams(): in the section 'pip' a section 'logging_hook' is missed")
	// }

	if config.Config.Pip.LoggingHook.HookURL == "" {
		fields += "hook_url,"
	}

	if config.Config.Pip.LoggingHook.CertShownByPipToHook == "" {
		fields += "cert_shown_by_pip_to_hook,"
	}

	if config.Config.Pip.LoggingHook.PrivkeyForCertShownByPipToHook == "" {
		fields += "privkey_for_cert_shown_by_pip_to_hook,"
	}

	if len(config.Config.Pip.LoggingHook.CertsPipAcceptsWhenShownByHook) == 0 {
		fields += "certs_pip_accepts_when_shown_by_hook,"
	}

	if fields != "" {
		return fmt.Errorf("init: InitGUILoggerParams(): in the section 'pip.logging_hook' the following required fields are missed: '%s'", strings.TrimSuffix(fields, ","))
	}

	// Read CA certs used for signing client certs and are accepted by the logging hook
	for _, acceptedPepCert := range config.Config.Pip.LoggingHook.CertsPipAcceptsWhenShownByHook {
		if err = loadCACertificate(acceptedPepCert, "client", config.Config.Pip.CaCertPoolPipAcceptsFromHook); err != nil {
			return fmt.Errorf("init: InitGUILoggerParams(): error loading certificates PIP accepts from the logging hook: %s", err.Error())
		}
	}

	config.Config.Pip.X509KeyPairShownByPipToHook, err = loadX509KeyPair(config.Config.Pip.LoggingHook.CertShownByPipToHook,
		config.Config.Pip.LoggingHook.PrivkeyForCertShownByPipToHook, "Hook", "")
	if err != nil {
		return fmt.Errorf("init: InitGUILoggerParams(): error loading certificate pair PIP shows to logging hook: %s", err.Error())
	}

	return nil
}
