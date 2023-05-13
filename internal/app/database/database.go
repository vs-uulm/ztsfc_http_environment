package database

import (
	"sync"

	rattr "github.com/vs-uulm/ztsfc_http_attributes"
)

var (
	Database         DatabaseT
	WaitDatabaseList sync.WaitGroup
)

type DatabaseT struct {
	UserDB   map[string]*rattr.User   `yaml:"user"`
	DeviceDB map[string]*rattr.Device `yaml:"device"`
}
