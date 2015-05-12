package main

import (
	"os"

	"github.com/sharelog/appserver/log"
)

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

	// TODO:
}
