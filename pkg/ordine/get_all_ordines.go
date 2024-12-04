package ordine

import (
	"errors"
	"fmt"
	"ordine-api/pkg/database"
)

func getAllOrdines() ([]Ordine, error) {
	var ordines []Ordine
	db, err := database.GetConnector()
	if err != nil {
		return ordines, errors.New(fmt.Sprintf("Error trying to connect to database %s", err.Error()))
	}

	result := db.Find(&ordines)

	return ordines, result.Error
}
