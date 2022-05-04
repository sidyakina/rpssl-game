package main

import (
	"github.com/caarlos0/env/v6"
)

type Config struct {
	GRPCPort                 int64  `env:"GRPC_PORT"`
	RandomNumberGeneratorURI string `env:"RANDOM_NUMBER_GENERATOR_URI"`
}

func parseConfig() (*Config, error) {
	config := &Config{}

	err := env.Parse(config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
