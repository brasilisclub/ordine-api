package ordine

import (
	"errors"
	"fmt"
	"ordine-api/pkg/database"
)

func createOrdine(o *Ordine) error {
	db, err := database.GetConnector()
	if err != nil {
		return errors.New(fmt.Sprintf("Error trying to connect to database %s", err.Error()))
	}

	result := db.Create(o)
	if result.Error != nil {
		return errors.New(fmt.Sprintf("Error trying to insert ordine on database %s", result.Error.Error()))
	}

	return nil
}
