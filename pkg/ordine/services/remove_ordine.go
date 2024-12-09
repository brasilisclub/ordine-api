package services

import (
	"errors"
	"fmt"
	"ordine-api/pkg/database"
	ord "ordine-api/pkg/ordine"
)

func RemoveOrdine(id string) error {
	db := database.GetConnector()

	result := db.Delete(&ord.Ordine{}, id)
	if result.Error != nil {
		return fmt.Errorf("error deleting ordine: %s", result.Error)
	}

	if result.RowsAffected == 0 {
		return errors.New("no ordine found with the provided ID")
	}
	return nil
}
