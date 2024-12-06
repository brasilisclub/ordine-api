package migrations

import (
	"ordine-api/pkg/auth"
	"ordine-api/pkg/database"
	"ordine-api/pkg/ordine"
	"ordine-api/pkg/product"

	"github.com/sirupsen/logrus"
)

func Migrate() {
	db := database.GetConnector()

	db.AutoMigrate(&ordine.Ordine{}, &product.Product{}, ordine.OrderProducts{}, &auth.User{})
	logrus.Info("Migrations successfully executed!")
}
