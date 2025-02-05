package tests

import "ordine-api/pkg/database"

func CreateInsertValueForTests(value interface{}) {
	c := database.GetConnector()
	c.Create(value)
}
