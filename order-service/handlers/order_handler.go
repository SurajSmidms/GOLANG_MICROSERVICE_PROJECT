package handlers

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"smarket/order-service/database"
	"smarket/order-service/models"
	"smarket/order-service/utils"
)

func CreateOrder(c *gin.Context) {
	token := c.GetHeader("Authorization")
	username, err := utils.ValidateToken(token, os.Getenv("ACCESS_SECRET"))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid or expired token"})
		return
	}

	var order models.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	order.UserID = 1 // demo user; in future we’ll map this via username
	if err := database.DB.Create(&order).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create order"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "order created successfully",
		"user":    username,
		"order":   order,
	})
}

func GetOrders(c *gin.Context) {
	token := c.GetHeader("Authorization")
	username, err := utils.ValidateToken(token, os.Getenv("ACCESS_SECRET"))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid or expired token"})
		return
	}

	var orders []models.Order
	if err := database.DB.Find(&orders).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch orders"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user":   username,
		"orders": orders,
	})
}
