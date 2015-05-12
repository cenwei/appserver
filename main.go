package main

import (
	"net/http"
	"os"

	"github.com/hprose/hprose-go/hprose"
	"github.com/sharelog/appserver/log"
)

var logLevelMapping = map[string]log.Level{
	"trace": log.LevelTrace,
	"debug": log.LevelDebug,
	"error": log.LevelError,
	"fatal": log.LevelFatal,
}

func main() {
	// read config from os.Args
	if len(os.Args) <= 1 {
		log.Fatalf("should provide config file path via os args")
	}
	configPath := os.Args[1]
	config := Configuration{}
	if err := ReadConfig(&config, configPath); err != nil {
		log.Fatalf("cant read config: %v", err)
	}

	log.Tracef("set log level to %s:%v", config.LOG.Level, logLevelMapping[config.LOG.Level])
	log.Tracef("listen server on %v", config.HTTP.Host)

	// set log level
	log.SetLogLevel(logLevelMapping[config.LOG.Level])

	// initialize hprose service
	handler := hprose.NewHttpService()
	handler.AddMethods(&publicServices{})
	if logLevelMapping[config.LOG.Level] < log.LevelError {
		handler.GetEnabled = true
		handler.DebugEnabled = true
	}

	// start server
	if err := http.ListenAndServe(config.HTTP.Host, handler); err != nil {
		log.Fatalf("cant listen and serve http: %v", err)
	}
}
