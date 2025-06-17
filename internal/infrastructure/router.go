package infrastructure

import (
	x_app "go-fiber-template/internal/app"

	"github.com/gofiber/fiber/v2"
)

func registerRoutes(r fiber.Router) {
	x_app.NewHttpHandler(r)
}
