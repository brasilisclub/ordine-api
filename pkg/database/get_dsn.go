package database

import (
	"fmt"
	"ordine-api/config"
)

func getDsn() string {
	if config.IsApplicationInTestMode() {
		return ":memory:"
	}

	user := config.Envs.DATABASE_USER
	password := config.Envs.DATABASE_PASSWORD
	host := config.Envs.DATABASE_HOST
	port := config.Envs.DATABASE_PORT

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/ordine?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port)
	return dsn
}
