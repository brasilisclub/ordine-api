package product

import (
	"errors"
	"fmt"
	"ordine-api/pkg/database"
	"strconv"
)

func removeProduct(id string) error {
	productId, err := strconv.Atoi(id)
	if err != nil {
		return errors.New(fmt.Sprintf("Invalid param: %s", err.Error()))
	}
	db, err := database.GetConnector()
	if err != nil {
		return errors.New(fmt.Sprintf("Error trying to connect to database %s", err.Error()))
	}

	result := db.Delete(&Product{}, productId)
	if result.Error != nil {
		return errors.New(fmt.Sprintf("Error deleting product: %s", result.Error.Error()))
	}

	if result.RowsAffected == 0 {
		return errors.New("No product found with the provided ID")
	}

	return nil

}