package handlers

import (
	"net/http"
	"smarket/payment-service/database"
	"smarket/payment-service/models"

	"github.com/gin-gonic/gin"
)

func CreatePayment(c *gin.Context) {
	var payment models.Payment
	if err := c.ShouldBindJSON(&payment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	payment.Status = "SUCCESS"

	if err := database.DB.Create(&payment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create payment"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "payment created", "payment": payment})
}

func GetPayments(c *gin.Context) {
	var payments []models.Payment
	if err := database.DB.Find(&payments).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch payments"})
		return
	}
	c.JSON(http.StatusOK, payments)
}

func GetPaymentByID(c *gin.Context) {
	id := c.Param("id")
	var payment models.Payment
	if err := database.DB.First(&payment, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "payment not found"})
		return
	}
	c.JSON(http.StatusOK, payment)
}
