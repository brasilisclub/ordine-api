package services

import (
	"errors"
	"fmt"
	"ordine-api/pkg/database"
	"ordine-api/pkg/product"
	"strconv"
)

func RemoveProduct(id string) error {
	productId, err := strconv.Atoi(id)
	if err != nil {
		return errors.New(fmt.Sprintf("Invalid param: %s", err.Error()))
	}
	db := database.GetConnector()

	result := db.Delete(&product.Product{}, productId)
	if result.Error != nil {
		return errors.New(fmt.Sprintf("Error deleting product: %s", result.Error.Error()))
	}

	if result.RowsAffected == 0 {
		return errors.New("No product found with the provided ID")
	}

	return nil

}
