package interfaces

import (
	"go-fiber-template/internal/domain/dto"

	"github.com/gofiber/fiber/v2"
)

type AuthService interface {
	Register(c *fiber.Ctx, req *dto.RegisterRequest) (*dto.RegisterResponse, error)
	Login(c *fiber.Ctx, req *dto.LoginRequest) (*dto.LoginResponse, error)
}
