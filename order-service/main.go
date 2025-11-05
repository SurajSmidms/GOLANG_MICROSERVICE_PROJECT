package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"smarket/order-service/database"
	"smarket/order-service/handlers"
)

func main() {
	database.ConnectDB()

	r := gin.Default()
	r.POST("/orders", handlers.CreateOrder)
	r.GET("/orders", handlers.GetOrders)

	log.Println("🚀 Order service running on port 8082")
	if err := r.Run(":8082"); err != nil {
		log.Fatal("❌ Failed to start order service:", err)
	}
}
