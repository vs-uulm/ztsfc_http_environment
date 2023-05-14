package init

import (
	"time"

	yt "github.com/leobrada/yaml_tools"
	"github.com/vs-uulm/ztsfc_http_pip/internal/app/config"
	"github.com/vs-uulm/ztsfc_http_pip/internal/app/database"
)

func InitDatabase(databaseFilePath string) error {

	// Databse is loaded in main.go: init()

	go reloadDatabase(databaseFilePath)

	return nil
}

func reloadDatabase(databaseFilePath string) {

	reloadInterval := time.Tick(1 * time.Minute)
	for range reloadInterval {
		// Load current state of database file
		database.WaitDatabaseList.Add(1)
		//database.Database.UserDB = make(map[string]*rattr.User)
		//database.Database.DeviceDB = make(map[string]*rattr.Device)
		err := yt.LoadYamlFile(databaseFilePath, &database.Database)
		if err != nil {
			config.SysLogger.Fatalf("main: init(): could not update database: %v: trying again in 1 minute", err)
		} else {
			config.SysLogger.Info("main: init(): InitDatabase(): reloadDatabase(): successfully reloaded database")
		}
		database.WaitDatabaseList.Add(-1)
	}

}
