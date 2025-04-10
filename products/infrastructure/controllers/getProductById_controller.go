package controllers

import (
	"log"
	"strconv"

	"github.com/JosephAntony37900/API-Hexagonal-1-Productor/products/application"
	"github.com/gin-gonic/gin"
)

type GetProductByIDController struct {
	getProductByID *application.GetProductByID
}

func NewGetProductByIDController(getProductByID *application.GetProductByID) *GetProductByIDController {
	return &GetProductByIDController{getProductByID: getProductByID}
}

func (c *GetProductByIDController) Handle(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Printf("Error convirtiendo ID a número: %v", err)
		ctx.JSON(400, gin.H{"error": "ID inválido"})
		return
	}

	log.Printf("Petición para obtener producto con ID %d, recibido", id)
	product, err := c.getProductByID.Run(id)
	if err != nil {
		log.Printf("Error buscando producto: %v", err)
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	if product == nil {
		log.Printf("Producto con ID %d no encontrado", id)
		ctx.JSON(404, gin.H{"error": "producto no encontrado"})
		return
	}

	log.Printf("Retornando producto con ID %d", id)
	ctx.JSON(200, product)
}