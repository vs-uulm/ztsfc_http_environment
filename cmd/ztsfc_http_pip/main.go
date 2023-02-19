package main

import (
    "flag"
    "log"
    "github.com/vs-uulm/ztsfc_http_pip/internal/app/router"
    "github.com/vs-uulm/ztsfc_http_pip/internal/app/config"
    "github.com/vs-uulm/ztsfc_http_pip/internal/app/database"
    //"github.com/vs-uulm/ztsfc_http_pip/internal/app/device"
    //"github.com/vs-uulm/ztsfc_http_pip/internal/app/user"
    yt "github.com/leobrada/yaml_tools"
    logger "github.com/vs-uulm/ztsfc_http_logger"
    confInit "github.com/vs-uulm/ztsfc_http_pip/internal/app/init"
    //ti "github.com/vs-uulm/ztsfc_http_pip/internal/app/threat_intelligence"
)

//var (
//    SysLogger *logger.Logger
//)

func init() {
    var confFilePath string
    var databaseFilePath string

    flag.StringVar(&confFilePath, "c", "./config/conf.yml", "Path to user defined yaml config file")
    flag.StringVar(&databaseFilePath, "d", "./database/database.yml", "Path to user defined yaml database file")
    flag.Parse()

    // Load config file
    err := yt.LoadYamlFile(confFilePath, &config.Config)
    if err != nil {
        log.Fatalf("main: init(): could not load yaml file: %v", err)
    }

    // Load database file
    err = yt.LoadYamlFile(databaseFilePath, &database.Database)
    if err != nil {
        log.Fatalf("main: init(): could not load yaml file: %v", err)
    }

    confInit.InitSysLoggerParams()
    config.SysLogger, err = logger.New(config.Config.SysLogger.LogFilePath,
        config.Config.SysLogger.LogLevel,
        config.Config.SysLogger.IfTextFormatter,
        logger.Fields{"type": "system"},
    )
    if err != nil {
        log.Fatalf("main: init(): could not initialize logger: %v", err)
    }
    config.SysLogger.Debugf("loading logger configuration from %s - OK", confFilePath)

    if err = confInit.InitConfig(); err != nil {
        config.SysLogger.Fatalf("main: init(): could not initialize Environment params: %v", err)
    }
}

func main() {
    //go ti.RunThreatIntelligence()
    pip := router.NewRouter()

    pip.ListenAndServeTLS()
}
