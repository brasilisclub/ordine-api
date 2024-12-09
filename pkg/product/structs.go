package product

import "time"

type Product struct {
	ID          uint       `json:"id" bson:"id"`
	CreatedAt   time.Time  `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at" bson:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty" bson:"deleted_at,omitempty"`
	Name        string     `json:"name" bson:"name"`
	Category    string     `json:"category" bson:"category"`
	Price       float32    `json:"price" bson:"price"`
	Description string     `json:"description" bson:"description"`
	Stock       int        `json:"stock" bson:"stock"`
}
type ProductRequestBody struct {
	Name        string  `json:"name" bson:"name"`
	Category    string  `json:"category" bson:"category"`
	Price       float32 `json:"price" bson:"price"`
	Description string  `json:"description" bson:"description"`
	Stock       int     `json:"stock" bson:"stock"`
}
