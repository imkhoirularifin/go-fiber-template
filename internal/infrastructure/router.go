package infrastructure

import (
	x_app "go-fiber-template/internal/app"
	"go-fiber-template/internal/auth"
	"go-fiber-template/internal/docs"
	"go-fiber-template/internal/product"
	"go-fiber-template/internal/user"

	"github.com/gofiber/fiber/v2"
)

func registerRoutes(r fiber.Router) {
	x_app.NewHttpHandler(r)
	docs.NewHttpHandler(r.Group("/docs"))
	auth.NewHttpHandler(r.Group("/auth"), authService)
	user.NewHttpHandler(r.Group("/users"), userService)
	product.NewHttpHandler(r.Group("/products"), productService)
}
