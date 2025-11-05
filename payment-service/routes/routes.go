package routes

import (
	"smarket/payment-service/handlers"
	"smarket/payment-service/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	payments := r.Group("/payments")
	payments.Use(middleware.AuthMiddleware()) // ✅ Protect all payment routes
	{
		payments.POST("/", handlers.CreatePayment)
		payments.GET("/", handlers.GetPayments)
		payments.GET("/:id", handlers.GetPaymentByID)
	}
}