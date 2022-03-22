package device

func LoadTestDevices() {
    m1MacMini, _ := NewDevice("M1 Mac Mini", "", false)
    DevicesByID[m1MacMini.DeviceID] = m1MacMini
}
