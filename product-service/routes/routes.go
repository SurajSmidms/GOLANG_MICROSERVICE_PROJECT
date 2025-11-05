package routes

import (
	"github.com/gin-gonic/gin"
	"smarket/product-service/handlers"
)

func RegisterRoutes(r *gin.Engine) {
	r.POST("/products", handlers.CreateProduct)
	r.GET("/products", handlers.GetProducts)
	r.GET("/products/:id", handlers.GetProductByID)
	r.PUT("/products/:id", handlers.UpdateProduct)
	r.DELETE("/products/:id", handlers.DeleteProduct)
}
