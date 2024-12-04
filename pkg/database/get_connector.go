package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetConnector() (*gorm.DB, error) {
	dsn := "root:root@tcp(database:3306)/ordine?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{}) // fazer singleton
	return db, err
}
