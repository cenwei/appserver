package main

import (
	"os"

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

	// set log level
	log.SetLogLevel(logLevelMapping[config.LOG.Level])
}
