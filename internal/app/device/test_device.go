package device

import (
    rattr "github.com/vs-uulm/ztsfc_http_attributes"
)

func LoadTestDevices() {
    m1MacMini, _ := rattr.NewDevice("M1 Mac Mini", "", false)
    m1MacBookPro, _ := rattr.NewDevice("M1 MacBook Pro 16 Inch", "", false)
    x1Lenovo, _ := rattr.NewDevice("X1 Kali Linux", "", false)
    scLegitimateDevice, _ := rattr.NewDevice("ZTSFC Testbed Web Client1", "", false)
    scLegitimateDevice2, _ := rattr.NewDevice("ZTSFC Testbed Web Client2", "", false)
    DevicesByID[m1MacMini.DeviceID] = m1MacMini
    DevicesByID[m1MacBookPro.DeviceID] = m1MacBookPro
    DevicesByID[x1Lenovo.DeviceID] = x1Lenovo
    DevicesByID[scLegitimateDevice.DeviceID] = scLegitimateDevice
    DevicesByID[scLegitimateDevice2.DeviceID] = scLegitimateDevice2
}
