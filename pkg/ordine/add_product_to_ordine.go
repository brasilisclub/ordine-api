package ordine

import (
	"errors"
	"fmt"
	"ordine-api/pkg/database"
	"ordine-api/pkg/product"

	"gorm.io/gorm"
)

func addProductToOrdine(ordine *Ordine, productId string) error {

	dbProduct, err := product.GetProductById(productId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("Product with id %s not founded", productId)
		}
		return fmt.Errorf("Error getting product on db: %w", err)
	}

	db := database.GetConnector()

	if err := db.Model(ordine).Association("Products").Append(&dbProduct); err != nil {
		return fmt.Errorf("Error trying to add product to ordine: %w", err)
	}

	return nil

}
