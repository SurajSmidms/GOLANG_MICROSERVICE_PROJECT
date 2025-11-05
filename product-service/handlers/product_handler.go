package handlers

import (
	"net/http"
	"smarket/product-service/database"
	"smarket/product-service/models"

	"github.com/gin-gonic/gin"
)

// Create a new product
func CreateProduct(c *gin.Context) {
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if result := database.DB.Create(&product); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Product created successfully", "product": product})
}

// Get all products
func GetProducts(c *gin.Context) {
	var products []models.Product
	database.DB.Find(&products)
	c.JSON(http.StatusOK, products)
}

// Get single product by ID
func GetProductByID(c *gin.Context) {
	id := c.Param("id")
	var product models.Product
	if result := database.DB.First(&product, id); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}
	c.JSON(http.StatusOK, product)
}

// Update product
func UpdateProduct(c *gin.Context) {
	id := c.Param("id")
	var product models.Product
	if result := database.DB.First(&product, id); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	var update models.Product
	if err := c.ShouldBindJSON(&update); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Model(&product).Updates(update)
	c.JSON(http.StatusOK, gin.H{"message": "Product updated", "product": product})
}

// Delete product
func DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	if result := database.DB.Delete(&models.Product{}, id); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Product deleted"})
}
