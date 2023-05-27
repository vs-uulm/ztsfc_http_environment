package router

import (
	"crypto/tls"
	"encoding/json"
	"net/http"

	"github.com/vs-uulm/ztsfc_http_pip/internal/app/config"
	"github.com/vs-uulm/ztsfc_http_pip/internal/app/database"

	//    "github.com/vs-uulm/ztsfc_http_pip/internal/app/device"
	//    "github.com/vs-uulm/ztsfc_http_pip/internal/app/user"
	"github.com/vs-uulm/ztsfc_http_pip/internal/app/system"

	rattr "github.com/vs-uulm/ztsfc_http_attributes"
)

const (
	// Request URIs for the API endpoint.
	getDeviceEndpoint           = "/get-device-attributes"
	updateDeviceEndpoint        = "/update-device-attributes"
	getUserEndpoint             = "/get-user-attributes"
	pushUserAttrUpdatesEndpoint = "/push-user-attr-updates"
	getSystemEndpoint           = "/get-system-attributes"
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
	http.HandleFunc(updateDeviceEndpoint, handleUpdateDeviceRequests)
	http.HandleFunc(getUserEndpoint, handleGetUserRequests)
	http.HandleFunc(pushUserAttrUpdatesEndpoint, handlePushUserAttrUpdateRequests)
	http.HandleFunc(getSystemEndpoint, handleGetSystemRequests)

	// Create HTTP frontend server
	router.frontend_server = &http.Server{
		Addr:      config.Config.Pip.ListenAddr,
		TLSConfig: router.frontend_tls_config,
	}

	return router
}

func handleGetDeviceRequests(w http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		config.SysLogger.Errorf("router: handleGetDeviceRequests(): PDP sent an unsupported HTTP request method")
		w.WriteHeader(405)
		return
	}

	q := req.URL.Query()

	dev := q.Get("device")
	if len(dev) == 0 {
		config.SysLogger.Infof("router: handleGetDeviceRequests(): get device request did not contain a device")
		w.WriteHeader(404)
		return
	}

	database.WaitDatabaseList.Wait()
	requestedDevice, ok := database.Database.DeviceDB[dev]
	if !ok {
		config.SysLogger.Infof("router: handleGetDeviceRequests(): PDP requested a device that does not exist in the DB")
		w.WriteHeader(404)
		return
	}

	config.SysLogger.Infof("router: handleGetDeviceRequests(): PDP requested the following device: %v", dev)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(requestedDevice)
}

func handleUpdateDeviceRequests(w http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		config.SysLogger.Errorf("router: handleUpdateDeviceRequests(): PDP sent an unsupported HTTP request method")
		w.WriteHeader(405)
		return
	}

	deviceUpdate := rattr.NewEmptyDevice()
	err := json.NewDecoder(req.Body).Decode(deviceUpdate)
	if err != nil {
		config.SysLogger.Errorf("router: handleUpdateDeviceRequests(): could not decode device update from PDP %v", err)
		return
	}

	database.WaitDatabaseList.Wait()
	deviceToUpdate, ok := database.Database.DeviceDB[deviceUpdate.DeviceID]
	if !ok {
		config.SysLogger.Errorf("router: handleUpdateDeviceRequests(): device '%s' PDP requested to update does not exist", deviceUpdate.DeviceID)
		return
	}

	deviceToUpdate.CurrentIP = deviceUpdate.CurrentIP
	config.SysLogger.Infof("router: handleGetDeviceRequests(): PDP updated the following device: %v", deviceToUpdate)
}

func handleGetUserRequests(w http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		config.SysLogger.Errorf("router: handleGetDeviceRequests(): PDP sent an unsupported HTTP request method")
		w.WriteHeader(405)
		return
	}

	q := req.URL.Query()

	usr := q.Get("user")
	if len(usr) == 0 {
		config.SysLogger.Infof("router: handleGetDeviceRequests(): get user request did not contain a user")
		w.WriteHeader(404)
		return
	}

	database.WaitDatabaseList.Wait()
	requestedUser, ok := database.Database.UserDB[usr]
	if !ok {
		config.SysLogger.Infof("router: handleGetDeviceRequests(): PDP requested a user that does not exist in the DB")
		w.WriteHeader(404)
		return
	}

	config.SysLogger.Infof("router: handleGetDeviceRequests(): PDP requested the following user: %v", usr)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(requestedUser)
}

// TODO: MUTEX for accessing the user object
func handlePushUserAttrUpdateRequests(w http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		// TODO: Check if its a PEP or PDP or whoever
		config.SysLogger.Errorf("router: handlePushUserAttrUpdateRequests(): PEP sent an unsupported HTTP request method")
		w.WriteHeader(405)
		return
	}

	q := req.URL.Query()

	usr := q.Get("user")
	if len(usr) == 0 {
		config.SysLogger.Infof("router: handlePushUserAttrUpdateRequests(): push user attribute update request did not contain a user")
		w.WriteHeader(404)
		return
	}

	failedAuthAttempts := q.Get("failed-auth-attempt")
	if len(failedAuthAttempts) != 0 {
		database.WaitDatabaseList.Wait()
		requestedUser, ok := database.Database.UserDB[usr]
		if !ok {
			config.SysLogger.Infof("router: handlePushUserAttrUpdateRequests(): user to update %s could not be found in User DB", usr)
			return
		}
		requestedUser.FailedAuthAttempts += 1
		database.UpdateDatabase()
		config.SysLogger.Infof("User %s has now %d failed authentication attemps", usr, requestedUser.FailedAuthAttempts)
	}

	successPWAuthentication := q.Get("success-auth-attempt")
	if len(successPWAuthentication) != 0 {
		database.WaitDatabaseList.Wait()
		requestedUser, ok := database.Database.UserDB[usr]
		if !ok {
			config.SysLogger.Infof("router: handlePushUserAttrUpdateRequests(): user to update %s could not be found in User DB", usr)
			return
		}
		requestedUser.FailedAuthAttempts = 0
		database.UpdateDatabase()
		config.SysLogger.Infof("User %s has now %d failed authentication attemps", usr, requestedUser.FailedAuthAttempts)
	}

	config.SysLogger.Debugf("User %s just got updated", usr)
}

func handleGetSystemRequests(w http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		config.SysLogger.Errorf("router: handleGetSystemRequests(): PDP sent an unsupported HTTP request method")
		w.WriteHeader(405)
		return
	}

	config.SysLogger.Infof("router: handleGetSystemRequests(): PDP requested the system attributes")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(system.System)
}

func (router *Router) ListenAndServeTLS() error {
	return router.frontend_server.ListenAndServeTLS("", "")
}
