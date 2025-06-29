package infrastructure

import (
	"go-fiber-template/internal/auth"
	"go-fiber-template/internal/domain/interfaces"
	"go-fiber-template/internal/email"
	"go-fiber-template/internal/product"
	"go-fiber-template/internal/user"
	"go-fiber-template/lib/config"
	"go-fiber-template/lib/database"
	"go-fiber-template/lib/xkafka"
	"go-fiber-template/lib/xlogger"
	"go-fiber-template/lib/xvalidator"

	"gorm.io/gorm"
)

var (
	cfg         config.AppConfig
	dbInstance  *database.Database
	db          *gorm.DB
	kafkaClient *xkafka.Client

	authService    interfaces.AuthService
	userService    interfaces.UserService
	emailService   interfaces.EmailService
	productService interfaces.ProductService
)

func init() {
	cfg = config.Setup()
	xlogger.Setup(cfg)
	xvalidator.Setup()

	dbInstance = database.New(database.Config{
		Driver: cfg.Database.Driver,
		Dsn:    cfg.Database.Dsn,
	})
	db = dbInstance.GetDB()

	kafkaClient = xkafka.Setup(cfg.Kafka)

	userRepository := user.NewRepository(db)
	productRepository := product.NewRepository(db)

	authService = auth.NewService(userRepository, kafkaClient)
	userService = user.NewService(userRepository)
	emailService = email.NewService(kafkaClient)
	productService = product.NewService(productRepository)
}
