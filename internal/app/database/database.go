package database

import (
    rattr "github.com/vs-uulm/ztsfc_http_attributes"
)

var (
    Database DatabaseT
)

type DatabaseT struct {
    UserDB map[string]*rattr.User `yaml:"user"`
}
