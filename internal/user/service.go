package user

import (
	"github.com/JPSBarbosa/GO-API-REST/internal/models"
)

type Service interface {
	FindById(id uint) (*models.User, error)
	Save(product models.User) (uint, error)
	Update(product models.User) (*models.User, error)
	DeleteById(id uint) error
}
