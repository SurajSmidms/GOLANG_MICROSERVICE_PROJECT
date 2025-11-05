package main

import (
	"log"
	"smarket/payment-service/database"
	"smarket/payment-service/models"
	"smarket/payment-service/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	database.ConnectDB()
	database.DB.AutoMigrate(&models.Payment{})

	r := gin.Default()
	routes.RegisterRoutes(r)

	log.Println("🚀 Payment service running on port 8084")
	if err := r.Run(":8084"); err != nil {
		log.Fatal("❌ Failed to start payment service:", err)
	}
}
