// service/product_service.go
package service

import (
	"ecommerce-api/models"
	"ecommerce-api/repositories"
	"errors"
)



type ProductService struct {
	ProductRepo *repositories.ProductRepository
}

func NewProductService(productRepo *repositories.ProductRepository) *ProductService {
	return &ProductService{ProductRepo: productRepo}
}

func (s *ProductService) CreateProduct(product *models.Product) error {
	return s.ProductRepo.Create(product)
}

func (s *ProductService) UpdateProduct(product *models.Product) error {
	// Check if product exists
	_, err := s.ProductRepo.FindByID(product.ID)
	if err != nil {
		return errors.New("product not found")
	}
	return s.ProductRepo.Update(product)
}

func (s *ProductService) DeleteProduct(id uint) error {
	// Check if product exists
	_, err := s.ProductRepo.FindByID(id)
	if err != nil {
		return errors.New("product not found")
	}
	return s.ProductRepo.Delete(id)
}

func (s *ProductService) GetProduct(id uint) (*models.Product, error) {
	return s.ProductRepo.FindByID(id)
}

func (s *ProductService) GetProducts(page, limit int) ([]models.Product, int64, error) {
	// Handle pagination
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}

	products, err := s.ProductRepo.FindAll(page, limit)
	if err != nil {
		return nil, 0, err
	}

	total, err := s.ProductRepo.Count()
	if err != nil {
		return nil, 0, err
	}

	return products, total, nil
}