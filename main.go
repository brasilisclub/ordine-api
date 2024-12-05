package main

import (
	"ordine-api/cmd"
	"ordine-api/pkg/database"
)

func main() {
	database.Connect()
	cmd.Execute()
	//cmd.Start()
}
