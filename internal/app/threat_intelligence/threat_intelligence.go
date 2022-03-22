package threat_intelligence

import (
    "fmt"
    "net/http"
    "encoding/json"
    "net"
    "encoding/base64"

    "github.com/vs-uulm/ztsfc_http_pip/internal/app/config"
)

type flowAlert struct {
    TimeReceived  string `json:"TimeReceived"`
    FlowDirection uint32 `json:"FlowDirection"`
    TimeFlowStart string `json:"TimeFlowStart"`
    TimeFlowEnd   string `json:"TimeFlowEnd"`
    Bytes   string `json:"Bytes"`
    Packets string `json:"Packets"`
    SrcAddr string `json:"SrcAddr"`
    DstAddr string `json:"DstAddr"`
    Etype uint32 `json:"Etype"`
    Proto uint32 `json:"Proto"`
    SrcPort uint32 `json:"SrcPort"`
    DstPort uint32 `json:"DstPort"`
    InIf  uint32 `json:"InIf"`
    OutIf uint32 `json:"OutIf"`
    IPTTL uint32 `json:"IPTTL"`
    TCPFlags uint32 `json:"TCPFlags"`
    RemoteAddr string `json:"RemoteAddr"`
}

func handleFlowAlert(w http.ResponseWriter, req *http.Request) {
    var alert flowAlert
    err := json.NewDecoder(req.Body).Decode(&alert)
    if err != nil {
        config.SysLogger.Errorf("threat_intelligence: runThreatIntelligence(): handleFlowAlert(): %v\n", err)
        return
    }

    addrIP, err := convertAddrFromStringToIP(alert.SrcAddr)
    if err != nil {
        config.SysLogger.Errorf("threat_intelligence: runThreatIntelligence(): handleFlowAlert(): %v\n", err)
        return
    }

    config.SysLogger.Infof("threat_intelligence: runThreatIntelligence(): handleFlowAlert(): exported suspicious IP %s\n", addrIP.String())
}

func RunThreatIntelligence() error {
    http.HandleFunc("/handleFlowAlert", handleFlowAlert)

    web_server := http.Server{
        Addr: ":8080",
    }

    err := web_server.ListenAndServe()
    if err != nil {
        return fmt.Errorf("threat_intelligence: runThreatIntelligence(): %v", err)
    }

    return nil
}

// CONVENIENCE TOOLS
func convertAddrFromStringToIP(addr string) (net.IP, error) {
    addrBytes, err := base64.StdEncoding.DecodeString(addr)
    if err != nil {
        return nil, fmt.Errorf("convertAddrFromStringToIP: error decoding alert from flow exporter: %v", err)
    }

    addrIP := net.IP(addrBytes)

    return addrIP, nil
}
