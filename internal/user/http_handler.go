package user

import (
	"go-fiber-template/internal/domain/dto"
	"go-fiber-template/internal/domain/interfaces"
	"go-fiber-template/lib/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/ryanbekhen/di"
)

type httpHandler struct {
	userService interfaces.UserService
}

func NewHttpHandler(r fiber.Router) {
	userService := di.MustResolve[interfaces.UserService]()
	authMiddleware := di.MustResolve[middleware.JWTAuthMiddleware]()

	handler := &httpHandler{
		userService: userService,
	}

	r.Get("/self", authMiddleware.Validate(), handler.FindSelf)
}

// @Summary		Find self user
// @Description	Find self user
// @Tags			User
// @Accept			application/json
// @Produce		application/json
// @Success		200	{object}	dto.ResponseDto{data=dto.UserDto}
// @Failure		400		{object}	dto.ResponseDto
// @Failure		401		{object}	dto.ResponseDto
// @Failure		404		{object}	dto.ResponseDto
// @Failure		500		{object}	dto.ResponseDto
// @Router			/users/self [get]
// @Security		Bearer
func (h *httpHandler) FindSelf(c *fiber.Ctx) error {
	data, err := h.userService.FindSelf(c)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(dto.ResponseDto{
		Message: "User found",
		Data:    data,
	})
}
