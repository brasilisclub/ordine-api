package product

import (
	"ordine-api/pkg/database"
)

func getAllProducts() ([]Product, error) {
	var products []Product
	db := database.GetConnector()

	result := db.Find(&products)

	return products, result.Error
}
