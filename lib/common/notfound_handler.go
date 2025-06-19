package common

import (
	"go-fiber-template/internal/domain/dto"

	"github.com/gofiber/fiber/v2"
)

func NotFoundHandler(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotFound).JSON(dto.ResponseDto{
		Message: "Look like you're lost",
	})
}
