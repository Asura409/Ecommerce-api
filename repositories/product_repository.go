// repository/product_repository.go
package repositories

import (
	"ecommerce-api/models"
	"gorm.io/gorm"
)


type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) Create(product *models.Product) error {
	return r.db.Create(product).Error
}

func (r *ProductRepository) Update(product *models.Product) error {
	return r.db.Save(product).Error
}

func (r *ProductRepository) Delete(id uint) error {
	return r.db.Delete(&models.Product{}, id).Error
}

func (r *ProductRepository) FindByID(id uint) (*models.Product, error) {
	var product models.Product
	err := r.db.First(&product, id).Error
	return &product, err
}

func (r *ProductRepository) FindAll(page, limit int) ([]models.Product, error) {
	var products []models.Product
	offset := (page - 1) * limit
	err := r.db.Offset(offset).Limit(limit).Find(&products).Error
	return products, err
}

func (r *ProductRepository) Count() (int64, error) {
	var count int64
	err := r.db.Model(&models.Product{}).Count(&count).Error
	return count, err
}