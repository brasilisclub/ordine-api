package ordine

import (
	"ordine-api/pkg/product"

	"gorm.io/gorm"
)

type Ordine struct {
	gorm.Model
	Table      int             `json:"table"`
	ClientName string          `json:"client_name"`
	Status     bool            `json:"status"`
	Products   []OrderProducts `json:"ordine_products" gorm:"foreignKey:OrdineID"`
}

type OrderProducts struct {
	gorm.Model
	OrdineID  uint            `json:"ordine_id"`
	ProductID uint            `json:"product_id"`
	Quantity  int             `json:"quantity"`
	Product   product.Product `gorm:"foreignKey:ProductID" json:"product"`
}

type OrderProductBody struct {
	ProductID int `json:"product_id"`
	Quantity  int `json:"quantity"`
}
