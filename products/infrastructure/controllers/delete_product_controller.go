package controllers

import (
	"strconv"

	"github.com/JosephAntony37900/API-Hexagonal-1-Productor/products/application"
	"github.com/gin-gonic/gin"
)

type DeleteProductController struct {
	deleteProduct *application.DeleteProduct
}

func NewDeleteProductController(deleteProduct *application.DeleteProduct) *DeleteProductController {
	return &DeleteProductController{deleteProduct: deleteProduct}
}

func (c *DeleteProductController) Handle(ctx *gin.Context) {
	// Obtener el ID del producto de la URL
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(400, gin.H{"error": "invalid product ID"})
		return
	}

	// Ejecutar el caso de uso para eliminar el producto
	if err := c.deleteProduct.Run(id); err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// Respuesta de éxito
	ctx.JSON(200, gin.H{"message": "product deleted successfully"})
}
