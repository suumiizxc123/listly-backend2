package main

import (
	"kcloudb1/internal/config"
	"kcloudb1/internal/models/client"
	"kcloudb1/internal/models/metal"
)

func main() {

	config.ConnectDatabase()

	config.DB.AutoMigrate(

		client.Client{},
		metal.Metal{},
		metal.MetalRate{},
	)
}
