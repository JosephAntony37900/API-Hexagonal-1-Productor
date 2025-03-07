package controllers

import (
	"strconv"

	"github.com/JosephAntony37900/API-Hexagonal-1-Productor/products/application"
	"github.com/gin-gonic/gin"
)

type UpdateProductController struct {
	updateProduct *application.UpdateProduct
}

func NewUpdateProductController(updateProduct *application.UpdateProduct) *UpdateProductController {
	return &UpdateProductController{updateProduct: updateProduct}
}

func (c *UpdateProductController) Handle(ctx *gin.Context) {
	// Obtener el ID del producto de la URL
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(400, gin.H{"error": "invalid product ID"})
		return
	}

	// Decodificar el cuerpo de la solicitud
	var request struct {
		Nombre  string  `json:"Nombre"`
		Precio float64 `json:"Precio"`
		Cantidad int `json:"Cantidad"`
	}
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(400, gin.H{"error": "invalid request body"})
		return
	}

	// Ejecutar el caso de uso para actualizar el producto
	if err := c.updateProduct.Run(id, request.Nombre, request.Precio, request.Cantidad); err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// Respuesta de éxito
	ctx.JSON(200, gin.H{"message": "product updated successfully"})
}
