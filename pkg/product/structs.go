package product

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name        string  `json:"name"`
	Category    string  `json:"category"`
	Price       float32 `json:"price"`
	Description string  `json:"description"`
	Stock       int     `json:"stock"`
}
