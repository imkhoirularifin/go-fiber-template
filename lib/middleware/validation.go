package middleware

import (
	"go-fiber-template/internal/domain/dto"
	"go-fiber-template/lib/xvalidator"

	"github.com/gofiber/fiber/v2"
)

func Validate[V any]() fiber.Handler {
	validate := xvalidator.XValidator
	return func(c *fiber.Ctx) error {
		var v V
		if err := c.BodyParser(&v); err != nil {
			return err
		}

		if validationErrors := validate.ValidateStruct(v); validationErrors != nil {
			return c.Status(fiber.StatusUnprocessableEntity).JSON(dto.ResponseDto{
				Message: "Validation Error",
				Errors:  validationErrors,
			})
		}

		c.Locals("parser", &v)
		return c.Next()
	}
}
