package services

import (
	"errors"
	"fmt"
	"ordine-api/pkg/database"
	"ordine-api/pkg/product"

	"gorm.io/gorm"
)

func UpdateProduct(id string, p *product.ProductRequestBody) (*product.Product, error) {
	db := database.GetConnector()

	dbProduct, err := GetProductById(id)
	if err != nil {
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
