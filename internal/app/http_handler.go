package app

import (
	"go-fiber-template/internal/domain/dto"

	"github.com/gofiber/contrib/fiberi18n/v2"
	"github.com/gofiber/fiber/v2"
)

type httpHandler struct{}

func NewHttpHandler(r fiber.Router) {
	handler := &httpHandler{}

	r.Get("/ping", handler.ping)
	r.Get("/welcome", handler.welcome)
}

// @Summary		Ping
// @Description	Ping the server
// @Tags			App
// @Accept			application/json
// @Produce		application/json
// @Success		200	{string}	string	"pong"
// @Router			/ping [get]
func (h *httpHandler) ping(c *fiber.Ctx) error {
	return c.SendString("pong")
}

// @Summary		Welcome
// @Description	Get welcome message
// @Tags			App
// @Accept			application/json
// @Produce		application/json
// @Param			Accept-Language	header	string	true	"Language code for localization"
// @Success		200	{object}	dto.ResponseDto
// @Router			/welcome [get]
func (h *httpHandler) welcome(c *fiber.Ctx) error {
	msg, err := fiberi18n.Localize(c, "welcome_message")
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(dto.ResponseDto{
		Message: msg,
	})
}
