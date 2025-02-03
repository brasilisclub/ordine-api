package tests

import "ordine-api/pkg/database"

func DropTablesForTests(structs ...interface{}) {
	c := database.GetConnector()
	c.Migrator().DropTable(structs...)
}
