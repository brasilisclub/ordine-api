package services

import (
	"fmt"
	"ordine-api/pkg/database"
	"ordine-api/pkg/ordine"
	ord "ordine-api/pkg/ordine"
)

func RemoveOrdine(id string) error {
	db := database.GetConnector()

	result := db.Where("id = ?", id).Delete(&ord.Ordine{})

	if result.Error != nil {
		return fmt.Errorf("error deleting ordine: %s", result.Error)
	}

	if result.RowsAffected == 0 {
		return ordine.ErrorOrdineNotFound
	}
	return nil
}
