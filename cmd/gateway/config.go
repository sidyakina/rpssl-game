package main

import (
	"time"

	"github.com/caarlos0/env/v6"
)

type Config struct {
	HTTPPort         int64         `env:"HTTP_PORT"`
	HTTPWriteTimeout time.Duration `env:"HTTP_WRITE_TIMEOUT" envDefault:"5s"`
	HTTPReadTimeout  time.Duration `env:"HTTP_READ_TIMEOUT" envDefault:"5s"`

	GPPCConnectTimeout time.Duration `env:"GRPC_CONNECT_TIMEOUT" envDefault:"5s"`
	GRPCRequestTimeout time.Duration `env:"GRPC_REQUESTS_TIMEOUT" envDefault:"2s"`
	GameServiceURI     string        `env:"GAME_SERVICE_URI"`
}

func parseConfig() (*Config, error) {
	config := &Config{}

	err := env.Parse(config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
