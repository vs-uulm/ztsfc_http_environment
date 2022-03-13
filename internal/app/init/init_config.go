package init

import (
    "fmt"

    logger "github.com/vs-uulm/ztsfc_http_logger"
)


func InitConfig(sysLogger *logger.Logger) error {
    if err := initThreatIntelligence(sysLogger); err != nil {
        return fmt.Errorf("init: InitConfig(): %v", err)
    }

    return nil
}
