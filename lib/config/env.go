package config

import (
	"github.com/caarlos0/env/v11"
	"github.com/go-playground/validator/v10"
	_ "github.com/joho/godotenv/autoload"
)

var (
	Config AppConfig
)

type AppConfig struct {
	AppName   string         `env:"APP_NAME" envDefault:"go-fiber-template"`
	Port      string         `env:"PORT" envDefault:"3000"`
	GoEnv     string         `env:"GO_ENV" envDefault:"development" validate:"oneof=development production"`
	LogFields []string       `env:"LOG_FIELDS" envSeparator:"," envDefault:"latency,status,method,url,error"`
	Jwt       JwtConfig      `envPrefix:"JWT_"`
	Database  DatabaseConfig `envPrefix:"DB_"`
	Apitally  ApitallyConfig `envPrefix:"APITALLY_"`
	Kafka     KafkaConfig    `envPrefix:"KAFKA_"`
}

type JwtConfig struct {
	SecretKey string `env:"SECRET_KEY" envDefault:"secret"`
	ExpiredAt int64  `env:"EXPIRED_AT" envDefault:"3600"`
}

type DatabaseConfig struct {
	Driver string `env:"DRIVER" envDefault:"postgres"`
	Dsn    string `env:"DSN" envDefault:"host=localhost user=postgres password=secret dbname=mydb port=5432 sslmode=disable"`
}

type ApitallyConfig struct {
	ClientId string `env:"CLIENT_ID,notEmpty"`
}

type KafkaConfig struct {
	Brokers []string `env:"BROKERS" envSeparator:"," envDefault:"localhost:9092"`
	GroupId string   `env:"GROUP_ID" envDefault:"go-fiber-template"`
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

	Config = cfg
	return cfg
}
