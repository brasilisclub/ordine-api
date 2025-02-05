package services

import (
	"ordine-api/pkg/database"
)

func InsertProductToOrder(ordineId uint, productId uint, quant int) error {
	order, err := GetOrderByIds(ordineId, productId)
	if err != nil {
		return err
	}

	c := database.GetConnector()

	order.Quantity = order.Quantity + quant

	err = c.Save(&order).Error

	return err
}
