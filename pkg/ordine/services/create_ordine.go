package services

import (
	"fmt"
	"ordine-api/pkg/database"
	ord "ordine-api/pkg/ordine"
)

func CreateOrdine(o *ord.OrdineRequestBody) (*ord.Ordine, error) {
	db := database.GetConnector()

	var dbOrdine ord.Ordine

	dbOrdine.Table = o.Table
	dbOrdine.ClientName = o.ClientName
	dbOrdine.Status = o.Status

	result := db.Create(&dbOrdine)
	if result.Error != nil {
		return &ord.Ordine{}, fmt.Errorf("error trying to insert ordine on database %s", result.Error.Error())
	}

	return &dbOrdine, nil
}
