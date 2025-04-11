package controllers

import (
	"log"
	"strconv"

	"github.com/JosephAntony37900/API-Hexagonal-1-Productor/Order/application"
	"github.com/gin-gonic/gin"
)

type GetOrderByIDController struct {
	getOrderByID *application.GetByIdOrder
}

func NewGetOrderByIDController(getOrderByID *application.GetByIdOrder) *GetOrderByIDController {
	return &GetOrderByIDController{getOrderByID: getOrderByID}
}

func (c *GetOrderByIDController) Handle(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Printf("Error convirtiendo ID a número: %v", err)
		ctx.JSON(400, gin.H{"error": "ID inválido"})
		return
	}

	log.Printf("Petición para obtener Ordero con ID %d, recibido", id)
	Order, err := c.getOrderByID.Run(id)
	if err != nil {
		log.Printf("Error buscando Ordero: %v", err)
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	if Order == nil {
		log.Printf("Ordero con ID %d no encontrado", id)
		ctx.JSON(404, gin.H{"error": "Ordero no encontrado"})
		return
	}

	log.Printf("Retornando Ordero con ID %d", id)
	ctx.JSON(200, Order)
}