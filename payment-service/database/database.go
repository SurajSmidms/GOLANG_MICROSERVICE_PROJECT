package database

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"smarket/payment-service/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	// ✅ Load .env from project root (two levels up)
	rootPath, _ := filepath.Abs("../")
	envPath := filepath.Join(rootPath, ".env")

	if err := godotenv.Load(envPath); err != nil {
		log.Println("⚠️  No .env file found at project root, using system environment variables")
	}

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	if dbUser == "" || dbHost == "" || dbPort == "" || dbName == "" {
		log.Fatal("❌ Missing one or more DB environment variables")
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		dbUser, dbPass, dbHost, dbPort, dbName)

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("❌ Failed to connect to database (Payment Service): %v", err)
	}

	err = DB.AutoMigrate(&models.Payment{})
	if err != nil {
		log.Fatalf("❌ Failed to migrate Payment model: %v", err)
	}

	log.Println("✅ Connected to MySQL (Payment Service)")
}
