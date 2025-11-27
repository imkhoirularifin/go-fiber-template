package config

import (
	"github.com/caarlos0/env/v11"
	"github.com/go-playground/validator/v10"
	_ "github.com/joho/godotenv/autoload"
)

// NewConfig creates a new AppConfig.
func NewConfig() *AppConfig {
	cfg := AppConfig{}
	validate := validator.New(validator.WithRequiredStructEnabled())

	if err := env.Parse(&cfg); err != nil {
		panic(err)
	}

	err := validate.Struct(cfg)
	if err != nil {
		panic(err)
	}

	return &cfg
}

type AppConfig struct {
	AppName   string         `env:"APP_NAME" envDefault:"go-fiber-template"`
	Port      string         `env:"PORT" envDefault:"3000"`
	GoEnv     string         `env:"GO_ENV" envDefault:"development" validate:"oneof=development production"`
	LogFields []string       `env:"LOG_FIELDS" envSeparator:"," envDefault:"latency,status,method,url,error"`
	Jwt       JwtConfig      `envPrefix:"JWT_"`
	Database  DatabaseConfig `envPrefix:"GOOSE_"`
	Cors      CorsConfig     `envPrefix:"CORS_"`
}

type JwtConfig struct {
	SecretKey             string `env:"SECRET_KEY,notEmpty"`
	AccessTokenExpiresIn  int64  `env:"ACCESS_TOKEN_EXPIRES_IN" envDefault:"3600"`    // 1 Hour
	RefreshTokenExpiresIn int64  `env:"REFRESH_TOKEN_EXPIRES_IN" envDefault:"604800"` // 7 Days
	Issuer                string `env:"ISSUER" envDefault:"go-fiber-template"`
}

type DatabaseConfig struct {
	Driver   string `env:"DRIVER" envDefault:"postgres"`
	DbString string `env:"DBSTRING,notEmpty"`
}

// CorsConfig is the configuration for the CORS.
type CorsConfig struct {
	AllowOrigins     string `env:"ALLOW_ORIGINS" envDefault:"*"`
	AllowCredentials bool   `env:"ALLOW_CREDENTIALS" envDefault:"false"`
}
