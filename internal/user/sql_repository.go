package user

import (
	"go-fiber-template/internal/domain/entity"
	"go-fiber-template/internal/domain/interfaces"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func (r *repository) Create(data *entity.User) error {
	return r.db.Create(data).Error
}

func (r *repository) Delete(id uint) error {
	return r.db.Delete(&entity.User{}, id).Error
}

func (r *repository) FindByEmail(email string) (*entity.User, error) {
	var user entity.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *repository) FindByID(id uint) (*entity.User, error) {
	var user entity.User
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *repository) Update(user *entity.User) error {
	return r.db.Save(user).Error
}

func NewRepository(db *gorm.DB) interfaces.UserRepository {
	return &repository{db: db}
}
