package config

import (
	"flag"
	"os"
)

type Config struct {
	Addr string
}

func NewConfig() *Config {
	addrEnv := os.Getenv("ADDRESS")
	addrFlag := flag.String("a", ":8080", "Server address")
	flag.Parse()

	if addrEnv != "" {
		*addrFlag = addrEnv
	}

	return &Config{
		Addr: *addrFlag,
	}
}
