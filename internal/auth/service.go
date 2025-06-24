package auth

import (
	"encoding/json"
	"fmt"
	"go-fiber-template/internal/domain/dto"
	"go-fiber-template/internal/domain/entity"
	"go-fiber-template/internal/domain/interfaces"
	"go-fiber-template/lib/utils"
	"go-fiber-template/lib/xjwt"
	"go-fiber-template/lib/xkafka"

	"github.com/gofiber/fiber/v2"
)

type service struct {
	userRepo    interfaces.UserRepository
	kafkaClient *xkafka.Client
}

func (s *service) Login(c *fiber.Ctx, req *dto.LoginRequest) (*dto.LoginResponse, error) {
	byEmail, err := s.userRepo.FindByEmail(req.Email)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusUnauthorized, "Invalid email or password")
	}

	if !utils.CheckPasswordHash(req.Password, byEmail.Password) {
		return nil, fiber.NewError(fiber.StatusUnauthorized, "Invalid email or password")
	}

	accessToken, err := xjwt.GenerateToken(byEmail, xjwt.TokenTypeAccess)
	if err != nil {
		return nil, err
	}

	if err := s.sendLoginNotification(c, byEmail); err != nil {
		return nil, err
	}

	return &dto.LoginResponse{
		AccessToken: accessToken,
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

	accessToken, err := xjwt.GenerateToken(user, xjwt.TokenTypeAccess)
	if err != nil {
		return nil, err
	}

	return &dto.RegisterResponse{
		UserID:      user.ID,
		AccessToken: accessToken,
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

func (s *service) sendLoginNotification(c *fiber.Ctx, user *entity.User) error {
	// Create email configuration for login notification
	emailConfig := &interfaces.EmailConfig{
		To:      user.Email,
		Subject: "Login Notification",
		Body:    fmt.Sprintf("Hello %s, you have successfully logged in to your account.", user.Name),
	}

	emailConfigByte, err := json.Marshal(emailConfig)
	if err != nil {
		return err
	}

	if err := s.kafkaClient.Produce(c.Context(), "auth.login", emailConfigByte); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to send login notification")
	}

	return nil
}

func NewService(
	userRepo interfaces.UserRepository,
	kafkaClient *xkafka.Client,
) interfaces.AuthService {
	return &service{
		userRepo:    userRepo,
		kafkaClient: kafkaClient,
	}
}
