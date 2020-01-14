package config

import (
	"github.com/BurntSushi/toml"
)

// Config represents the configuration in config.toml
type Config struct {
	Stats StatsConfig
}

// StatsConfig is a configuration for how to get the stats
type StatsConfig struct {
	Type        string
	URL         string
	Division    string
	TOAKey      string
	TOAOrigin   string
	TOAEventKey string
}

// Configure loads the config.toml file into memory
func Configure(file string) Config {
	config := Config{}
	_, err := toml.DecodeFile(file, &config)
	if err != nil {
		panic(err)
	}
	return config
}
