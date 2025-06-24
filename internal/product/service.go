package product

import (
	"go-fiber-template/internal/domain/dto"
	"go-fiber-template/internal/domain/entity"
	"go-fiber-template/internal/domain/interfaces"

	"github.com/gofiber/fiber/v2"
)

type service struct {
	productRepo interfaces.ProductRepository
}

// Create implements interfaces.ProductService.
func (s *service) Create(c *fiber.Ctx, req *dto.CreateProductRequest) (*dto.ProductDto, error) {
	product := &entity.Product{
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Stock:       req.Stock,
	}

	if err := s.productRepo.Create(product); err != nil {
		return nil, err
	}

	return constructProductDto(product), nil
}

// Delete implements interfaces.ProductService.
func (s *service) Delete(c *fiber.Ctx, id uint) error {
	if err := s.productRepo.Delete(id); err != nil {
		return err
	}

	return nil
}

// FindAll implements interfaces.ProductService.
func (s *service) FindAll(c *fiber.Ctx) ([]dto.ProductDto, error) {
	products, err := s.productRepo.FindAll()
	if err != nil {
		return nil, err
	}

	var productDtos []dto.ProductDto
	for _, product := range products {
		productDtos = append(productDtos, *constructProductDto(&product))
	}

	return productDtos, nil
}

// FindByID implements interfaces.ProductService.
func (s *service) FindByID(c *fiber.Ctx, id uint) (*dto.ProductDto, error) {
	product, err := s.productRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	return constructProductDto(product), nil
}

// Update implements interfaces.ProductService.
func (s *service) Update(c *fiber.Ctx, id uint, req *dto.UpdateProductRequest) (*dto.ProductDto, error) {
	product, err := s.productRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	product.Name = req.Name
	product.Description = req.Description
	product.Price = req.Price
	product.Stock = req.Stock

	if err := s.productRepo.Update(product); err != nil {
		return nil, err
	}

	return constructProductDto(product), nil
}

func constructProductDto(product *entity.Product) *dto.ProductDto {
	return &dto.ProductDto{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Stock:       product.Stock,
		CreatedAt:   product.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:   product.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}

func NewService(productRepo interfaces.ProductRepository) interfaces.ProductService {
	return &service{productRepo: productRepo}
}
