package config

import (
	"github.com/caarlos0/env/v11"
	_ "github.com/joho/godotenv/autoload"
)

type AppConfig struct {
	Port string `env:"PORT" envDefault:"3000"`
}

func Setup() (AppConfig, error) {
	var cfg AppConfig
	if err := env.Parse(&cfg); err != nil {
		return AppConfig{}, err
	}

	return cfg, nil
}
