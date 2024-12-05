package database

import (
	"sync"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	instance *gorm.DB
	once     sync.Once
	err      error
)

func GetConnector() (*gorm.DB, error) {
	once.Do(func() {
		dsn := "root:root@tcp(database:3306)/ordine?charset=utf8mb4&parseTime=True&loc=Local"
		instance, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			return
		}
		logrus.Info("Database connection successfully established!")
	})

	return instance, err
}
