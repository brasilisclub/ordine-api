package ordine

import (
	"errors"
	"fmt"
	"strconv"

	"gorm.io/gorm"
)

func addProductsToOrdine(ordineId string, items []OrderProductBody) (*Ordine, error) {

	dbOrdine, err := GetOrdineById(ordineId)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("Ordine with id %s not founded", ordineId)
		}
		return nil, fmt.Errorf("Error getting ordine on db: %w", err)

	}

	for _, item := range items {

		productId := strconv.FormatUint(uint64(item.ProductID), 10)
		err := addProductToOrdine(dbOrdine.ID, productId, item.Quantity)

		if err != nil {
			return nil, err
		}
	}

	return &dbOrdine, nil

}
