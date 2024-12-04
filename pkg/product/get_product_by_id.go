package product

import (
	"errors"
	"fmt"
	"ordine-api/pkg/database"
)

func getProductById(id string) (Product, error) {
	var product Product

	db, err := database.GetConnector()
	if err != nil {
		return product, errors.New(fmt.Sprintf("Error trying to connect to database %s", err.Error()))
	}

	result := db.First(&product, id)
	if result.Error != nil {
		return product, errors.New(fmt.Sprintf("Error deleting product: %s", result.Error.Error()))
	}

	if result.RowsAffected == 0 {
		return product, errors.New("No product found with the provided ID")
	}

	return product, nil
}
