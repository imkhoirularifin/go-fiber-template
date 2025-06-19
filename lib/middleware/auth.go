package middleware

import (
	"go-fiber-template/internal/domain/dto"
	"go-fiber-template/lib/config"
	"go-fiber-template/lib/xjwt"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

// Protected protect routes
func Protected() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{
			JWTAlg: jwt.SigningMethodHS256.Name,
			Key:    []byte(config.Config.Jwt.SecretKey),
		},
		ContextKey:     "user",
		ErrorHandler:   jwtError,
		SuccessHandler: jwtSuccess,
	})
}

func jwtError(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ResponseDto{
			Message: "Missing or malformed JWT",
		})
	}
	return c.Status(fiber.StatusUnauthorized).JSON(dto.ResponseDto{
		Message: "Invalid or expired JWT",
	})
}

func jwtSuccess(c *fiber.Ctx) error {
	// Get the user from the context
	jwtToken, ok := c.Locals("user").(*jwt.Token)
	if !ok {
		return fiber.NewError(fiber.StatusUnauthorized, "Invalid JWT token")
	}
	customClaims, err := xjwt.MapClaimsToTokenClaims(jwtToken)
	if err != nil {
		return err
	}
	// Set the user in the context
	c.Locals("claims", customClaims)
	return c.Next()
}
