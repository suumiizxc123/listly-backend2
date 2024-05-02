package main

import (
	"io"
	"kcloudb1/internal/config"
	"kcloudb1/internal/routes"
	"log"
	"net/http"
	"os"

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
	f, err := os.Create("gin.log")
	if err != nil {
		log.Fatal(err)
	}

	r.Use(Cors())
	gin.DefaultWriter = io.MultiWriter(f)

	v1 := r.Group("/api/v1")

	routes.SysUserRoute(v1)
	routes.UserRoute(v1)

	r.Run()
}
