package product

import (
	"go-fiber-template/internal/domain/dto"
	"go-fiber-template/internal/domain/interfaces"
	"go-fiber-template/lib/middleware"
	"go-fiber-template/lib/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type httpHandler struct {
	productService interfaces.ProductService
}

func NewHttpHandler(r fiber.Router, productService interfaces.ProductService) {
	handler := &httpHandler{
		productService: productService,
	}

	r.Post("/", middleware.Validate[dto.CreateProductRequest](), handler.Create)
	r.Get("/", handler.FindAll)
	r.Get("/:id", handler.FindByID)
	r.Put("/:id", middleware.Validate[dto.UpdateProductRequest](), handler.Update)
	r.Delete("/:id", handler.Delete)
}

func (h *httpHandler) Create(c *fiber.Ctx) error {
	req := utils.ExtractStructFromValidator[dto.CreateProductRequest](c)
	data, err := h.productService.Create(c, req)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(dto.ResponseDto{
		Message: "Product created successfully",
		Data:    data,
	})
}

func (h *httpHandler) FindAll(c *fiber.Ctx) error {
	products, err := h.productService.FindAll(c)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(dto.ResponseDto{
		Message: "Products fetched successfully",
		Data:    products,
	})
}

func (h *httpHandler) FindByID(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return err
	}

	data, err := h.productService.FindByID(c, uint(id))
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(dto.ResponseDto{
		Message: "Product fetched successfully",
		Data:    data,
	})
}

func (h *httpHandler) Update(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return err
	}

	req := utils.ExtractStructFromValidator[dto.UpdateProductRequest](c)
	data, err := h.productService.Update(c, uint(id), req)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(dto.ResponseDto{
		Message: "Product updated successfully",
		Data:    data,
	})
}

func (h *httpHandler) Delete(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return err
	}

	if err := h.productService.Delete(c, uint(id)); err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(dto.ResponseDto{
		Message: "Product deleted successfully",
	})
}
