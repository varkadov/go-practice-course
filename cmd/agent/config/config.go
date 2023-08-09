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
	var cfg Config
	var (
		Addr           string
		PollInterval   int
		ReportInterval int
	)

	if err := env.Parse(&cfg); err != nil {
		log.Fatal(err)
	}

	flag.StringVar(&Addr, "a", "localhost:8080", "Server address")
	flag.IntVar(&PollInterval, "p", 2, "Pool Interval")
	flag.IntVar(&ReportInterval, "r", 10, "Report interval")
	flag.Parse()

	if cfg.Addr == "" {
		cfg.Addr = Addr
	}
	if cfg.PollInterval == 0 {
		cfg.PollInterval = PollInterval
	}
	if cfg.ReportInterval == 0 {
		cfg.ReportInterval = ReportInterval
	}

	return &cfg
}
