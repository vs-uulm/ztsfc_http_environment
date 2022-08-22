package user

import (
    rattr "github.com/vs-uulm/ztsfc_http_attributes"
)

func LoadTestUser() {
    reinerCalumny := rattr.NewUser("reiner.calumny@ztsfc.com", 0, 8, 16, []string{"service1.testbed.informatik.uni-ulm.de"})
    UserByID[reinerCalumny.UserID] = reinerCalumny
}
