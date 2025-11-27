package xjwt

import "github.com/google/uuid"

type Config struct {
	SecretKey             string
	AccessTokenExpiresIn  int64
	RefreshTokenExpiresIn int64
	Issuer                string
}

var DefaultConfig = Config{
	SecretKey:             uuid.NewString(),
	AccessTokenExpiresIn:  3600,   // 1 hour
	RefreshTokenExpiresIn: 604800, // 7 days
}

func setConfig(config ...Config) Config {
	if len(config) == 0 {
		return DefaultConfig
	}

	cfg := config[0]
	if cfg.SecretKey == "" {
		cfg.SecretKey = DefaultConfig.SecretKey
	}
	if cfg.AccessTokenExpiresIn == 0 {
		cfg.AccessTokenExpiresIn = DefaultConfig.AccessTokenExpiresIn
	}
	if cfg.RefreshTokenExpiresIn == 0 {
		cfg.RefreshTokenExpiresIn = DefaultConfig.RefreshTokenExpiresIn
	}

	return cfg
}
