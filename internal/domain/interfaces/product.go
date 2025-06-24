package interfaces

import (
	"go-fiber-template/internal/domain/dto"
	"go-fiber-template/internal/domain/entity"

	"github.com/gofiber/fiber/v2"
)

type ProductRepository interface {
	Create(data *entity.Product) error
	FindByID(id uint) (*entity.Product, error)
	FindAll() ([]entity.Product, error)
	Update(data *entity.Product) error
	Delete(id uint) error
}

type ProductService interface {
	Create(c *fiber.Ctx, req *dto.CreateProductRequest) (*dto.ProductDto, error)
	FindByID(c *fiber.Ctx, id uint) (*dto.ProductDto, error)
	FindAll(c *fiber.Ctx) ([]dto.ProductDto, error)
	Update(c *fiber.Ctx, id uint, req *dto.UpdateProductRequest) (*dto.ProductDto, error)
	Delete(c *fiber.Ctx, id uint) error
}
