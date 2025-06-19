package user

import (
	"go-fiber-template/internal/domain/dto"
	"go-fiber-template/internal/domain/entity"
	"go-fiber-template/internal/domain/interfaces"

	"github.com/gofiber/fiber/v2"
)

type service struct {
	userRepo interfaces.UserRepository
}

func (s *service) FindByID(c *fiber.Ctx, id uint) (*dto.UserDto, error) {
	user, _ := s.userRepo.FindByID(id)
	if user == nil {
		return nil, fiber.NewError(fiber.StatusNotFound, "user not found")
	}

	return constructUserDto(user), nil
}

func constructUserDto(user *entity.User) *dto.UserDto {
	return &dto.UserDto{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: user.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}

func NewService(userRepo interfaces.UserRepository) interfaces.UserService {
	return &service{
		userRepo: userRepo,
	}
}
