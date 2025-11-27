package product

import (
	"go-fiber-template/internal/domain/dto"
	"go-fiber-template/internal/domain/interfaces"
	"go-fiber-template/lib/middleware"
	"go-fiber-template/lib/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/ryanbekhen/di"
)

type httpHandler struct {
	productService interfaces.ProductService
}

func NewHttpHandler(r fiber.Router) {
	productService := di.MustResolve[interfaces.ProductService]()
	authMiddleware := di.MustResolve[middleware.JWTAuthMiddleware]()

	handler := &httpHandler{
		productService: productService,
	}

	r.Post("/",
		authMiddleware.Validate(),
		middleware.Validate[dto.CreateProductRequest](middleware.PlacementBody),
		handler.Create,
	)
	r.Get("/",
		authMiddleware.Validate(),
		handler.FindAll,
	)
	r.Get("/:id",
		authMiddleware.Validate(),
		handler.FindByID,
	)
	r.Put("/:id",
		authMiddleware.Validate(),
		middleware.Validate[dto.UpdateProductRequest](middleware.PlacementBody),
		handler.Update,
	)
	r.Delete("/:id",
		authMiddleware.Validate(),
		handler.Delete,
	)
}

// @Summary		Create a new product
// @Description	Create a new product
// @Tags			Product
// @Accept			application/json
// @Produce		application/json
// @Param			request	body		dto.CreateProductRequest	true	"Create product request"
// @Success		201		{object}	dto.ResponseDto{data=dto.ProductDto}
// @Failure		400		{object}	dto.ResponseDto
// @Failure		401		{object}	dto.ResponseDto
// @Failure		500		{object}	dto.ResponseDto
// @Router			/products [post]
// @Security		Bearer
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

// @Summary		Find all products
// @Description	Find all products
// @Tags			Product
// @Accept			application/json
// @Produce		application/json
// @Success		200		{object}	dto.ResponseDto{data=[]dto.ProductDto}
// @Failure		400		{object}	dto.ResponseDto
// @Failure		401		{object}	dto.ResponseDto
// @Failure		500		{object}	dto.ResponseDto
// @Router			/products [get]
// @Security		Bearer
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

// @Summary		Find product by ID
// @Description	Find product by ID
// @Tags			Product
// @Accept			application/json
// @Produce		application/json
// @Param			id	path		string	true	"Product ID"
// @Success		200		{object}	dto.ResponseDto{data=dto.ProductDto}
// @Failure		400		{object}	dto.ResponseDto
// @Failure		401		{object}	dto.ResponseDto
// @Failure		500		{object}	dto.ResponseDto
// @Router			/products/:id [get]
// @Security		Bearer
func (h *httpHandler) FindByID(c *fiber.Ctx) error {
	idParam := c.Params("id", "")
	if idParam == "" {
		return fiber.ErrBadRequest
	}

	id, err := utils.ConvertStringToUint(idParam)
	if err != nil {
		return err
	}

	data, err := h.productService.FindByID(c, id)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(dto.ResponseDto{
		Message: "Product fetched successfully",
		Data:    data,
	})
}

// @Summary		Update a product
// @Description	Update a product
// @Tags			Product
// @Accept			application/json
// @Produce		application/json
// @Security		Bearer
// @Param			id	path		string	true	"Product ID"
// @Param			request	body		dto.UpdateProductRequest	true	"Update product request"
// @Success		200		{object}	dto.ResponseDto{data=dto.ProductDto}
// @Failure		400		{object}	dto.ResponseDto
// @Failure		401		{object}	dto.ResponseDto
// @Failure		500		{object}	dto.ResponseDto
// @Router			/products/:id [put]
// @Security		Bearer
func (h *httpHandler) Update(c *fiber.Ctx) error {
	idParam := c.Params("id", "")
	if idParam == "" {
		return fiber.ErrBadRequest
	}

	id, err := utils.ConvertStringToUint(idParam)
	if err != nil {
		return err
	}

	req := utils.ExtractStructFromValidator[dto.UpdateProductRequest](c)
	data, err := h.productService.Update(c, id, req)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(dto.ResponseDto{
		Message: "Product updated successfully",
		Data:    data,
	})
}

// @Summary		Delete a product
// @Description	Delete a product
// @Tags			Product
// @Accept			application/json
// @Produce		application/json
// @Security		Bearer
// @Param			id	path		string	true	"Product ID"
// @Success		200		{object}	dto.ResponseDto
// @Failure		400		{object}	dto.ResponseDto
// @Failure		401		{object}	dto.ResponseDto
// @Failure		500		{object}	dto.ResponseDto
// @Router			/products/:id [delete]
// @Security		Bearer
func (h *httpHandler) Delete(c *fiber.Ctx) error {
	idParam := c.Params("id", "")
	if idParam == "" {
		return fiber.ErrBadRequest
	}

	id, err := utils.ConvertStringToUint(idParam)
	if err != nil {
		return err
	}

	if err := h.productService.Delete(c, id); err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(dto.ResponseDto{
		Message: "Product deleted successfully",
	})
}
