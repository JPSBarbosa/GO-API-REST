package Service

import (
	"github.com/JPSBarbosa/GO-API-REST/Models"
	"gorm.io/gorm"
)

type ProductsService struct {
	db *gorm.DB
}

func NewProductService(db *gorm.DB) *ProductsService {
	return &ProductsService{
		db: db,
	}
}

func (s *ProductsService) FindById(id uint) (Models.Products, error) {
	product := new(Models.Products)
	resp := s.db.First(&product, id)
	if resp.Error != nil {
		return Models.Products{}, resp.Error
	}
	return *product, nil
}

func (s *ProductsService) SaveProduct(product Models.Products) (uint, error) {
	result := s.db.Create(&product)
	if result.Error != nil {
		return 0, result.Error
	}
	return product.ID, nil
}

func (s *ProductsService) UpdateProduct(product Models.Products, id uint) (Models.Products, error) {
	exist := new(Models.Products)
	result := s.db.First(&exist, id)
	if result.Error != nil {
		return Models.Products{}, result.Error
	}
	exist.Nome = product.Nome
	exist.Descricao = product.Nome
	exist.Preco = product.Preco
	exist.Quantidade = product.Quantidade

	resp := s.db.Save(&exist)
	if resp.Error != nil {
		return Models.Products{}, resp.Error
	}

	return *exist, nil
}

func (s *ProductsService) DeleteById(id uint) error {
	result := s.db.Delete(&Models.Products{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
