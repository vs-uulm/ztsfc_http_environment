package init

import (
    "fmt"
    "strings"
    "github.com/vs-uulm/ztsfc_http_pip/internal/app/config"
)

func initThreatIntelligence() error {
    fields := ""
   
    // TODO: Check if the field make sense as well!
    if config.Config.ThreatIntelligence.ListenAddr == "" {
        fields += "listen_addr,"
    }

    if fields != "" {
        return fmt.Errorf("initThreatIntelligence(): in the section 'threat intelligence' the following required fields are missed: '%s'",
            strings.TrimSuffix(fields, ","))
    }

    return nil
}


