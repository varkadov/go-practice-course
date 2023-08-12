package config

import (
	"flag"
	"log"

	"github.com/caarlos0/env/v6"
)

type Config struct {
	Addr           *string `env:"ADDRESS"`
	PollInterval   *int    `env:"POLL_INTERVAL"`
	ReportInterval *int    `env:"REPORT_INTERVAL"`
}

func NewConfig() *Config {
	var (
		config         Config
		addr           *string
		pollInterval   *int
		reportInterval *int
	)

	if err := env.Parse(&config); err != nil {
		log.Fatal(err)
	}

	flag.StringVar(addr, "a", "localhost:8080", "Server address")
	flag.IntVar(pollInterval, "p", 2, "Pool Interval")
	flag.IntVar(reportInterval, "r", 10, "Report interval")
	flag.Parse()

	if config.Addr == nil {
		config.Addr = addr
	}
	if config.PollInterval == nil {
		config.PollInterval = pollInterval
	}
	if config.ReportInterval == nil {
		config.ReportInterval = reportInterval
	}

	return &config
}
