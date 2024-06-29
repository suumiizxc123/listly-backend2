package main

import (
	"kcloudb1/internal/config"
	"kcloudb1/internal/models/admin"
	"kcloudb1/internal/models/client"
	"kcloudb1/internal/models/metal"
	"kcloudb1/internal/models/order"
	"kcloudb1/internal/models/payment"
)

func main() {

	config.ConnectDatabase()

	config.DB.AutoMigrate(

		client.Client{},
		metal.Metal{},
		metal.MetalRate{},
		order.Order{},
		order.OrderPayment{},
		order.Balance{},
		order.BalanceHistory{},
		payment.QPayToken{},
		admin.Admin{},
	)
}
