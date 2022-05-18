package main

import (
	"flag"
	"log"

	yt "github.com/leobrada/yaml_tools"
	logger "github.com/vs-uulm/ztsfc_http_logger"
	"github.com/vs-uulm/ztsfc_http_pip/internal/app/config"
	"github.com/vs-uulm/ztsfc_http_pip/internal/app/device"
	confInit "github.com/vs-uulm/ztsfc_http_pip/internal/app/init"
	"github.com/vs-uulm/ztsfc_http_pip/internal/app/router"
	ti "github.com/vs-uulm/ztsfc_http_pip/internal/app/threat_intelligence"
)

//var (
//    SysLogger *logger.Logger
//)

func init() {
	var confFilePath string

	flag.StringVar(&confFilePath, "c", "./config/conf.yml", "Path to user defined yaml config file")
	flag.Parse()

	err := yt.LoadYamlFile(confFilePath, &config.Config)
	if err != nil {
		log.Fatalf("main: init(): could not load yaml file: %s", err.Error())
	}

	confInit.InitSysLoggerParams()
	config.SysLogger, err = logger.New(config.Config.SysLogger.LogFilePath,
		config.Config.SysLogger.LogLevel,
		config.Config.SysLogger.IfTextFormatter,
		logger.Fields{"type": "system"},
	)
	if err != nil {
		log.Fatalf("main: init(): could not initialize logger: %s", err.Error())
	}
	config.SysLogger.Debugf("loading pip configuration from '%s' - OK", confFilePath)

	if err = confInit.InitConfig(); err != nil {
		config.SysLogger.Fatalf("main: init(): could not initialize environment params: %s", err.Error())
	}

	// For testing
	device.LoadTestDevices()
}

func main() {
	var err error
	go ti.RunThreatIntelligence()

	//device.PrintDevices()

	pip := router.NewRouter()

	err = pip.ListenAndServeTLS()
	if err != nil {
		log.Fatalf("main: main(): listen and serve failed: %s", err.Error())
	}
}
