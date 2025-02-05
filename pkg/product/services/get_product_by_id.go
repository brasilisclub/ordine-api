package services

import (
	"errors"
	"fmt"
	"ordine-api/pkg/database"
	"ordine-api/pkg/product"

	"gorm.io/gorm"
)

func GetProductById(id string) (product.Product, error) {
	var product product.Product

	db := database.GetConnector()

	result := db.First(&product, "id = ?", id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return product, gorm.ErrRecordNotFound
		}
		return product, errors.New(fmt.Sprintf("Error getting product: %s", result.Error.Error()))
	}

	if result.RowsAffected == 0 {
		return product, errors.New("No product found with the provided ID")
	}

	return product, nil
}
