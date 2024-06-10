package main

import (
	"kcloudb1/internal/config"
	"kcloudb1/internal/routes/client_route"
	"kcloudb1/internal/routes/metal_route"
	"kcloudb1/internal/routes/order_route"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")

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

	r.Run(":8080")
}
