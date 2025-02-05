package ordine

import (
	"ordine-api/pkg/ordine/order"
	"time"
)

type OrdineRequestBody struct {
	Table      int    `json:"table" bson:"tables" binding:"required"`
	ClientName string `json:"client_name" bson:"client_name" binding:"required"`
	Status     bool   `json:"status" bson:"status" binding:"required"`
}
type Ordine struct {
	ID         uint                  `json:"id" bson:"id" gorm:"primaryKey"`
	CreatedAt  time.Time             `json:"created_at" bson:"created_at"`
	UpdatedAt  time.Time             `json:"updated_at" bson:"updated_at"`
	DeletedAt  *time.Time            `json:"deleted_at,omitempty" bson:"deleted_at,omitempty"`
	Table      int                   `json:"table" bson:"tables"`
	ClientName string                `json:"client_name" bson:"client_name"`
	Status     bool                  `json:"status" bson:"status"`
	Products   []order.OrderProducts `json:"ordine_products" gorm:"foreignKey:OrdineID" bson:"products,omitempty"`
}
