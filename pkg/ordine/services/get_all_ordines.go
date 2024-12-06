package services

import (
	"ordine-api/pkg/database"
	ord "ordine-api/pkg/ordine"
)

func GetAllOrdines() ([]ord.Ordine, error) {
	var ordines []ord.Ordine
	db := database.GetConnector()

	result := db.Preload("Products.Product").Find(&ordines)

	return ordines, result.Error
}
