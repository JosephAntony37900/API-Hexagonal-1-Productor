package infraestructure

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	orderApp "github.com/JosephAntony37900/API-Hexagonal-1-Productor/Order/application"
    orderController "github.com/JosephAntony37900/API-Hexagonal-1-Productor/Order/infraestructure/controllers"
    orderRepo "github.com/JosephAntony37900/API-Hexagonal-1-Productor/Order/infraestructure/repository"
    orderRoutes "github.com/JosephAntony37900/API-Hexagonal-1-Productor/Order/infraestructure/routes"
	"github.com/JosephAntony37900/API-Hexagonal-1-Productor/Order/infraestructure/adapters"
)

func InitOrderDependencies(Engine *gin.Engine, db *sql.DB){
	adapters.InitRabbitMQ()
	defer adapters.CloseRabbitMQ()

	channel := adapters.GetChannel()


	orderRepository := orderRepo.NewOrderRepoMySQL(db, channel)     

	createOrder := orderApp.NewCreateOrder(orderRepository) 
    getOrders := orderApp.NewGetOrders(orderRepository) 

	createOrderController := orderController.NewCreateOrderController(createOrder)
    getOrdersController := orderController.NewGetOrderController(getOrders)

	go adapters.ConsumeConfirmedOrders(orderRepository)
	go adapters.ConsumeRejectedOrders(orderRepository)

	orderRoutes.OrderRoutes(Engine, createOrderController, getOrdersController)


}