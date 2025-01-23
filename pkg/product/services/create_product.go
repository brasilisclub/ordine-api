package services

import (
	"errors"
	"fmt"
	"ordine-api/pkg/database"
	"ordine-api/pkg/product"
)

func CreateProduct(p *product.ProductRequestBody) (*product.Product, error) {
	db := database.GetConnector()

	var dbProduct product.Product

	dbProduct.Name = p.Name
	dbProduct.Category = p.Category
	dbProduct.Price = p.Price
	dbProduct.Description = p.Description
	dbProduct.Stock = p.Stock

	result := db.Create(&dbProduct)
	if result.Error != nil {
		return nil, errors.New(fmt.Sprintf("Error trying to insert product on database %s", result.Error.Error()))
	}

	return &dbProduct, nil
}
