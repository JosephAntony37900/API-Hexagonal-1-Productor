package routes

import (
	"github.com/JosephAntony37900/API-Hexagonal-1-Productor/products/infrastructure/controllers"
	"github.com/gin-gonic/gin"
)

func SetupProductRoutes(r *gin.Engine, createProductController *controllers.CreateProductController, getProductsController *controllers.GetProductsController, updateProductController *controllers.UpdateProductController, deleteProductController *controllers.DeleteProductController, getProductsByMinPriceController *controllers.GetProductsByMinPriceController,
	getProductByIDController *controllers.GetProductByIDController) {
	r.POST("/products", createProductController.Handle)
	r.GET("/products", getProductsController.Handle)
	r.GET("/products/:id", getProductByIDController.Handle)
	r.PUT("/products/:id", updateProductController.Handle)
	r.DELETE("/products/:id", deleteProductController.Handle)
	r.GET("/products/price/:minPrice", getProductsByMinPriceController.Handle)
}
