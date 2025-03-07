package controllers

import (
    "encoding/json"
    "net/http"
    "strconv"

    "github.com/JosephAntony37900/API-Hexagonal-1-Productor/Order/application"
    "github.com/gin-gonic/gin"
)

type UpdateOrderRequest struct {
    Producto          string `json:"producto"`
    Estado            string `json:"estado"`
    Pais              string `json:"pais"`
    Entidad_federativa string `json:"entidad_federativa"`
    Cp                string `json:"cp"`
}

type UpdateOrderController struct {
    useCase *application.UpdateOrder
}

func NewUpdateOrderController(useCase *application.UpdateOrder) *UpdateOrderController {
    return &UpdateOrderController{useCase: useCase}
}

func (c *UpdateOrderController) Handle(ctx *gin.Context) {
    orderIDParam := ctx.Param("id")
    orderID, err := strconv.Atoi(orderIDParam)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID de pedido inválido"})
        return
    }

    var req UpdateOrderRequest
    if err := json.NewDecoder(ctx.Request.Body).Decode(&req); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Error en el formato de la solicitud"})
        return
    }

    err = c.useCase.Run(orderID, req.Producto, req.Estado, req.Pais, req.Entidad_federativa, req.Cp)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, gin.H{"message": "Estado del pedido actualizado con éxito"})
}
