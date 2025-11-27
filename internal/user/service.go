package user

import (
	"go-fiber-template/internal/domain/dto"
	"go-fiber-template/internal/domain/entity"
	"go-fiber-template/internal/domain/interfaces"
	"go-fiber-template/lib/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/ryanbekhen/di"
)

type service struct {
	userRepo interfaces.UserRepository
}

func (s *service) FindSelf(c *fiber.Ctx) (*dto.UserDto, error) {
	claims := utils.ExtractTokenClaimFromCtx(c)

	id, err := utils.ConvertStringToUint(claims.Subject)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusUnauthorized, "invalid token")
	}

	user, _ := s.userRepo.FindByID(id)
	if user == nil {
		return nil, fiber.NewError(fiber.StatusNotFound, "user not found")
	}

	return userToDto(user), nil
}

func userToDto(user *entity.User) *dto.UserDto {
	return &dto.UserDto{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func RegisterService() {
	userRepo := di.MustResolve[interfaces.UserRepository]()

	di.RegisterFactory(func() interfaces.UserService {
		return &service{
			userRepo: userRepo,
		}
	})
}
