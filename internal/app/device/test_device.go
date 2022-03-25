package device

import (
    rattr "github.com/vs-uulm/ztsfc_http_attributes"
)

func LoadTestDevices() {
    m1MacMini, _ := rattr.NewDevice("M1 Mac Mini", "", false)
    DevicesByID[m1MacMini.DeviceID] = m1MacMini
}
