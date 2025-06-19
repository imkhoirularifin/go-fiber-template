package config

import (
	"github.com/caarlos0/env/v11"
	"github.com/go-playground/validator/v10"
	_ "github.com/joho/godotenv/autoload"
)

type AppConfig struct {
	Port  string `env:"PORT" envDefault:"3000"`
	GoEnv string `env:"GO_ENV" envDefault:"development" validate:"oneof=development production"`
	// Log field docs: https://github.com/gofiber/contrib/blob/main/fiberzerolog/config.go
	LogFields []string `env:"LOG_FIELDS" envSeparator:","`
}

func Setup() AppConfig {
	var cfg AppConfig
	validate := validator.New(validator.WithRequiredStructEnabled())

	if err := env.Parse(&cfg); err != nil {
		panic(err)
	}

	err := validate.Struct(cfg)
	if err != nil {
		panic(err)
	}

	return cfg
}
