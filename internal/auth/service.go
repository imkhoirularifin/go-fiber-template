package auth

import (
	"go-fiber-template/internal/domain/dto"
	"go-fiber-template/internal/domain/entity"
	"go-fiber-template/internal/domain/interfaces"
	"go-fiber-template/lib/utils"
	"go-fiber-template/pkg/xjwt"

	"github.com/gofiber/fiber/v2"
	"github.com/ryanbekhen/di"
)

type service struct {
	userRepo  interfaces.UserRepository
	jwtClient xjwt.Client
}

func (s *service) Login(c *fiber.Ctx, req *dto.LoginRequest) (*dto.LoginResponse, error) {
	byEmail, err := s.userRepo.FindByEmail(req.Email)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusUnauthorized, "Invalid email or password")
	}

	if !utils.CheckPasswordHash(req.Password, byEmail.Password) {
		return nil, fiber.NewError(fiber.StatusUnauthorized, "Invalid email or password")
	}

	tokenResponse, err := s.jwtClient.GenerateToken(xjwt.GenerateTokenRequest{
		User: xjwt.UserInfo{
			ID:    string(byEmail.ID),
			Name:  byEmail.Name,
			Email: byEmail.Email,
		},
	})
	if err != nil {
		return nil, err
	}

	return &dto.LoginResponse{
		AccessToken:  tokenResponse.AccessToken,
		RefreshToken: tokenResponse.RefreshToken,
		ExpiresAt:    tokenResponse.ExpiresAt,
	}, nil
}

func (s *service) Register(c *fiber.Ctx, req *dto.RegisterRequest) (*dto.RegisterResponse, error) {
	user := &entity.User{
		Name:  req.Name,
		Email: req.Email,
	}

	if err := s.validateUnique(user); err != nil {
		return nil, err
	}

	hash, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}
	user.Password = hash

	if err := s.userRepo.Create(user); err != nil {
		return nil, err
	}

	tokenResponse, err := s.jwtClient.GenerateToken(xjwt.GenerateTokenRequest{
		User: xjwt.UserInfo{
			ID:    string(user.ID),
			Name:  user.Name,
			Email: user.Email,
		},
	})
	if err != nil {
		return nil, err
	}

	return &dto.RegisterResponse{
		UserID:       user.ID,
		AccessToken:  tokenResponse.AccessToken,
		RefreshToken: tokenResponse.RefreshToken,
		ExpiresAt:    tokenResponse.ExpiresAt,
	}, nil
}

func (s *service) RefreshToken(c *fiber.Ctx, req *dto.RefreshTokenRequest) (*dto.RefreshTokenResponse, error) {
	tokenClaims, err := s.jwtClient.ValidateToken(xjwt.ValidateTokenRequest{
		Token: req.RefreshToken,
	})
	if err != nil {
		return nil, err
	}

	tokenResponse, err := s.jwtClient.GenerateToken(xjwt.GenerateTokenRequest{
		User: xjwt.UserInfo{
			ID:    tokenClaims.TokenClaims.Subject,
			Name:  tokenClaims.TokenClaims.UserName,
			Email: tokenClaims.TokenClaims.UserEmail,
		},
	})
	if err != nil {
		return nil, err
	}

	return &dto.RefreshTokenResponse{
		AccessToken:  tokenResponse.AccessToken,
		RefreshToken: tokenResponse.RefreshToken,
		ExpiresAt:    tokenResponse.ExpiresAt,
	}, nil
}

func (s *service) validateUnique(user *entity.User) error {
	if user.Email != "" {
		byEmail, _ := s.userRepo.FindByEmail(user.Email)
		if byEmail != nil && byEmail.ID != user.ID {
			return fiber.NewError(fiber.StatusConflict, "Email already exists")
		}
	}

	return nil
}

func RegisterService() {
	userRepo := di.MustResolve[interfaces.UserRepository]()
	jwtClient := di.MustResolve[xjwt.Client]()

	di.RegisterFactory(func() interfaces.AuthService {
		return &service{
			userRepo:  userRepo,
			jwtClient: jwtClient,
		}
	})
}
