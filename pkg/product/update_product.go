package product

import (
	"errors"
	"fmt"
	"ordine-api/pkg/database"

	"gorm.io/gorm"
)

func updateProduct(id string, p *Product) (*Product, error) {
	db, err := database.GetConnector()
	if err != nil {
		return nil, fmt.Errorf("Error trying to connect to database: %w", err)
	}

	var dbProduct Product
	if err := db.First(&dbProduct, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("Product with id %s not founded", id)
		}
		return nil, fmt.Errorf("Error getting product on db: %w", err)
	}

	dbProduct.Name = p.Name
	dbProduct.Category = p.Category
	dbProduct.Price = p.Price
	dbProduct.Description = p.Description
	dbProduct.Stock = p.Stock

	if err := db.Save(&dbProduct).Error; err != nil {
		return nil, fmt.Errorf("Error trying update product: %w", err)
	}

	return &dbProduct, nil
}
