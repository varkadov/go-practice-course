package config

import (
	"flag"
	"github.com/caarlos0/env/v6"
	"log"
)

type Config struct {
	Addr string `env:"ADDRESS"`
}

func NewConfig() *Config {
	var (
		config Config
		addr   string
	)

	if err := env.Parse(&config); err != nil {
		log.Fatal(err)
	}

	flag.StringVar(&addr, "a", ":8080", "Server address")
	flag.Parse()

	if config.Addr == "" {
		config.Addr = addr
	}

	return &config
}
