package main

import (
    "log"

    helpers "github.com/JosephAntony37900/API-Hexagonal-1-Productor/helpers"

    // Productos
    productApp "github.com/JosephAntony37900/API-Hexagonal-1-Productor/products/application"
    productController "github.com/JosephAntony37900/API-Hexagonal-1-Productor/products/infrastructure/controllers"
    productRepo "github.com/JosephAntony37900/API-Hexagonal-1-Productor/products/infrastructure/repository"
    productRoutes "github.com/JosephAntony37900/API-Hexagonal-1-Productor/products/infrastructure/routes"

    // Usuarios 
    userApp "github.com/JosephAntony37900/API-Hexagonal-1-Productor/Users/application"
    userController "github.com/JosephAntony37900/API-Hexagonal-1-Productor/Users/infraestructure/controllers"
    userRepo "github.com/JosephAntony37900/API-Hexagonal-1-Productor/Users/infraestructure/repository"
    userRoutes "github.com/JosephAntony37900/API-Hexagonal-1-Productor/Users/infraestructure/routes"

    // Pedidos
    orderApp "github.com/JosephAntony37900/API-Hexagonal-1-Productor/Order/application"
    orderController "github.com/JosephAntony37900/API-Hexagonal-1-Productor/Order/infraestructure/controllers"
    orderRepo "github.com/JosephAntony37900/API-Hexagonal-1-Productor/Order/infraestructure/repository"
    orderRoutes "github.com/JosephAntony37900/API-Hexagonal-1-Productor/Order/infraestructure/routes"

    "github.com/gin-gonic/gin"
)

func main() {
    // Conexión a MySQL
    db, err := helpers.NewMySQLConnection()
    if err != nil {
        log.Fatalf("Error conectando a la BD: %v", err)
    }
    defer db.Close()

    // Repositorios
    productRepository := productRepo.NewProductRepoMySQL(db)
    userRepository := userRepo.NewCreateUserRepoMySQL(db)
    orderRepository := orderRepo.NewOrderRepoMySQL(db) 

    // Casos de uso para productos
    createProduct := productApp.NewCreateProduct(productRepository)
    getProducts := productApp.NewGetProducts(productRepository)
    updateProduct := productApp.NewUpdateProduct(productRepository)
    deleteProduct := productApp.NewDeleteProduct(productRepository)
    getProductsByMinPrice := productApp.NewGetProductsByMinPrice(productRepository)

    // Casos de uso para usuarios
    createUser := userApp.NewCreateUser(userRepository)
    getUsers := userApp.NewGetUsers(userRepository)
    deleteUsers := userApp.NewDeleteUser(userRepository)
    updateUsers := userApp.NewUpdateUser(userRepository)

    // Casos de uso para pedidos
    createOrder := orderApp.NewCreateOrder(orderRepository) 
    getOrders := orderApp.NewGetOrders(orderRepository)  

    // Controladores para productos
    createProductController := productController.NewCreateProductController(createProduct)
    getProductsController := productController.NewGetProductsController(getProducts)
    updateProductController := productController.NewUpdateProductController(updateProduct)
    deleteProductController := productController.NewDeleteProductController(deleteProduct)
    getProductsByMinPriceController := productController.NewGetProductsByMinPriceController(getProductsByMinPrice)

    // Controladores para usuarios
    createUserController := userController.NewCreateUserController(createUser)
    getUserController := userController.NewUsersController(getUsers)
    deleteUserController := userController.NewDeleteUserController(deleteUsers)
    updateUserController := userController.NewUpdateUserController(updateUsers)

    // Controladores para pedidos
    createOrderController := orderController.NewCreateOrderController(createOrder)
    getOrdersController := orderController.NewGetOrderController(getOrders)

    // Configuración del enrutador de Gin
    r := gin.Default()

    // Configurar CORS
    r.Use(helpers.SetupCORS())

    // Configurar rutas de productos
    productRoutes.SetupProductRoutes(r, createProductController, getProductsController, updateProductController, deleteProductController, getProductsByMinPriceController)

    // Configurar rutas de usuarios
    userRoutes.SetupUserRoutes(r, createUserController, getUserController, deleteUserController, updateUserController)

    // Configurar rutas de pedidos ✅
    orderRoutes.OrderRoutes(r, createOrderController, getOrdersController)

    // Iniciar servidor
    log.Println("Server started at :8080")
    if err := r.Run(":8080"); err != nil {
        log.Fatalf("Error starting server: %v", err)
    }
}
