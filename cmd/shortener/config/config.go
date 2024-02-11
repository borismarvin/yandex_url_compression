package config

import (
	"flag"
)

type Config struct {
	ServerAddress string
	BaseURL       string
}

func NewConfig() *Config {
	serverAddr := flag.String("a", "localhost:8888", "HTTP server start address")
	baseAddr := flag.String("b", "http://localhost:8888/", "Base address")
	flag.Parse()

	return &Config{
		ServerAddress: *serverAddr,
		BaseURL:       *baseAddr,
	}
}
