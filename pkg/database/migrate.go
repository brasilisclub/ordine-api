package database

import (
	"ordine-api/pkg/ordine"

	"github.com/sirupsen/logrus"
)

func Migrate() {
	db, err := GetConnector()
	if err != nil {
		logrus.Errorf("Error getting database connector %s", err.Error())
	}
	db.AutoMigrate(&ordine.Ordine{}, &ordine.Product{})
	logrus.Info("Migrations successfully executed!")
}
