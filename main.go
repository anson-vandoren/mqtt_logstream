package main

import (
	"os"
	"os/signal"
	"syscall"

	"ansonvandoren.com/mqtt_logstream/config"
)

const ENV_CONFIG_PATH = "MQTT_LOGSTREAM_CONFIG"

func main() {
	// load config from file specified in environment variable, if defined
	configPath := os.Getenv(ENV_CONFIG_PATH)
	if configPath == "" {
		configPath = "./config.yml"
	}
	configuration := config.Load(configPath)

	logstream := NewLogstreamConnection(configuration.Logstream)
	defer logstream.Close()

	// no particular need to keep a reference to MQTT listener in this simple example
	_ = NewMqttListener(configuration.Broker, logstream.Handler)

	// listen until terminated
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c
}
