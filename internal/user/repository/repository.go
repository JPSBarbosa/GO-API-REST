package repository

import (
	"github.com/JPSBarbosa/GO-API-REST/internal/models"
	"github.com/JPSBarbosa/GO-API-REST/internal/user"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) user.Repository {
	return &userRepository{db: db}
}

func (t *userRepository) FindById(id uint) (*models.User, error) {
	user := new(models.User)
	resp := t.db.First(&user, id)
	if resp.Error != nil {
		return nil, resp.Error
	}
	return user, nil
}

func (t *userRepository) Save(user models.User) (uint, error) {
	result := t.db.Create(&user)
	if result.Error != nil {
		return 0, result.Error
	}
	return user.ID, nil
}

func (t *userRepository) Update(user models.User) (*models.User, error) {
	exist := new(models.User)
	result := t.db.First(&exist, user.ID)
	if result.Error != nil {
		return nil, result.Error
	}
	exist.Name = user.Name
	exist.Email = user.Email
	exist.BirthDate = user.BirthDate

	resp := t.db.Save(&exist)
	if resp.Error != nil {
		return nil, resp.Error
	}

	return exist, nil
}

func (t *userRepository) DeleteById(id uint) error {
	result := t.db.Delete(&models.User{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
