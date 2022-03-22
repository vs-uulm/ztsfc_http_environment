package router

import (
	"crypto/tls"
	"encoding/json"
	"net/http"

    "github.com/vs-uulm/ztsfc_http_pip/internal/app/config"
    "github.com/vs-uulm/ztsfc_http_pip/internal/app/device"
)

const (
	// Request URIs for the API endpoint.
	getDeviceEndpoint = "/get-device-attributes"
)

type Router struct {
	frontend_tls_config *tls.Config
	frontend_server     *http.Server
}

func NewRouter() *Router {

	// Create new Router
	router := new(Router)

	// Create TLS config for frontend server
	router.frontend_tls_config = &tls.Config{
		Rand:                   nil,
		Time:                   nil,
		MinVersion:             tls.VersionTLS13,
		MaxVersion:             tls.VersionTLS13,
		SessionTicketsDisabled: true,
		Certificates:           []tls.Certificate{config.Config.Pip.X509KeyPairShownByPipToPdp},
		ClientAuth:             tls.RequireAndVerifyClientCert,
		ClientCAs:              config.Config.Pip.CaCertPoolPipAcceptsFromPdp,
	}

	// Create MUX server
    http.HandleFunc(getDeviceEndpoint, handleGetDeviceRequests)

	// Create HTTP frontend server
	router.frontend_server = &http.Server{
		Addr:      config.Config.Pip.ListenAddr,
		TLSConfig: router.frontend_tls_config,
	}

	return router
}

func handleGetDeviceRequests(w http.ResponseWriter, req *http.Request) {
    q := req.URL.Query()

    dev := q.Get("device");
    if  len(dev) == 0 {
        config.SysLogger.Infof("router: handleGetDeviceRequests(): get device request did not contain a device")
        w.WriteHeader(404)
        return
    }

    requestedDevice, ok := device.DevicesByID[dev]
    if !ok {
        config.SysLogger.Infof("router: handleGetDeviceRequests(): PDP requested a device that does not exist in the DB")
        w.WriteHeader(404)
        return
    }

    config.SysLogger.Infof("router: handleGetDeviceRequests(): PDP requested the following device: %v", requestedDevice)
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(requestedDevice)
}

func (router *Router) ListenAndServeTLS() error {
	return router.frontend_server.ListenAndServeTLS("", "")
}
