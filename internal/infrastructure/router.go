package infrastructure

import (
	x_app "nothing-go/internal/app"

	"github.com/gofiber/fiber/v2"
)

func registerRoutes(r fiber.Router) {
	x_app.NewHttpHandler(r)
}
