package docs

import (
	"go-fiber-template/docs"

	"github.com/gofiber/contrib/swagger"
	"github.com/gofiber/fiber/v2"
	"github.com/swaggo/swag"
)

type HttpDocsHandler struct{}

func NewHttpHandler(r fiber.Router) {
	handler := &HttpDocsHandler{}
	r.Get("/swagger.json", handler.ServeSwaggerJSON)
	r.Use(swagger.New(swagger.Config{
		BasePath: "/api/v1",
		Path:     "/docs",
		FilePath: "./docs/swagger.json",
	}))
}

func (h *HttpDocsHandler) ServeSwaggerJSON(c *fiber.Ctx) error {
	docs.SwaggerInfo.Host = c.Hostname()
	doc, err := swag.ReadDoc()
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to read swagger doc")
	}
	return c.SendString(doc)
}
