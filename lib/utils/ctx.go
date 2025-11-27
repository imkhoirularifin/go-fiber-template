package utils

import (
	"go-fiber-template/pkg/xjwt"

	"github.com/gofiber/fiber/v2"
)

func ExtractStructFromValidator[V any](c *fiber.Ctx) *V {
	v, ok := c.Locals("parser").(*V)
	if !ok {
		return v
	}
	return v
}

func ExtractTokenClaimFromCtx(c *fiber.Ctx) *xjwt.TokenClaims {
	v, ok := c.Locals("tokenClaims").(*xjwt.TokenClaims)
	if !ok {
		return nil
	}

	return v
}
