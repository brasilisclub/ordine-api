package database

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var (
	connection *gorm.DB
)

func Connect() {

	dsn := getDsn()
	c, err := getConnection(dsn)
	if err != nil {
		logrus.Fatalf("error connecting to database %s", err.Error())
	}
	connection = c

}

func GetConnector() *gorm.DB {
	if connection == nil {
		Connect()
	}
	return connection
}
