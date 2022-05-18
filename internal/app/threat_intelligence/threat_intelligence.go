package threat_intelligence

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net"
	"net/http"

	"github.com/vs-uulm/ztsfc_http_pip/internal/app/config"
	"github.com/vs-uulm/ztsfc_http_pip/internal/app/device"
	"github.com/vs-uulm/ztsfc_http_pip/internal/app/jsonlogsender"
	"github.com/vs-uulm/ztsfc_http_pip/internal/app/system"

	rattr "github.com/vs-uulm/ztsfc_http_attributes"
)

var logSender *jsonlogsender.JSONLogSender

type flowAlert struct {
	TimeReceived  string `json:"TimeReceived"`
	FlowDirection uint32 `json:"FlowDirection"`
	TimeFlowStart string `json:"TimeFlowStart"`
	TimeFlowEnd   string `json:"TimeFlowEnd"`
	Bytes         string `json:"Bytes"`
	Packets       string `json:"Packets"`
	SrcAddr       string `json:"SrcAddr"`
	DstAddr       string `json:"DstAddr"`
	Etype         uint32 `json:"Etype"`
	Proto         uint32 `json:"Proto"`
	SrcPort       uint32 `json:"SrcPort"`
	DstPort       uint32 `json:"DstPort"`
	InIf          uint32 `json:"InIf"`
	OutIf         uint32 `json:"OutIf"`
	IPTTL         uint32 `json:"IPTTL"`
	TCPFlags      uint32 `json:"TCPFlags"`
	RemoteAddr    string `json:"RemoteAddr"`
}

func handleFlowAlert(w http.ResponseWriter, req *http.Request) {
	var alert flowAlert
	err := json.NewDecoder(req.Body).Decode(&alert)
	if err != nil {
		config.SysLogger.Errorf("threat_intelligence: runThreatIntelligence(): handleFlowAlert(): %s", err.Error())
		return
	}

	// Direct Reaction
	addrIP, err := convertAddrFromStringToIP(alert.SrcAddr)
	if err != nil {
		config.SysLogger.Errorf("threat_intelligence: runThreatIntelligence(): handleFlowAlert(): %s", err.Error())
		return
	}

	affectedDevice := rattr.FindDeviceByIPInIDMap(config.SysLogger, addrIP.String(), device.DevicesByID)
	if affectedDevice == nil {
		config.SysLogger.Infof("threat_intelligence: runThreatIntelligence(): handleFlowAlert(): exported suspicious IP "+
			"'%s' is currently not assigned to a managed device\n",
			addrIP.String())
	} else {
		affectedDevice.Revoked = true
		config.SysLogger.Infof("threat_intelligence: runThreatIntelligence(): handleFlowAlert(): exported suspicious IP "+
			"'%s' belongs to managed device '%s' that is now revoked\n",
			addrIP.String(), affectedDevice.DeviceID)
	}

	// Indirect Reaction
	system.System.ThreatLevel = 1

	err = logSender.SendThreatLevel("ztsfc_pip", system.System.ThreatLevel)
	if err != nil {
		config.SysLogger.Errorf("threat_intelligence: runThreatIntelligence(): handleFlowAlert(): unable to send the Treat Level value: %s", err.Error())
		return
	}

}

func RunThreatIntelligence() error {
	var err error
	logSender, err = jsonlogsender.New(config.Config.Pip.LoggingHook.HookURL)
	if err != nil {
		return err
	}

	err = logSender.SendThreatLevel("ztsfc_pip", system.System.ThreatLevel)
	if err != nil {
		return fmt.Errorf("RunThreatIntelligence(): unable to send the system threat level: %s", err.Error())
	}

	http.HandleFunc("/handleFlowAlert", handleFlowAlert)

	web_server := http.Server{
		Addr: ":8080",
	}

	err = web_server.ListenAndServe()
	if err != nil {
		return fmt.Errorf("threat_intelligence: runThreatIntelligence(): %s", err.Error())
	}

	return nil
}

// CONVENIENCE TOOLS
func convertAddrFromStringToIP(addr string) (net.IP, error) {
	addrBytes, err := base64.StdEncoding.DecodeString(addr)
	if err != nil {
		return nil, fmt.Errorf("convertAddrFromStringToIP: error decoding alert from flow exporter: %s", err.Error())
	}

	addrIP := net.IP(addrBytes)

	return addrIP, nil
}
