package product

import (
	"go-fiber-template/internal/domain/dto"
	"go-fiber-template/internal/domain/entity"
	"go-fiber-template/internal/domain/interfaces"

	"github.com/gofiber/fiber/v2"
	"github.com/ryanbekhen/di"
)

type service struct {
	productRepo interfaces.ProductRepository
}

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

	return productToDto(product), nil
}

func (s *service) Delete(c *fiber.Ctx, id uint) error {
	if err := s.productRepo.Delete(id); err != nil {
		return err
	}

	return nil
}

func (s *service) FindAll(c *fiber.Ctx) ([]dto.ProductDto, error) {
	products, err := s.productRepo.FindAll()
	if err != nil {
		return nil, err
	}

	var productDtos []dto.ProductDto
	for _, product := range products {
		productDtos = append(productDtos, *productToDto(&product))
	}

	return productDtos, nil
}

func (s *service) FindByID(c *fiber.Ctx, id uint) (*dto.ProductDto, error) {
	product, err := s.productRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	return productToDto(product), nil
}

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

	return productToDto(product), nil
}

func productToDto(data *entity.Product) *dto.ProductDto {
	return &dto.ProductDto{
		ID:          data.ID,
		Name:        data.Name,
		Description: data.Description,
		Price:       data.Price,
		Stock:       data.Stock,
		CreatedAt:   data.CreatedAt,
		UpdatedAt:   data.UpdatedAt,
	}
}

func RegisterService() {
	productRepo := di.MustResolve[interfaces.ProductRepository]()

	di.RegisterFactory(func() interfaces.ProductService {
		return &service{
			productRepo: productRepo,
		}
	})
}
