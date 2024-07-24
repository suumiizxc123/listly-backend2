package client_route

import (
	"kcloudb1/internal/handlers/client_handler"
	"kcloudb1/internal/middleware"

	"github.com/gin-gonic/gin"
)

func ClientRoute(r *gin.RouterGroup) {
	clientRoute := r.Group("/client")
	{
		clientRoute.POST("/generate-otp", client_handler.GenerateOTP)
		clientRoute.POST("/verify-otp", client_handler.VerifyOTP)
		clientRoute.POST("/register", client_handler.Register)
		clientRoute.POST("/login-password", client_handler.LoginByPassword)
		clientRoute.POST("/check-token", client_handler.CheckToken)

		clientRoute.POST("/forgot-password", client_handler.ForgotPassword)
		clientRoute.POST("/verify-otp-change-password", client_handler.VerifyOTPChangePassword)
		clientRoute.POST("/change-password", client_handler.ChangePassword)

		clientRoute.GET("/profile", middleware.Auth(), client_handler.GetProfile)

	}
}
