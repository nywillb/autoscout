package main

import (
	"github.com/BurntSushi/toml"
)

type Config struct {
	Stats StatsConfig
}

type StatsConfig struct {
	URL      string
	Division string
}

func configure() Config {
	config := Config{}
	_, err := toml.DecodeFile("config.toml", &config)
	if err != nil {
		panic(err)
	}
	return config
}
