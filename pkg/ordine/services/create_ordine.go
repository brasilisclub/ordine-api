package services

import (
	"fmt"
	"ordine-api/pkg/database"
	ord "ordine-api/pkg/ordine"
)

func CreateOrdine(o *ord.Ordine) error {
	db := database.GetConnector()

	result := db.Create(o)
	if result.Error != nil {
		return fmt.Errorf("error trying to insert ordine on database %s", result.Error.Error())
	}

	return nil
}
