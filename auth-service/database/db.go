package database

import (
    "log"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "smarket/auth-service/models"
)

var DB *gorm.DB

func Connect() {
    dsn := "root:root@tcp(localhost:3306)/smarket?charset=utf8mb4&parseTime=True&loc=Local"
    var err error
    DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("❌ Failed to connect to database:", err)
    }

    log.Println("✅ Connected to MySQL")

    // Auto-migrate all models
    err = DB.AutoMigrate(&models.User{})
    if err != nil {
        log.Fatal("❌ AutoMigrate failed:", err)
    }
}
