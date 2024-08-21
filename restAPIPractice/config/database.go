package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

// DBConnect veritabanı bağlantısını başlatır
func DBConnect() *gorm.DB {
	dsn := "host=localhost user=postgres password=12345 dbname=demoDB port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	return db
}
