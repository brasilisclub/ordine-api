package ordine

import (
	"errors"
	"fmt"
	"ordine-api/pkg/database"
)

func createProduct(p *Product) error {
	db, err := database.GetConnector()
	if err != nil {
		return errors.New(fmt.Sprintf("Error trying to connect to database %s", err.Error()))
	}

	result := db.Create(p)
	if result.Error != nil {
		return errors.New(fmt.Sprintf("Error trying to insert product on database %s", result.Error.Error()))
	}

	return err
}
