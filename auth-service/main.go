package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"smarket/auth-service/database"
	"smarket/auth-service/handlers"
	"smarket/auth-service/models"
)

func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("❌ Error loading .env file")
	}

	database.Connect()
	database.DB.AutoMigrate(&models.User{})

	r := gin.Default()

	r.POST("/register", handlers.Register)
	r.POST("/login", handlers.Login)
	r.GET("/validate", handlers.Validate)
	r.POST("/refresh", handlers.Refresh)

	r.Run(":8081")
}
