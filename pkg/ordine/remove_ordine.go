package ordine

import (
	"errors"
	"fmt"
	"ordine-api/pkg/database"
)

func RemoveOrdine(id string) error {
	db := database.GetConnector()

	result := db.Delete(&Ordine{}, id)
	if result.Error != nil {
		return errors.New(fmt.Sprintf("Error deleting ordine: %s", result.Error))
	}

	if result.RowsAffected == 0 {
		return errors.New("No ordine found with the provided ID")
	}
	return nil
}
