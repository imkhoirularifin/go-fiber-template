package app

import (
	"github.com/gofiber/fiber/v2"
)

type httpHandler struct{}

func NewHttpHandler(r fiber.Router) {
	handler := &httpHandler{}

	r.Get("/ping", handler.ping)
}

// @Summary		Ping
// @Description	Ping the server
// @Tags			Health Check
// @Accept			application/json
// @Produce		application/json
// @Success		200	{string}	string	"pong"
// @Router			/ping [get]
func (h *httpHandler) ping(c *fiber.Ctx) error {
	return c.SendString("pong")
}
