package interfaces

import (
	"go-fiber-template/internal/domain/dto"
	"go-fiber-template/internal/domain/entity"

	"github.com/gofiber/fiber/v2"
)

type UserRepository interface {
	Create(data *entity.User) error
	FindByID(id uint) (*entity.User, error)
	FindByEmail(email string) (*entity.User, error)
	Update(user *entity.User) error
	Delete(id uint) error
}

type UserService interface {
	FindByID(c *fiber.Ctx, id uint) (*dto.UserDto, error)
}
