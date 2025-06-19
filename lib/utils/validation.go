package utils

import "github.com/gofiber/fiber/v2"

func ExtractStructFromValidator[V any](c *fiber.Ctx) *V {
	v, ok := c.Locals("parser").(*V)
	if !ok {
		return v
	}
	return v
}
