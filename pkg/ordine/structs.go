package ordine

import (
	"ordine-api/pkg/product"

	"gorm.io/gorm"
)

type Ordine struct {
	gorm.Model
	Table      int               `json:"table"`
	ClientName string            `json:"client_name"`
	Status     bool              `json:"status"`
	Products   []product.Product `json:"products" gorm:"many2many:ordine_products;"`
}
