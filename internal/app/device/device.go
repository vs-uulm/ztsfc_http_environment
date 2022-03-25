package device

import (
    rattr "github.com/vs-uulm/ztsfc_http_attributes"
)

var (
    DevicesByID = make(map[string]*rattr.Device)
    DevicesByIP = make(map[string]*rattr.Device)
)

//type Device struct {
//    DeviceID string `json:"deviceID"`
//    CurrentIP string `json:"currentIP"`
//    Revoked bool `json:"revoked"`
//}
//
//func NewDevice(_deviceID, _currentIP string, _revoked bool) (*Device, error) {
//    newDevice := new(Device)
//    newDevice.DeviceID = _deviceID
//    newDevice.CurrentIP = _currentIP
//    newDevice.Revoked = _revoked
//    return newDevice, nil
//}
//
//func PrintDevices() {
//    for _, deviceObj := range DevicesByID {
//        config.SysLogger.Infof("%v\n", deviceObj)
//    }
//}
