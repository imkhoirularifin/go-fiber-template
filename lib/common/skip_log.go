package common

import "github.com/gofiber/fiber/v2"

func SkipLog(c *fiber.Ctx) bool {
	return c.Method() == fiber.MethodOptions
}
