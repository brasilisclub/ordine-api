package services

import (
	"errors"
	"fmt"
	"ordine-api/pkg/database"
	"ordine-api/pkg/ordine"

	"gorm.io/gorm"
)

func UpdateOrdine(id string, o *ordine.OrdineRequestBody) (*ordine.Ordine, error) {
	db := database.GetConnector()

	dbOrdine, err := GetOrdineById(id)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ordine.ErrorOrdineNotFound
		}
		return nil, fmt.Errorf("error getting ordine on db: %w", err)

	}

	dbOrdine.Table = o.Table
	dbOrdine.ClientName = o.ClientName
	dbOrdine.Status = o.Status

	if err := db.Save(&dbOrdine).Error; err != nil {
		return nil, fmt.Errorf("error trying update product: %w", err)
	}

	return &dbOrdine, nil
}
