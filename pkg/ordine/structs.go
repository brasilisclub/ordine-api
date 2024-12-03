package ordine

import "gorm.io/gorm"

type Ordine struct {
	gorm.Model
	Table      int       `json:"table"`
	ClientName string    `json:"client_name"`
	Status     bool      `json:"status"`
	Products   []Product `json:"products" gorm:"many2many:ordine_products;"`
}
type Product struct {
	gorm.Model
	Name        string  `json:"name"`
	Category    string  `json:"category"`
	Price       float32 `json:"price"`
	Description string  `json:"description"`
	Stock       int     `json:"stock"`
}
