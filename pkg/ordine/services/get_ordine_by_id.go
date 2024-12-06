package services

import (
	"errors"
	"fmt"
	"ordine-api/pkg/database"
	ord "ordine-api/pkg/ordine"
)

func GetOrdineById(id string) (ord.Ordine, error) {
	var ordine ord.Ordine

	db := database.GetConnector()

	result := db.Preload("Products.Product").First(&ordine, "id = ?", id)

	if result.Error != nil {
		return ordine, fmt.Errorf("error getting ordine: %s", result.Error.Error())
	}

	if result.RowsAffected == 0 {
		return ordine, errors.New("no ordine found with the provided ID")
	}

	return ordine, nil
}
