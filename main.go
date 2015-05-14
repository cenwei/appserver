package main

import (
	"net/http"
	"os"

	"github.com/hprose/hprose-go/hprose"
	"github.com/sharelog/appserver/log"
	"github.com/sharelog/appserver/symmetric"
)

// headers
const (
	HeaderAccessToken = "Access-Token"
)

// session keys
const (
	SessionXXTEA = "XXTEA"
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
	log.Tracef("start normal service on %v", config.HTTP.Host)
	log.Tracef("start ssl service on %v%s", config.HTTP.Host, config.HTTP.SSLRoute)
	log.Tracef("......")

	// set log level
	log.SetLogLevel(logLevelMapping[config.LOG.Level])

	// initialize hprose service
	debug := logLevelMapping[config.LOG.Level] < log.LevelError
	handler := NewHTTPHproseService(&normalStub{}, debug)
	SSLHandler := NewHTTPHproseService(&sslStub{}, debug)
	handler.SetFilter(symmetricEncryption{
		getter:            nil, // TODO: need to pass session in
		symmetric:         symmetric.XXTEA{},
		headerAccessToken: HeaderAccessToken,
		sessionKey:        SessionXXTEA,
	})

	// register handlers for given pattern
	http.Handle(config.HTTP.SSLRoute, SSLHandler)
	http.Handle("/", handler)

	// start server
	if err := http.ListenAndServe(config.HTTP.Host, nil); err != nil {
		log.Fatalf("cant listen and serve http: %v", err)
	}
}

// NewHTTPHproseService initialize an http hprose service
func NewHTTPHproseService(stub interface{}, debug bool) *hprose.HttpService {
	service := hprose.NewHttpService()
	service.AddMethods(stub)
	service.GetEnabled = debug
	service.DebugEnabled = debug
	return service
}
