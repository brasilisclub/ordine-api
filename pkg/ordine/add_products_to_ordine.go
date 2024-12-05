package ordine

import (
	"errors"
	"fmt"
	"strconv"

	"gorm.io/gorm"
)

func addProductsToOrdine(ordineId string, ids []int) (*Ordine, error) {

	dbOrdine, err := getOrdineById(ordineId)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("Ordine with id %s not founded", ordineId)
		}
		return nil, fmt.Errorf("Error getting ordine on db: %w", err)

	}

	for _, id := range ids {

		productId := strconv.Itoa(id)
		err := addProductToOrdine(&dbOrdine, productId)

		if err != nil {
			return nil, err
		}
	}

	return &dbOrdine, nil

}
