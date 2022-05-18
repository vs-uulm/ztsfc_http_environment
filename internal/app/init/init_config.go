package init

import (
	"fmt"
)

func InitConfig() error {
	if err := initThreatIntelligence(); err != nil {
		return fmt.Errorf("init: InitConfig(): %s", err.Error())
	}

	if err := initPip(); err != nil {
		return fmt.Errorf("init: InitConfig(): %s", err.Error())
	}

	return nil
}
