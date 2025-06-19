package infrastructure

import (
	"go-fiber-template/internal/auth"
	"go-fiber-template/internal/domain/interfaces"
	"go-fiber-template/internal/user"
	"go-fiber-template/lib/config"
	"go-fiber-template/lib/database"
	"go-fiber-template/lib/xlogger"
	"go-fiber-template/lib/xvalidator"

	"gorm.io/gorm"
)

var (
	cfg        config.AppConfig
	dbInstance *database.Database
	db         *gorm.DB

	userRepository interfaces.UserRepository

	authService interfaces.AuthService
	userService interfaces.UserService
)

func init() {
	cfg = config.Setup()
	xlogger.Setup(cfg)
	xvalidator.Setup()
	setupDB()

	userRepository = user.NewRepository(db)

	authService = auth.NewService(userRepository)
	userService = user.NewService(userRepository)
}
