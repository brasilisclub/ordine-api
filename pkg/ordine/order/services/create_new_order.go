package services

import (
	"ordine-api/pkg/database"
	"ordine-api/pkg/ordine/order"
)

func CreateNewOrder(ordineId uint, productId uint, quant int) error {

	c := database.GetConnector()

	newOrderProduct := order.OrderProducts{
		OrdineID:  ordineId,
		ProductID: productId,
		Quantity:  quant,
	}
	err := c.Create(&newOrderProduct).Error

	return err

}
