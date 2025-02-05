package services

import (
	"errors"

	"ordine-api/pkg/ordine"
	"ordine-api/pkg/ordine/order"
	"ordine-api/pkg/ordine/services"
	"ordine-api/pkg/product"
	"ordine-api/pkg/utils"

	"gorm.io/gorm"
)

func AddProductsToOrdine(ordineId string, items []order.OrderProductBody) (*ordine.Ordine, error) {

	var ord ordine.Ordine

	if err := services.OrdineExists(ordineId); err != nil {
		if errors.Is(gorm.ErrRecordNotFound, err) {
			return &ord, ordine.ErrorOrdineNotFound
		}
		return &ord, err
	}

	for _, item := range items {

		convertedOrdId, _ := utils.StringToUint(ordineId)

		err := addProductToOrdine(convertedOrdId, uint(item.ProductID), item.Quantity)

		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return &ord, product.ErrorProductNotFound
			}
			return &ord, err
		}
	}

	ord, _ = services.GetOrdineById(ordineId)

	return &ord, nil

}
