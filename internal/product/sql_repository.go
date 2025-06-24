package product

import (
	"go-fiber-template/internal/domain/entity"
	"go-fiber-template/internal/domain/interfaces"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

// Create implements interfaces.ProductRepository.
func (r *repository) Create(data *entity.Product) error {
	return r.db.Create(data).Error
}

// Delete implements interfaces.ProductRepository.
func (r *repository) Delete(id uint) error {
	return r.db.Delete(&entity.Product{}, id).Error
}

// FindAll implements interfaces.ProductRepository.
func (r *repository) FindAll() ([]entity.Product, error) {
	var products []entity.Product
	if err := r.db.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

// FindByID implements interfaces.ProductRepository.
func (r *repository) FindByID(id uint) (*entity.Product, error) {
	var product entity.Product
	if err := r.db.First(&product, id).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

// Update implements interfaces.ProductRepository.
func (r *repository) Update(data *entity.Product) error {
	return r.db.Save(data).Error
}

func NewRepository(db *gorm.DB) interfaces.ProductRepository {
	return &repository{db: db}
}
