package main

import (
    "flag"
    "log"
    "github.com/vs-uulm/ztsfc_http_environment/internal/app/config"
    yt "github.com/leobrada/yaml_tools"
    logger "github.com/vs-uulm/ztsfc_http_logger"
    confInit "github.com/vs-uulm/ztsfc_http_environment/internal/app/init"
    ti "github.com/vs-uulm/ztsfc_http_environment/internal/app/threat_intelligence"
)

var (
    sysLogger *logger.Logger
)

func init() {
    var confFilePath string

    flag.StringVar(&confFilePath, "c", "./config/conf.yml", "Path to user defined yaml config file")
    flag.Parse()

    err := yt.LoadYamlFile(confFilePath, &config.Config)
    if err != nil {
        log.Fatalf("main: init(): could not load yaml file: %v", err)
    }

    confInit.InitSysLoggerParams()
    sysLogger, err = logger.New(config.Config.SysLogger.LogFilePath,
        config.Config.SysLogger.LogLevel,
        config.Config.SysLogger.IfTextFormatter,
        logger.Fields{"type": "system"},
    )
    if err != nil {
        log.Fatalf("main: init(): could not initialize logger: %v", err)
    }
    sysLogger.Debugf("loading logger configuration from %s - OK", confFilePath)

    if err = confInit.InitConfig(sysLogger); err != nil {
        sysLogger.Fatalf("main: init(): could not initialize Environment params: %v", err)
    }
}

func main() {
    go ti.RunThreatIntelligence(sysLogger)

    for {
    }
}
