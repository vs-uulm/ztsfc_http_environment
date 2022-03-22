package init

import (
    "fmt"
)


func InitConfig() error {
    if err := initThreatIntelligence(); err != nil {
        return fmt.Errorf("init: InitConfig(): %v", err)
    }

    if err := initPip(); err != nil {
        return fmt.Errorf("init: InitConfig(): %v", err)
    }

    return nil
}
