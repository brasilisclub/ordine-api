package services

import (
	"errors"
	"fmt"
	"ordine-api/pkg/database"
	ord "ordine-api/pkg/ordine"

	"gorm.io/gorm"
)

func GetOrdineById(id string) (ord.Ordine, error) {
	var ordine ord.Ordine

	db := database.GetConnector()

	result := db.Preload("Products.Product").First(&ordine, "id = ?", id)

	if result.Error != nil {
		if errors.Is(gorm.ErrRecordNotFound, result.Error) {
			return ordine, ord.ErrorOrdineNotFound
		}
		return ordine, fmt.Errorf("internal error getting ordine: %s", result.Error.Error())
	}

	return ordine, nil
}
