package ordine

import (
	"ordine-api/pkg/product"
	"time"
)

type OrdineRequestBody struct {
	Table      int    `json:"table" bson:"tables"`
	ClientName string `json:"client_name" bson:"client_name"`
	Status     bool   `json:"status" bson:"status"`
}
type Ordine struct {
	ID         uint            `json:"id" bson:"id"`
	CreatedAt  time.Time       `json:"created_at" bson:"created_at"`
	UpdatedAt  time.Time       `json:"updated_at" bson:"updated_at"`
	DeletedAt  *time.Time      `json:"deleted_at,omitempty" bson:"deleted_at,omitempty"`
	Table      int             `json:"table" bson:"tables"`
	ClientName string          `json:"client_name" bson:"client_name"`
	Status     bool            `json:"status" bson:"status"`
	Products   []OrderProducts `json:"ordine_products" gorm:"foreignKey:OrdineID" bson:"products,omitempty"`
}

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
	ProductID int `json:"product_id"`
	Quantity  int `json:"quantity"`
}
