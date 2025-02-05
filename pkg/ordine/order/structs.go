package order

import (
	"ordine-api/pkg/product"
	"time"
)

type OrderProducts struct {
	ID        uint            `json:"id" bson:"id"`
	CreatedAt time.Time       `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time       `json:"updated_at" bson:"updated_at"`
	DeletedAt *time.Time      `json:"deleted_at,omitempty" bson:"deleted_at,omitempty"`
	OrdineID  uint            `json:"ordine_id" bson:"ordine_id"`
	ProductID uint            `json:"product_id" bson:"product_id"`
	Quantity  int             `json:"quantity" bson:"quantity"`
	Product   product.Product `gorm:"foreignKey:ProductID" json:"product" bson:"product,omitempty"`
}

type OrderProductBody struct {
	ProductID int `json:"product_id" binding:"required"`
	Quantity  int `json:"quantity" binding:"required"`
}
