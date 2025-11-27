package middleware

import (
	"go-fiber-template/pkg/xjwt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/ryanbekhen/di"
)

type JWTAuthMiddleware interface {
	Validate() fiber.Handler
}

type jwtAuthMiddleware struct {
	jwtClient xjwt.Client
}

func (j *jwtAuthMiddleware) Validate() fiber.Handler {
	return func(c *fiber.Ctx) error {
		token := c.Get("Authorization")
		if token == "" {
			return fiber.ErrUnauthorized
		}
		token = strings.TrimPrefix(token, "Bearer ")

		tokenClaims, err := j.jwtClient.ValidateToken(xjwt.ValidateTokenRequest{
			Token: token,
		})
		if err != nil {
			return fiber.ErrUnauthorized
		}

		// set to local ctx
		c.Locals("tokenClaims", tokenClaims)

		return c.Next()
	}
}

func RegisterJWTAuth() {
	jwtClient := di.MustResolve[xjwt.Client]()

	di.RegisterFactory(func() JWTAuthMiddleware {
		return &jwtAuthMiddleware{
			jwtClient: jwtClient,
		}
	})
}
