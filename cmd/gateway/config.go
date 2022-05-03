package main

import "github.com/caarlos0/env/v6"

type Config struct {
	Port int64 `env:"HTTP_PORT"`
}

func parseConfig() (*Config, error) {
	config := &Config{}

	err := env.Parse(config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
