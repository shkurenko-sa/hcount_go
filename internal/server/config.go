package server

import (
	"log"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Host          string `toml:"host"`
	Port          string `toml:"port"`
	MongoHost     string `toml:"mongo_host"`
	MongoPort     string `toml:"mongo_port"`
	MongoUser     string `toml:"mongo_user"`
	MongoPassword string `toml:"mongo_password"`
	MongoDB       string `toml:"mongo_db"`
}

func NewConfig() *Config {
	var cfg *Config
	_, err := toml.DecodeFile("config/env.toml", &cfg)
	if err != nil {
		log.Fatal(err)
	}
	return cfg
}
