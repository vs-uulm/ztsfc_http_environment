package user

import (
    rattr "github.com/vs-uulm/ztsfc_http_attributes"
)

var (
    UserByID = make(map[string]*rattr.User)
)
