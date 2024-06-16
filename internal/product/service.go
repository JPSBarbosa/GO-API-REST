package product

import "github.com/JPSBarbosa/GO-API-REST/internal/models"

type Service interface {
	FindById(id uint) (*models.Product, error)
	Save(product models.Product) (uint, error)
	Update(product models.Product) (*models.Product, error)
	DeleteById(id uint) error
}
