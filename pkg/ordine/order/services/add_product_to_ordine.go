package services

import (
	"fmt"
	"ordine-api/pkg/product"
	"ordine-api/pkg/product/services"
)

func addProductToOrdine(ordId uint, prodId uint, quant int) error {
	if quant <= 0 {
		return fmt.Errorf("quantity <= 0")
	}

	if !services.ProductExists(prodId) {
		return product.ErrorProductNotFound
	}

	if !OrderExists(ordId, prodId) {
		return CreateNewOrder(ordId, prodId, quant)
	}

	err := InsertProductToOrder(ordId, prodId, quant)

	return err
}
