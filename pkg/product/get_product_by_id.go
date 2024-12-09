package product

import (
	"errors"
	"fmt"
	"ordine-api/pkg/database"
)

func GetProductById(id string) (Product, error) {
	var product Product

	db := database.GetConnector()

	result := db.First(&product, "id = ?", id)
	if result.Error != nil {
		return product, errors.New(fmt.Sprintf("Error getting product: %s", result.Error.Error()))
	}

	if result.RowsAffected == 0 {
		return product, errors.New("No product found with the provided ID")
	}

	return product, nil
}
