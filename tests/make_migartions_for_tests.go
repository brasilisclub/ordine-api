package tests

import "ordine-api/pkg/database"

func MakeMigrationsForTests(structs ...interface{}) {
	c := database.GetConnector()
	c.AutoMigrate(structs...)
}
