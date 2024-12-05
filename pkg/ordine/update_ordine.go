package ordine

import (
	"errors"
	"fmt"
	"ordine-api/pkg/database"

	"gorm.io/gorm"
)

func updateOrdine(id string, o *Ordine) (*Ordine, error) {
	db, err := database.GetConnector()
	if err != nil {
		return nil, fmt.Errorf("Error trying to connect to database: %w", err)
	}

	var dbOrdine Ordine
	if err := db.First(&dbOrdine, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("Ordine with id %s not founded", id)
		}
		return nil, fmt.Errorf("Error getting ordine on db: %w", err)
	}
	dbOrdine.Table = o.Table
	dbOrdine.ClientName = o.ClientName
	dbOrdine.Status = o.Status

	if err := db.Save(&dbOrdine).Error; err != nil {
		return nil, fmt.Errorf("Error trying update product: %w", err)
	}

	return &dbOrdine, nil
}
