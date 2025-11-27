package product

import (
	"go-fiber-template/internal/domain/entity"
	"go-fiber-template/internal/domain/interfaces"

	"github.com/ryanbekhen/di"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func (r *repository) Create(data *entity.Product) error {
	return r.db.Create(data).Error
}

func (r *repository) Delete(id uint) error {
	return r.db.Delete(&entity.Product{}, id).Error
}

func (r *repository) FindAll() ([]entity.Product, error) {
	var products []entity.Product
	if err := r.db.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (r *repository) FindByID(id uint) (*entity.Product, error) {
	var product entity.Product
	if err := r.db.First(&product, id).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *repository) Update(data *entity.Product) error {
	return r.db.Save(data).Error
}

func RegisterRepository() {
	db := di.MustResolve[*gorm.DB]()

	di.RegisterFactory(func() interfaces.ProductRepository {
		return &repository{
			db: db,
		}
	})
}
