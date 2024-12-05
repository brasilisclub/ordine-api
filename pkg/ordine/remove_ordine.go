package ordine

import (
	"errors"
	"fmt"
	"ordine-api/pkg/database"
)

func removeOrdine(id string) error {
	db, err := database.GetConnector()
	if err != nil {
		return errors.New(fmt.Sprintf("Error trying to connect to database %s", err.Error()))
	}

	result := db.Delete(&Ordine{}, id)
	if result.Error != nil {
		return errors.New(fmt.Sprintf("Error deleting ordine: %s", result.Error))
	}

	if result.RowsAffected == 0 {
		return errors.New("No ordine found with the provided ID")
	}
	return nil
}
