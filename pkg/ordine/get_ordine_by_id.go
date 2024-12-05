package ordine

import (
	"errors"
	"fmt"
	"ordine-api/pkg/database"
)

func getOrdineById(id string) (Ordine, error) {
	var ordine Ordine

	db, err := database.GetConnector()
	if err != nil {
		return ordine, errors.New(fmt.Sprintf("Error trying to connect to database %s", err.Error()))
	}

	result := db.First(&ordine, "id = ?", id)

	if result.Error != nil {
		return ordine, errors.New(fmt.Sprintf("Error getting ordine: %s", result.Error.Error()))
	}

	if result.RowsAffected == 0 {
		return ordine, errors.New("No ordine found with the provided ID")
	}

	return ordine, nil
}
