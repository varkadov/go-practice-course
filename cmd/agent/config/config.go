package config

import (
	"flag"
	"github.com/caarlos0/env/v6"
	"log"
)

type Config struct {
	Addr           string `env:"ADDRESS"`
	PollInterval   int    `env:"POLL_INTERVAL"`
	ReportInterval int    `env:"REPORT_INTERVAL"`
}

func NewConfig() *Config {
	var (
		config         Config
		Addr           string
		PollInterval   int
		ReportInterval int
	)

	if err := env.Parse(&config); err != nil {
		log.Fatal(err)
	}

	flag.StringVar(&Addr, "a", "localhost:8080", "Server address")
	flag.IntVar(&PollInterval, "p", 2, "Pool Interval")
	flag.IntVar(&ReportInterval, "r", 10, "Report interval")
	flag.Parse()

	if config.Addr == "" {
		config.Addr = Addr
	}
	if config.PollInterval == 0 {
		config.PollInterval = PollInterval
	}
	if config.ReportInterval == 0 {
		config.ReportInterval = ReportInterval
	}

	return &config
}
