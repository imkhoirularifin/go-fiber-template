package infrastructure

import (
	"go-fiber-template/lib/config"
	"go-fiber-template/pkg/database"
	"go-fiber-template/pkg/xjwt"
	"go-fiber-template/pkg/xlogger"
	"go-fiber-template/pkg/xvalidator"
)

func setupConfig() config.AppConfig {
	cfg := config.NewConfig()
	return *cfg
}

func setupXlogger() {
	xlogger.Setup(cfg)
}

func setupXValidator() *xvalidator.Client {
	xvalidatorClient := xvalidator.NewClient(
		xvalidator.WithCustomValidator(&xvalidator.DateValidator{}),
		xvalidator.WithCustomValidator(&xvalidator.PasswordValidator{}),
	)
	return xvalidatorClient
}

func setupDatabase() *database.Database {
	// var logLevel string
	// if cfg.GoEnv == "production" {
	// 	logLevel = "error"
	// } else {
	// 	logLevel = "info"
	// }

	dbInstance := database.New(database.Config{
		Driver: cfg.Database.Driver,
		Dsn:    cfg.Database.DbString,
		// LogLevel: logLevel,
	})

	return dbInstance
}

func setupJwtClient() xjwt.Client {
	return xjwt.NewClient(xjwt.Config{
		SecretKey:             cfg.Jwt.SecretKey,
		AccessTokenExpiresIn:  cfg.Jwt.AccessTokenExpiresIn,
		RefreshTokenExpiresIn: cfg.Jwt.RefreshTokenExpiresIn,
		Issuer:                cfg.Jwt.Issuer,
	})
}
