package infrastructure

import (
	x_app "go-fiber-template/internal/app"
	"go-fiber-template/internal/auth"
	"go-fiber-template/internal/docs"
	"go-fiber-template/internal/product"
	"go-fiber-template/internal/user"
	"go-fiber-template/lib/common"

	"github.com/gofiber/fiber/v2"
)

func registerRoutes(app *fiber.App) {
	api := app.Group("/api/v1")
	x_app.NewHttpHandler(api)
	docs.NewHttpHandler(api.Group("/docs"))
	auth.NewHttpHandler(api.Group("/auth"), authService)
	user.NewHttpHandler(api.Group("/users"), userService)
	product.NewHttpHandler(api.Group("/products"), productService)
	app.Use(common.NotFoundHandler)
}
