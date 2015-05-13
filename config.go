package main

import "github.com/scalingdata/gcfg"

// Configuration is appServer's configuration
type Configuration struct {
	HTTP struct {
		Host     string
		SSLRoute string `gcfg:"route-ssl"`
	}
	Security struct {
		XXTEAKey string `gcfg:"xxtea-key"`
	}
	LOG struct {
		Level string
	}
}

// ReadConfig reads a configuration from file specified by path
func ReadConfig(config *Configuration, path string) error {
	if err := gcfg.ReadFileInto(config, path); err != nil {
		return err
	}
	if config.HTTP.Host == "" {
		config.HTTP.Host = ":8888"
	}
	return nil
}
