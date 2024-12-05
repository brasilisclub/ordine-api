package ordine

import (
	"ordine-api/pkg/database"
)

func getAllOrdines() ([]Ordine, error) {
	var ordines []Ordine
	db := database.GetConnector()

	result := db.Preload("Products").Find(&ordines)

	return ordines, result.Error
}