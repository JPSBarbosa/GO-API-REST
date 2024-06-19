package service

import (
	"github.com/JPSBarbosa/GO-API-REST/internal/models"
	"github.com/JPSBarbosa/GO-API-REST/internal/product"
)

// productService implements the product.Service interface
type productService struct {
	repo product.Repository
}

// NewProductService creates a new instance of productService
func NewProductService(repo product.Repository) product.Service {
	return &productService{repo: repo}
}

// FindById finds a product by ID
// Delegates the call to the repository's FindById method
func (s *productService) FindById(id uint) (*models.Product, error) {
	return s.repo.FindById(id)
}

// Save saves a new product
// Delegates the call to the repository's Save method
func (s *productService) Save(product models.Product) (uint, error) {
	return s.repo.Save(product)
}

// Update updates an existing product
// Delegates the call to the repository's Update method
func (s *productService) Update(product models.Product) (*models.Product, error) {
	return s.repo.Update(product)
}

// DeleteById deletes a product by ID
// Delegates the call to the repository's DeleteById method
func (s *productService) DeleteById(id uint) error {
	return s.repo.DeleteById(id)
}
