package service

import (
	"github.com/JPSBarbosa/GO-API-REST/internal/models"
	"github.com/JPSBarbosa/GO-API-REST/internal/product"
)

type productService struct {
	repo product.Repository
}

func NewProductService(repo product.Repository) product.Service {
	return &productService{repo: repo}
}

func (s *productService) FindById(id uint) (*models.Product, error) {
	return s.repo.FindById(id)
}

func (s *productService) Save(product models.Product) (uint, error) {
	return s.repo.Save(product)
}

func (s *productService) Update(product models.Product) (*models.Product, error) {
	return s.repo.Update(product)
}

func (s *productService) DeleteById(id uint) error {
	return s.repo.DeleteById(id)
}
