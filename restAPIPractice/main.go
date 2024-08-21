package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"restAPIPractice/controllers"
	"restAPIPractice/models" // models paketini ekle
	"restAPIPractice/repositories"
	"restAPIPractice/services"
)

func main() {
	// Gin router'ını oluştur
	router := gin.Default()

	// Veritabanı bağlantısını oluştur
	dsn := "host=localhost user=postgres password=12345 dbname=demoDB port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	// Model migrasyonları
	err = db.AutoMigrate(&models.Product{})
	if err != nil {
		panic("Failed to migrate database")
	}

	// Repository ve Service'i oluştur
	productRepo := repositories.NewProductRepository(db)
	productService := services.NewProductService(productRepo)
	productController := controllers.NewProductController(productService)

	// Rotayı kaydet
	productController.RegisterRoutes(router)

	// Sunucuyu başlat
	router.Run(":8080")
}

//
