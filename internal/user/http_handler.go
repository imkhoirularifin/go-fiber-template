package user

import (
	"go-fiber-template/internal/domain/dto"
	"go-fiber-template/internal/domain/interfaces"
	"go-fiber-template/lib/middleware"

	"github.com/gofiber/fiber/v2"
)

type httpHandler struct {
	userService interfaces.UserService
}

func NewHttpHandler(r fiber.Router, userService interfaces.UserService) {
	handler := &httpHandler{
		userService: userService,
	}

	r.Get("/:id", middleware.Protected(), handler.FindByID)
}

// @Summary		Find user by ID
// @Description	Find user by ID
// @Tags			User
// @Accept			application/json
// @Produce		application/json
// @Security		Bearer
// @Param			id	path		int	true	"User ID"
// @Success		200	{object}	dto.ResponseDto{data=dto.UserDto}
// @Failure		400	{object}	dto.ResponseDto
// @Failure		404	{object}	dto.ResponseDto
// @Failure		500	{object}	dto.ResponseDto
// @Router			/users/{id} [get]
func (h *httpHandler) FindByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid user ID")
	}

	data, err := h.userService.FindByID(c, uint(id))
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(dto.ResponseDto{
		Message: "User found",
		Data:    data,
	})
}
