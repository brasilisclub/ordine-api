package database

import (
	"os"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	connection *gorm.DB
)

func Connect() {
	dsn := "root:root@tcp(database:3306)/ordine?charset=utf8mb4&parseTime=True&loc=Local"
	c, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logrus.Fatal("error connecting to database")
		os.Exit(1)
	}
	connection = c
}

func TestConnect() {
	if connection == nil {
		c, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
		if err != nil {
			logrus.Fatal("error connecting to database")
			os.Exit(1)
		}
		connection = c
	}

}

func GetConnector() *gorm.DB {
	return connection
}
