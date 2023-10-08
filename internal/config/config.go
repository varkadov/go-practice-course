package config

import (
	"flag"
	"log"

	"github.com/caarlos0/env/v6"
)

type AgentConfig struct {
	Addr           string
	PollInterval   int
	ReportInterval int
}

type ServerConfig struct {
	Addr            string
	StoreInterval   int
	FileStoragePath string
	Restore         bool
}

func NewAgentConfig() *AgentConfig {
	var (
		config struct {
			Addr           *string `env:"ADDRESS"`
			PollInterval   *int    `env:"POLL_INTERVAL"`
			ReportInterval *int    `env:"REPORT_INTERVAL"`
		}
		addr           string
		pollInterval   int
		reportInterval int
	)

	if err := env.Parse(&config); err != nil {
		log.Fatal(err)
	}

	flag.StringVar(&addr, "a", "localhost:8080", "Server address")
	flag.IntVar(&pollInterval, "p", 2, "Pool interval")
	flag.IntVar(&reportInterval, "r", 10, "Report interval")
	flag.Parse()

	if config.Addr == nil {
		config.Addr = &addr
	}
	if config.PollInterval == nil {
		config.PollInterval = &pollInterval
	}
	if config.ReportInterval == nil {
		config.ReportInterval = &reportInterval
	}

	return &AgentConfig{
		Addr:           *config.Addr,
		PollInterval:   *config.PollInterval,
		ReportInterval: *config.ReportInterval,
	}
}

func NewServerConfig() *ServerConfig {
	var (
		config struct {
			Addr            *string `env:"ADDRESS"`
			StoreInterval   *int    `env:"STORE_INTERVAL"`
			FileStoragePath *string `env:"FILE_STORAGE_PATH"`
			Restore         *bool   `env:"RESTORE"`
		}
		addr            string
		storeInterval   int
		fileStoragePath string
		restore         bool
	)

	if err := env.Parse(&config); err != nil {
		log.Fatal(err)
	}

	flag.StringVar(&addr, "a", "localhost:8080", "Server address")
	flag.IntVar(&storeInterval, "i", 300, "Store interval")
	flag.StringVar(&fileStoragePath, "f", "/tmp/metrics-db.json", "File storage path")
	flag.BoolVar(&restore, "r", true, "Restore previous stored file")
	flag.Parse()

	if config.Addr == nil {
		config.Addr = &addr
	}
	if config.StoreInterval == nil {
		config.StoreInterval = &storeInterval
	}
	if config.FileStoragePath == nil {
		config.FileStoragePath = &fileStoragePath
	}
	if config.Restore == nil {
		config.Restore = &restore
	}

	return &ServerConfig{
		Addr:            *config.Addr,
		StoreInterval:   *config.StoreInterval,
		FileStoragePath: *config.FileStoragePath,
		Restore:         *config.Restore,
	}
}
