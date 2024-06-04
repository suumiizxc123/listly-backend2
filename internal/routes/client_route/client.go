package client_route

import (
	"kcloudb1/internal/handlers/client_handler"

	"github.com/gin-gonic/gin"
)

func ClientRoute(r *gin.RouterGroup) {
	clientRoute := r.Group("/client")
	{
		clientRoute.POST("/generate-otp", client_handler.GenerateOTP)
		clientRoute.POST("/verify-otp", client_handler.VerifyOTP)
		clientRoute.POST("/register", client_handler.Register)
		clientRoute.POST("/login-password", client_handler.LoginByPassword)
	}
}
