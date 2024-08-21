package services

import (
	"restAPIPractice/models"
	"restAPIPractice/repositories"
)

type ProductService interface {
	GetProductByID(id uint) (models.Product, error)
	GetAllProducts() ([]models.Product, error)
	CreateProduct(product models.Product) error
	UpdateProduct(product models.Product) error
	DeleteProduct(id uint) error
}

type productService struct {
	productRepo repositories.ProductRepository
}

// NewProductService, yeni bir ProductService nesnesi döndürür.
func NewProductService(repo repositories.ProductRepository) ProductService {
	return &productService{
		productRepo: repo,
	}
}

// GetProductByID, belirli bir ürün ID'sine sahip ürünü döndürür.
func (s *productService) GetProductByID(id uint) (models.Product, error) {
	return s.productRepo.GetByID(id)
}

// GetAllProducts, tüm ürünleri döndürür.
func (s *productService) GetAllProducts() ([]models.Product, error) {
	return s.productRepo.GetAll()
}

// CreateProduct, yeni bir ürün ekler.
func (s *productService) CreateProduct(product models.Product) error {
	return s.productRepo.Create(product)
}

// UpdateProduct, mevcut ürünü günceller.
func (s *productService) UpdateProduct(product models.Product) error {
	return s.productRepo.Update(product)
}

// DeleteProduct, belirli bir ürün ID'sine sahip ürünü siler.
func (s *productService) DeleteProduct(id uint) error {
	return s.productRepo.Delete(id)
}
