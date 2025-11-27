package infrastructure

import (
	"go-fiber-template/internal/product"
	"go-fiber-template/internal/user"
	"go-fiber-template/lib/config"
	"go-fiber-template/lib/middleware"
	"go-fiber-template/pkg/xjwt"
	"go-fiber-template/pkg/xvalidator"

	"github.com/ryanbekhen/di"
	"gorm.io/gorm"
)

func registerDependencies() {
	// register database
	di.RegisterFactory(func() *gorm.DB {
		return db
	})
	// register config
	di.RegisterFactory(func() config.AppConfig {
		return cfg
	})
	// register xvalidator
	di.RegisterFactory(func() *xvalidator.Client {
		return xvalidatorClient
	})
	// register jwt client
	di.RegisterFactory(func() xjwt.Client {
		return jwtClient
	})
}

func registerRepositories() {
	product.RegisterRepository()
	user.RegisterRepository()
}

func registerMiddlewares() {
	middleware.RegisterJWTAuth()
}

func registerServices() {
	product.RegisterService()
	user.RegisterService()
}
