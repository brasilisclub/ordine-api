package migrations

import (
	"ordine-api/pkg/database"
	"ordine-api/pkg/ordine"
	"ordine-api/pkg/product"

	"github.com/sirupsen/logrus"
)

func Migrate() {
	db := database.GetConnector()

	db.AutoMigrate(&ordine.Ordine{}, &product.Product{})
	logrus.Info("Migrations successfully executed!")
}
