package migrations

import (
	"ordine-api/pkg/auth"
	"ordine-api/pkg/database"
	"ordine-api/pkg/ordine"
	"ordine-api/pkg/ordine/order"
	"ordine-api/pkg/product"

	"github.com/sirupsen/logrus"
)

func Migrate() {
	db := database.GetConnector()

	db.AutoMigrate(&ordine.Ordine{}, &product.Product{}, &order.OrderProducts{}, &auth.User{})
	logrus.Info("Migrations successfully executed!")
}
