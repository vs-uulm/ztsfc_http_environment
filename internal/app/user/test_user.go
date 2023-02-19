package user

import (
    rattr "github.com/vs-uulm/ztsfc_http_attributes"
)

func LoadTestUser() {
    leo := rattr.NewUser("Leo", 0, 8, 16, []string{"wiki.bwnet.informatik.uni-ulm.de"})
    UserByID[leo.UserID] = leo
}
