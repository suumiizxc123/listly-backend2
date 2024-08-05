package main

import (
	"kcloudb1/cron"
	"kcloudb1/internal/config"
	"kcloudb1/internal/routes/admin_route"
	"kcloudb1/internal/routes/client_route"
	"kcloudb1/internal/routes/metal_route"
	"kcloudb1/internal/routes/order_route"
	"kcloudb1/internal/routes/upload_route"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.Header("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}

func main() {
	config.ConnectDatabase()
	config.RedisConfig()

	r := gin.Default()

	r.Use(Cors())

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Welcome to the main page!")
	})

	v1 := r.Group("/api/v1")

	metal_route.MetalRoute(v1)
	client_route.ClientRoute(v1)

	order_route.OrderRoute(v1)
	order_route.BalanceRoute(v1)
	order_route.PaymentRoute(v1)
	order_route.WithDrawRoute(v1)

	admin_route.AdminRoute(v1)
	// admin_route.FAQRoute(v1)
	admin_route.NewsRoute(v1)
	admin_route.SliderRoute(v1)
	admin_route.IngredientRoute(v1)
	admin_route.ProductRoute(v1)

	upload_route.UploadRoute(v1)

	go cron.CronJob()
	r.Run(":8080")
}
