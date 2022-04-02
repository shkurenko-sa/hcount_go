package server

import (
	"log"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Host string `toml:"host"`
	Port string `toml:"port"`
}

func NewConfig() *Config {
	var cfg *Config
	_, err := toml.DecodeFile("config/env.toml", &cfg)
	if err != nil {
		log.Fatal(err)
	}
	return cfg
}
