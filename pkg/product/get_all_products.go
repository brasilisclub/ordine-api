package product

import (
	"errors"
	"fmt"
	"ordine-api/pkg/database"
)

func getAllProducts() ([]Product, error) {
	var products []Product
	db, err := database.GetConnector()
	if err != nil {
		return products, errors.New(fmt.Sprintf("Error trying to connect to database %s", err.Error()))
	}

	result := db.Find(&products)

	return products, result.Error
}
