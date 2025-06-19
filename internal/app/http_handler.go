package app

import (
	"github.com/gofiber/fiber/v2"
)

type httpHandler struct{}

func NewHttpHandler(r fiber.Router) {
	handler := &httpHandler{}

	r.Get("/ping", handler.ping)
}

func (h *httpHandler) ping(c *fiber.Ctx) error {
	return c.SendString("pong")
}
