package ordine

import (
	"errors"
	"fmt"
	"ordine-api/pkg/database"
	"ordine-api/pkg/product"

	"gorm.io/gorm"
)

func addProductToOrdine(ordId uint, prodId string, quant int) error {
	if quant <= 0 {
		return fmt.Errorf("Quantity <= 0")
	}

	dbProduct, err := product.GetProductById(prodId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("Product with ID %s not founded", prodId)
		}
		return fmt.Errorf("Error getting product on database: %w", err)
	}

	db := database.GetConnector()

	var dbOrderProduct OrderProducts
	err = db.Where("ordine_id = ? AND product_id = ?", ordId, dbProduct.ID).First(&dbOrderProduct).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		newOrderProduct := OrderProducts{
			OrdineID:  ordId,
			ProductID: dbProduct.ID,
			Quantity:  quant,
		}
		if err := db.Create(&newOrderProduct).Error; err != nil {
			return fmt.Errorf("Error trying to create order: %w", err)
		}
	} else if err != nil {

		return fmt.Errorf("Error getting order %w", err)
	} else {

		dbOrderProduct.Quantity += quant
		if err := db.Save(&dbOrderProduct).Error; err != nil {
			return fmt.Errorf("Error trying to update order %w", err)
		}
	}

	return nil
}
