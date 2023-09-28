package database

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"sync"

	"gopkg.in/yaml.v3"

	rattr "github.com/vs-uulm/ztsfc_http_attributes"
	"github.com/vs-uulm/ztsfc_http_pip/internal/app/config"
)

var (
	DatabaseFilePath string
	Database         DatabaseT
	WaitDatabaseList sync.Mutex
)

type DatabaseT struct {
	UserDB   map[string]*rattr.User   `yaml:"user"`
	DeviceDB map[string]*rattr.Device `yaml:"device"`
}

func UpdateDatabase() error {

	WaitDatabaseList.Lock()
	// Create a backup of the file
	backup, err := os.OpenFile(DatabaseFilePath+".bak", os.O_RDWR|os.O_CREATE, 0664)
	if err != nil {
		fmt.Errorf("database: UpdateDatabase(): error creating backup file: %v", err)
	}
	defer backup.Close()
	source, err := os.Open(DatabaseFilePath)
	if err != nil {
		fmt.Errorf("database: UpdateDatabase(): error opening database file: %v", err)
	}
	defer source.Close()
	_, err = io.Copy(backup, source)
	if err != nil {
		fmt.Errorf("database: UpdateDatabase(): error writing database backup: %v", err)
	}

	updatedDatabase, err := yaml.Marshal(Database)
	if err != nil {
		return fmt.Errorf("database: UpdateDatabase(): %v", err)
	}
	err = ioutil.WriteFile(DatabaseFilePath, updatedDatabase, 0664)
	if err != nil {
		return fmt.Errorf("database: UpdateDatabase(): %v", err)
	}
	WaitDatabaseList.Unlock()
	config.SysLogger.Infof("router: UpdateDatabase(): database has been successfully updated")
	return nil
}
