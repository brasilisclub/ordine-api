package migrations

import (
	"ordine-api/pkg/database"
	"ordine-api/pkg/ordine"
	"ordine-api/pkg/product"

	"github.com/sirupsen/logrus"
)

func Migrate() {
	db, err := database.GetConnector()
	if err != nil {
		logrus.Errorf("Error getting database connector %s", err.Error())
	}
	db.AutoMigrate(&ordine.Ordine{}, &product.Product{})
	logrus.Info("Migrations successfully executed!")
}
