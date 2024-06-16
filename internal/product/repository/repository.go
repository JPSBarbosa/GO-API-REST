package repository

import (
	"github.com/JPSBarbosa/GO-API-REST/internal/models"
	"github.com/JPSBarbosa/GO-API-REST/internal/product"
	"gorm.io/gorm"
)

type productsService struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) product.Repository {
	return &productsService{db: db}
}

func (s *productsService) FindById(id uint) (*models.Product, error) {
	product := new(models.Product)
	resp := s.db.First(&product, id)
	if resp.Error != nil {
		return nil, resp.Error
	}
	return product, nil
}

func (s *productsService) Save(product models.Product) (uint, error) {
	result := s.db.Create(&product)
	if result.Error != nil {
		return 0, result.Error
	}
	return product.ID, nil
}

func (s *productsService) Update(product models.Product) (*models.Product, error) {
	exist := new(models.Product)
	result := s.db.First(&exist, product.ID)
	if result.Error != nil {
		return nil, result.Error
	}
	exist.Name = product.Name
	exist.Description = product.Description
	exist.Price = product.Price
	exist.Quantity = product.Quantity

	resp := s.db.Save(&exist)
	if resp.Error != nil {
		return nil, resp.Error
	}

	return exist, nil
}

func (s *productsService) DeleteById(id uint) error {
	result := s.db.Delete(&models.Product{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
