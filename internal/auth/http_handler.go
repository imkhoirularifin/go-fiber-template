package auth

import (
	"go-fiber-template/internal/domain/dto"
	"go-fiber-template/internal/domain/interfaces"
	"go-fiber-template/lib/middleware"
	"go-fiber-template/lib/utils"

	"github.com/gofiber/fiber/v2"
)

type httpHandler struct {
	authService interfaces.AuthService
}

func NewHttpHandler(r fiber.Router, authService interfaces.AuthService) {
	handler := &httpHandler{
		authService: authService,
	}

	r.Post("/register", middleware.Validate[dto.RegisterRequest](), handler.Register)
	r.Post("/login", middleware.Validate[dto.LoginRequest](), handler.Login)
}

// @Summary		Register a new user
// @Description	Register a new user with email and password
// @Tags			Auth
// @Accept			application/json
// @Produce		application/json
// @Param			request	body		dto.RegisterRequest	true	"User registration request"
// @Success		201		{object}	dto.ResponseDto{data=dto.RegisterResponse}
// @Failure		400		{object}	dto.ResponseDto
// @Failure		409		{object}	dto.ResponseDto
// @Failure		500		{object}	dto.ResponseDto
// @Router			/auth/register [post]
func (h *httpHandler) Register(c *fiber.Ctx) error {
	req := utils.ExtractStructFromValidator[dto.RegisterRequest](c)
	data, err := h.authService.Register(c, req)
	if err != nil {
		return err
	}
	return c.Status(fiber.StatusCreated).JSON(dto.ResponseDto{
		Message: "User registered successfully",
		Data:    data,
	})
}

// @Summary		User login
// @Description	User login with email and password
// @Tags			Auth
// @Accept			application/json
// @Produce		application/json
// @Param			request	body		dto.LoginRequest	true	"User login request"
// @Success		200		{object}	dto.ResponseDto{data=dto.LoginResponse}
// @Failure		400		{object}	dto.ResponseDto
// @Failure		401		{object}	dto.ResponseDto
// @Failure		500		{object}	dto.ResponseDto
// @Router			/auth/login [post]
func (h *httpHandler) Login(c *fiber.Ctx) error {
	req := utils.ExtractStructFromValidator[dto.LoginRequest](c)
	data, err := h.authService.Login(c, req)
	if err != nil {
		return err
	}
	return c.Status(fiber.StatusOK).JSON(dto.ResponseDto{
		Message: "Login successful",
		Data:    data,
	})
}
