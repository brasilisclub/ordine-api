package product

import (
	"errors"
	"fmt"
	"ordine-api/pkg/database"
)

func createProduct(p *Product) error {
	db := database.GetConnector()

	result := db.Create(p)
	if result.Error != nil {
		return errors.New(fmt.Sprintf("Error trying to insert product on database %s", result.Error.Error()))
	}

	return nil
}
