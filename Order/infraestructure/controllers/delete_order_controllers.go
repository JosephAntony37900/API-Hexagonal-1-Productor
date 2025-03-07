package controllers

import (
    "net/http"
    "strconv"

    "github.com/JosephAntony37900/API-Hexagonal-1-Productor/Order/application"
    "github.com/gin-gonic/gin"
)

type DeleteOrderController struct {
    useCase *application.DeleteOrder
}

func NewDeleteOrderController(useCase *application.DeleteOrder) *DeleteOrderController {
    return &DeleteOrderController{useCase: useCase}
}

func (c *DeleteOrderController) Handle(ctx *gin.Context) {
    orderIDParam := ctx.Param("id")
    orderID, err := strconv.Atoi(orderIDParam)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID de pedido inválido"})
        return
    }

    err = c.useCase.Run(orderID)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, gin.H{"message": "Pedido eliminado con éxito"})
}
