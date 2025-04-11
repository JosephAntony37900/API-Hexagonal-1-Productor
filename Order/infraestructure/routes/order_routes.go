package routes

import (
    "github.com/gin-gonic/gin"
    "github.com/JosephAntony37900/API-Hexagonal-1-Productor/Order/infraestructure/controllers"
)

func OrderRoutes(router *gin.Engine, createOrderController *controllers.CreateOrderController, getOrderController *controllers.GetOrderController, getOrderById *controllers.GetOrderByIDController) {
    router.POST("/orders", func(c *gin.Context) {
        createOrderController.Handle(c.Writer, c.Request)
    })

    router.GET("/orders/:usuario_id", getOrderController.Handle)
    router.GET("/orders/one/:id", getOrderById.Handle)
}

