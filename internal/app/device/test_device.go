package device

import (
    rattr "github.com/vs-uulm/ztsfc_http_attributes"
)

func LoadTestDevices() {
    m1MacMini, _ := rattr.NewDevice("M1 Mac Mini", "", false)
    m1MacBookPro, _ := rattr.NewDevice("M1 MacBook Pro 16 Inch", "", false)
    x1Lenovo, _ := rattr.NewDevice("X1 Kali Linux", "", false)
    DevicesByID[m1MacMini.DeviceID] = m1MacMini
    DevicesByID[m1MacBookPro.DeviceID] = m1MacBookPro
    DevicesByID[x1Lenovo.DeviceID] = x1Lenovo
}
