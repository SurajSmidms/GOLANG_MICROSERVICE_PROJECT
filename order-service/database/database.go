package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"smarket/order-service/models"
)

var DB *gorm.DB

func ConnectDB() {
	// Load env from project root (not service folder)
	err := godotenv.Load("../.env")
	if err != nil {
		log.Printf("⚠️ Could not load .env file from project root: %v", err)
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("❌ Failed to connect to MySQL (Order Service): %v", err)
	}

	DB = db

	err = db.AutoMigrate(&models.Order{})
	if err != nil {
		log.Fatalf("❌ Failed to migrate order model: %v", err)
	}

	log.Println("✅ Connected to MySQL (Order Service)")
}
