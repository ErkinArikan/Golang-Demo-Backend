package models

import (
	"gorm.io/gorm"
)

// Product modelinin tanımı
type Product struct {
	gorm.Model
	Name     string  `json:"name"`
	Price    float32 `json:"price"`
	Discount float32 `json:"discount"`
	Store    string  `json:"store"`
}
