package order

import (
	"errors"
	"fmt"
	ord "ordine-api/pkg/ordine"
	"strconv"

	"gorm.io/gorm"
)

func AddProductsToOrdine(ordineId string, items []ord.OrderProductBody) (*ord.Ordine, error) {

	dbOrdine, err := ord.GetOrdineById(ordineId)

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
