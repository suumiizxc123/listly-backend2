package main

import (
	"kcloudb1/internal/config"
	"kcloudb1/internal/models/client"
)

func main() {

	config.ConnectDatabase()

	config.DB.AutoMigrate(

		client.Client{},
	)
}
