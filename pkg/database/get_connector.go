package database

import (
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	connection *gorm.DB
)

func Connect() {

	dsn := getDsn()
	c, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logrus.Fatal("error connecting to database")
	}
	connection = c

}

func GetConnector() *gorm.DB {
	if connection == nil {
		Connect()
	}
	return connection
}
