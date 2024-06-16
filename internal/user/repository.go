package user

import "github.com/JPSBarbosa/GO-API-REST/internal/models"

type Repository interface {
	FindById(id uint) (*models.User, error)
	Save(user models.User) (uint, error)
	Update(user models.User) (*models.User, error)
	DeleteById(id uint) error
}
