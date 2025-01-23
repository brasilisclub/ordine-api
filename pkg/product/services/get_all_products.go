package services

import (
	"ordine-api/pkg/database"
	"ordine-api/pkg/product"
)

func GetAllProducts() ([]product.Product, error) {
	var products []product.Product
	db := database.GetConnector()

	result := db.Find(&products)

	return products, result.Error
}
