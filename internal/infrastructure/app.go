package infrastructure

import (
	x_app "go-fiber-template/internal/app"
	"go-fiber-template/internal/auth"
	"go-fiber-template/internal/docs"
	"go-fiber-template/internal/product"
	"go-fiber-template/internal/user"
	"go-fiber-template/lib/common"
	"go-fiber-template/lib/config"

	apitally "github.com/apitally/apitally-go/fiber"
	"github.com/gofiber/contrib/fiberi18n/v2"
	"github.com/gofiber/contrib/fiberzerolog"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

type App struct {
	container *Container
	server    *fiber.App
}

func NewApp(container *Container) *App {
	app := fiber.New(config.FiberCfg(container.Config))
	setupMiddleware(app, container.Config)
	setupRoutes(app, container)
	return &App{
		container: container,
		server:    app,
	}
}

func setupMiddleware(app *fiber.App, cfg config.AppConfig) {
	app.Use(fiberi18n.New(config.I18nConfig))
	app.Use(apitally.Middleware(app, config.ApitallyCfg(cfg)))
	app.Use(fiberzerolog.New(config.FiberZerologCfg(cfg)))
	app.Use(recover.New())
	app.Use(cors.New(config.CorsCfg))
	app.Use(cache.New(config.CacheCfg))
}

func setupRoutes(app *fiber.App, container *Container) {
	api := app.Group("/api/v1")
	x_app.NewHttpHandler(api)
	docs.NewHttpHandler(api.Group("/docs"))
	auth.NewHttpHandler(api.Group("/auth"), container.AuthService)
	user.NewHttpHandler(api.Group("/users"), container.UserService)
	product.NewHttpHandler(api.Group("/products"), container.ProductService)
	app.Use(common.NotFoundHandler)
}

func (a *App) GetServer() *fiber.App {
	return a.server
}

func (a *App) GetContainer() *Container {
	return a.container
}

func Run() {
	container := NewContainer()
	app := NewApp(container)
	server := NewServer(app)
	server.Start()
}
