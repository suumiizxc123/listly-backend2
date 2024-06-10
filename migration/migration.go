package main

import (
	"kcloudb1/internal/config"
	"kcloudb1/internal/models/client"
	"kcloudb1/internal/models/metal"
	"kcloudb1/internal/models/order"
)

func main() {

	config.ConnectDatabase()

	config.DB.AutoMigrate(

		client.Client{},
		metal.Metal{},
		metal.MetalRate{},
		order.Order{},
		order.OrderPayment{},
	)
}
