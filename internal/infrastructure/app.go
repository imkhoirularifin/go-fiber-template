package infrastructure

import (
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

// setupApp creates and configures the Fiber application
func setupApp() *fiber.App {
	app := fiber.New(config.FiberCfg(cfg))

	// Setup middleware
	setupMiddleware(app)

	// Setup routes
	setupRoutes(app)

	return app
}

// setupMiddleware configures all middleware for the application
func setupMiddleware(app *fiber.App) {
	app.Use(fiberi18n.New(config.I18nConfig))
	app.Use(apitally.Middleware(app, config.ApitallyCfg(cfg)))
	app.Use(fiberzerolog.New(config.FiberZerologCfg(cfg)))
	app.Use(recover.New())
	app.Use(cors.New(config.CorsCfg))
	app.Use(cache.New(config.CacheCfg))
}

// setupRoutes configures all routes for the application
func setupRoutes(app *fiber.App) {
	api := app.Group("/api/v1")
	registerRoutes(api)
	app.Use(common.NotFoundHandler)
}
