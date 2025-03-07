package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/JosephAntony37900/API-Hexagonal-1-Productor/Order/application"
)

type GetOrderController struct {
	getOrders *application.GetOrders
}

func NewGetOrderController(getOrders *application.GetOrders) *GetOrderController {
	return &GetOrderController{getOrders: getOrders}
}

func (ctrl *GetOrderController) Handle(c *gin.Context) {
	usuarioID, err := strconv.Atoi(c.Param("usuario_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de usuario inválido"})
		return
	}

	orders, err := ctrl.getOrders.Run(usuarioID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, orders)
}
