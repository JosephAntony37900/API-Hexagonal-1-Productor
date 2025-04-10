package main

import (
    "log"
    "os"
    "os/signal"
    "syscall"

    helpers "github.com/JosephAntony37900/API-Hexagonal-1-Productor/helpers"
    init_product "github.com/JosephAntony37900/API-Hexagonal-1-Productor/products/infrastructure"
    init_users "github.com/JosephAntony37900/API-Hexagonal-1-Productor/Users/infraestructure"
    init_order "github.com/JosephAntony37900/API-Hexagonal-1-Productor/Order/infraestructure"
    "github.com/gin-gonic/gin"
    "github.com/JosephAntony37900/API-Hexagonal-1-Productor/Order/infraestructure/adapters"
)

func main() {
    db, err := helpers.NewMySQLConnection()
    if err != nil {
        log.Fatalf("Error conectando a la BD: %v", err)
    }
    defer db.Close()

    r := gin.Default()

    // CORS
    r.Use(helpers.SetupCORS())

    init_product.InitProductDependencies(r, db)
    init_order.InitOrderDependencies(r, db)
    init_users.InitUsersDependencies(r, db)

    // Manejar señales para cerrar correctamente el rabbit
    sigChan := make(chan os.Signal, 1)
    signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
    go func() {
        <-sigChan
        log.Println("Cerrando conexiones...")
        adapters.CloseRabbitMQ()
        os.Exit(0)
    }()

    log.Println("Server escuchando en :8080")
    if err := r.Run(":8080"); err != nil {
        log.Fatalf("Error iniciando el server: %v", err)
    }
}