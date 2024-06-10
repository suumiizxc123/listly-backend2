package middleware

import (
	"fmt"
	"kcloudb1/internal/config"
	"kcloudb1/internal/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func CheckSecret() gin.HandlerFunc {
	return func(c *gin.Context) {

		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, utils.Error(
				[]string{"Forbidden", "Зөвшөөрөгдөөгүй нэвтрэлт"},
				"Forbidden",
			))
			c.Abort()
			return
		}

		// Assuming the token is in the format "Bearer <token>"
		token := strings.TrimPrefix(authHeader, "Bearer ")

		if token != "BblH6rsyEWlWOB6x2hkm6m1Ga3ITHCba" {
			c.JSON(http.StatusNotFound, utils.Error(
				[]string{"Unauthorized", "Хэрэглэгч зөвшөөрөгдөөгүй төхөөрөмжнөөс нэвтрэсэн байна"},
				"X-Secret",
			))
			c.Abort()
			return
		}
		c.Next()
	}
}

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {

		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, utils.Error(
				[]string{"Forbidden", "Зөвшөөрөгдөөгүй нэвтрэлт"},
				"Forbidden",
			))
			c.Abort()
			return
		}

		claims, err := VerifyToken(authHeader)
		if err != nil {
			c.JSON(http.StatusUnauthorized, utils.Error(
				[]string{"Unauthorized", "Нэвтрээгүй байна"},
				err.Error(),
			))
			c.Abort()
			return
		}

		// clientID := claims["clientID"].(int64)
		tokenString := claims["token"].(string)
		fmt.Sprintln("Token string:", tokenString)

		clientIDStr, _ := config.RS.Get(tokenString).Result()

		fmt.Println("Client ID:", clientIDStr)
		c.Set("clientID", clientIDStr)

		c.Next()
	}
}
