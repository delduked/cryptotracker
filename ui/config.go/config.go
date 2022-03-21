package config

import (
	"os"
)

var Hostname = "http://" + apiHostname + ":" + apiPort

var apiHostname = getApiHostname()
var apiPort = getApiPort()

func getApiHostname() string {
	importHostname, ok := os.LookupEnv("API_HOSTNAME")
	if !ok {
		return "192.168.0.32"
	}
	return importHostname

}

func getApiPort() string {
	importHostname, ok := os.LookupEnv("API_PORT")
	if !ok {
		return "8000"
	}
	return importHostname

}
