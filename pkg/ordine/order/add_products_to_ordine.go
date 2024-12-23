package order

import (
	"errors"
	"fmt"
	ord "ordine-api/pkg/ordine"
	"ordine-api/pkg/ordine/services"
	"strconv"

	"gorm.io/gorm"
)

func AddProductsToOrdine(ordineId string, items []ord.OrderProductBody) (*ord.Ordine, error) {

	dbOrdine, err := services.GetOrdineById(ordineId)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("ordine with id %s not founded", ordineId)
		}
		return nil, fmt.Errorf("error getting ordine on db: %w", err)

	}

	for _, item := range items {

		productId := strconv.FormatUint(uint64(item.ProductID), 10)
		err := addProductToOrdine(dbOrdine.ID, productId, item.Quantity)

		if err != nil {
			return nil, err
		}
	}

	newOrdine, _ := services.GetOrdineById(strconv.FormatUint(uint64(dbOrdine.ID), 10))

	return &newOrdine, nil

}
