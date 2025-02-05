package database

import (
	"ordine-api/config"

	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func getConnection(dsn string) (db *gorm.DB, err error) {
	if config.IsApplicationInTestMode() {
		return gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	}

	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}
