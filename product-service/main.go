package main

import (
	"log"
	"smarket/product-service/database"
	"smarket/product-service/models"
	"smarket/product-service/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	database.ConnectDB()
	database.DB.AutoMigrate(&models.Product{})

	r := gin.Default()
	routes.RegisterRoutes(r)

	log.Println("🚀 Product service running on port 8083")
	if err := r.Run(":8083"); err != nil {
		log.Fatal("❌ Failed to start product service:", err)
	}
}
