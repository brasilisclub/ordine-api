package ordine

import (
	"ordine-api/pkg/database"

	"github.com/sirupsen/logrus"
)

func Migrate() {
	db, err := database.GetConnector()
	if err != nil {
		logrus.Errorf("Error getting database connector %s", err.Error())
	}
	db.AutoMigrate(&Ordine{}, &Product{})
	logrus.Info("Migrations successfully executed!")
}
