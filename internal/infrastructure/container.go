package infrastructure

import (
	"go-fiber-template/lib/config"
	"go-fiber-template/pkg/database"
	"go-fiber-template/pkg/xvalidator"

	"gorm.io/gorm"
)

var (
	cfg              config.AppConfig
	dbInstance       *database.Database
	db               *gorm.DB
	xvalidatorClient *xvalidator.Client
)

func init() {
	// setup dependencies
	cfg = setupConfig()
	setupXlogger()
	xvalidatorClient = setupXValidator()
	dbInstance = setupDatabase()
	db = dbInstance.GetDB()

	// register dependencies
	registerDependencies()

	// register repositories
	registerRepositories()

	// register middlewares
	registerMiddlewares()

	// register services
	registerServices()
}
