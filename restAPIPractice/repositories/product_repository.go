package repositories

import (
	"gorm.io/gorm"
	"restAPIPractice/models"
)

type ProductRepository interface {
	GetAll() ([]models.Product, error)
	GetByID(id uint) (models.Product, error)
	Create(product models.Product) error
	Update(product models.Product) error
	Delete(id uint) error
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{
		db: db,
	}
}

// GetAll tüm ürünleri döndürür
func (r *productRepository) GetAll() ([]models.Product, error) {
	var products []models.Product
	result := r.db.Find(&products)
	return products, result.Error
}

// GetByID belirli bir ürün ID'sine sahip ürünü döndürür
func (r *productRepository) GetByID(id uint) (models.Product, error) {
	var product models.Product
	result := r.db.First(&product, id)
	return product, result.Error
}

// Create yeni bir ürün ekler
func (r *productRepository) Create(product models.Product) error {
	result := r.db.Create(&product)
	return result.Error
}

// Update mevcut ürünü günceller
func (r *productRepository) Update(product models.Product) error {
	result := r.db.Save(&product)
	return result.Error
}

// Delete belirli bir ürün ID'sine sahip ürünü siler
func (r *productRepository) Delete(id uint) error {
	result := r.db.Delete(&models.Product{}, id)
	return result.Error
}
