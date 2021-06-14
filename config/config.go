package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Broker    BrokerConfig
	Logstream LogstreamConfig
}

type BrokerConfig struct {
	Address  string
	Topics   []string
	ClientID string
}

type LogstreamConfig struct {
	Address   string
	Port      uint16
	AuthToken string
	Fields    map[string]string
}

// Load returns a `Config` struct read in from YAML configuration file at `configPath`
func Load(configPath string) Config {
	viper.SetConfigFile(configPath)

	// read in from file
	var config Config
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s\n", err)
	}

	// unmarshall into struct
	err := viper.Unmarshal(&config)
	if err != nil {
		log.Fatalf("Unable to decode config, %v\n", err)
	}

	return config
}
