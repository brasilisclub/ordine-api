package ordine

import (
	"errors"
	"fmt"
	"ordine-api/pkg/database"
)

func CreateOrdine(o *Ordine) error {
	db := database.GetConnector()

	result := db.Create(o)
	if result.Error != nil {
		return errors.New(fmt.Sprintf("Error trying to insert ordine on database %s", result.Error.Error()))
	}

	return nil
}
