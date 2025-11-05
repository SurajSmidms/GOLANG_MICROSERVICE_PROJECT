package routes

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"smarket/api-gateway/middleware"
	"smarket/api-gateway/utils"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	// load service base urls from env
	authBase := strings.TrimRight(os.Getenv("AUTH_SERVICE_URL"), "/")
	orderBase := strings.TrimRight(os.Getenv("ORDER_SERVICE_URL"), "/")
	productBase := strings.TrimRight(os.Getenv("PRODUCT_SERVICE_URL"), "/")
	paymentBase := strings.TrimRight(os.Getenv("PAYMENT_SERVICE_URL"), "/")

	// Health
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message":  "API Gateway running",
			"services": gin.H{"auth": authBase, "orders": orderBase, "products": productBase, "payments": paymentBase},
		})
	})

	// AUTH - public routes (register, login, refresh)
	r.Any("/auth/*path", func(c *gin.Context) {
		// build target using path (strip leading slash from path param)
		path := c.Param("path")
		if path == "" || path == "/" {
			// if path empty, forward to base / (optional)
			utils.ProxyRequest(c, authBase+"/")
			return
		}
		target := fmt.Sprintf("%s%s", authBase, path)
		utils.ProxyRequest(c, target)
	})

	// Protected groups: orders, products, payments
	// Use AuthMiddleware to validate token via Auth service
	protected := r.Group("/")
	protected.Use(middleware.AuthMiddleware())

	// Orders
	protected.Any("/orders/*path", func(c *gin.Context) {
		path := c.Param("path")
		if path == "" || path == "/" {
			utils.ProxyRequest(c, orderBase+"/")
			return
		}
		utils.ProxyRequest(c, fmt.Sprintf("%s%s", orderBase, path))
	})

	// Products
	protected.Any("/products/*path", func(c *gin.Context) {
		path := c.Param("path")
		if path == "" || path == "/" {
			utils.ProxyRequest(c, productBase+"/")
			return
		}
		utils.ProxyRequest(c, fmt.Sprintf("%s%s", productBase, path))
	})

	// Payments
	protected.Any("/payments/*path", func(c *gin.Context) {
		path := c.Param("path")
		if path == "" || path == "/" {
			utils.ProxyRequest(c, paymentBase+"/")
			return
		}
		utils.ProxyRequest(c, fmt.Sprintf("%s%s", paymentBase, path))
	})

	// Fallback - not found
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"error": "route not found on gateway"})
	})
}
