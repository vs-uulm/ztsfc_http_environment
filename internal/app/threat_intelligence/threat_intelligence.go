package threat_intelligence

import (
    "net/http"
    "io/ioutil"

    logger "github.com/vs-uulm/ztsfc_http_logger"
)

func handleFlowAlert(w http.ResponseWriter, req *http.Request) {
    alertBytes, err := ioutil.ReadAll(req.Body)
    if err != nil {
        // TODO: Better error handling
        return
    }

    //alert := ioutil.NopCloser(bytes.NewBuffer(body))

    fmt.Printf("Alert: %s", string(alertBytes))
}

func runThreatIntelligence(sysLogger *logger.Logger) error {
    http.HandleFunc("/handleFlowAlert", handleFlowAlert)

    web_server := http.Server{
        Addr: ":8080",
    }

    err := web_server.ListenAndServe()
    if err != {
        return fmt.Errorf("threat_intelligence: runThreatIntelligence(): %v", err)
    }

    return nil
}
