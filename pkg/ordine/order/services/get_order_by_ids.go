package services

import (
	"ordine-api/pkg/database"
	"ordine-api/pkg/ordine/order"
)

func GetOrderByIds(ordineId uint, productId uint) (order.OrderProducts, error) {
	var dbOrderProduct order.OrderProducts

	c := database.GetConnector()
	err := c.Where("ordine_id = ? AND product_id = ?", ordineId, productId).First(&dbOrderProduct).Error

	return dbOrderProduct, err

}
