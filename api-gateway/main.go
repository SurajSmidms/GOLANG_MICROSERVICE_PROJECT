package main

import (
	"log"
	"os"

	"smarket/api-gateway/routes"

	"github.com/joho/godotenv"
	"github.com/gin-gonic/gin"
)

func main() {
	// try to load project level .env (one level up)
	_ = godotenv.Load("../.env")

	// optional: log loaded config (minimal)
	log.Println("🔁 Starting API Gateway...")

	r := gin.Default()
	routes.RegisterRoutes(r)

	port := os.Getenv("GATEWAY_PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("🚀 API Gateway running on port %s\n", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatal("❌ Failed to start API Gateway:", err)
	}
}
