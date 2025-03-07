package controllers

import (
	"log"

	"github.com/JosephAntony37900/API-Hexagonal-1-Productor/products/application"
	"github.com/gin-gonic/gin"
)

type CreateProductController struct {
	createProduct *application.CreateProduct
}

func NewCreateProductController(createProduct *application.CreateProduct) *CreateProductController {
	return &CreateProductController{createProduct: createProduct}
}

func (c *CreateProductController) Handle(ctx *gin.Context) {
	log.Println("Recibe la petición de la creación del producto")

	// Estructura para decodificar el JSON de la solicitud
	var request struct {
		Nombre  string  `json:"Nombre"`
		Precio float64 `json:"Precio"`
		Cantidad int `json:"Cantidad"`
	}

	// Decodificar el cuerpo de la solicitud
	if err := ctx.ShouldBindJSON(&request); err != nil {
		log.Printf("Error decoding request body: %v", err)
		ctx.JSON(400, gin.H{"error": "invalid request body"})
		return
	}
	log.Printf("Creating product: Nombre=%s, Precio=%f, Cantidad=%f", request.Nombre, request.Precio, request.Cantidad)

	// Ejecutar el caso de uso para crear el producto
	if err := c.createProduct.Run(request.Nombre, request.Precio, request.Cantidad); err != nil {
		log.Printf("Error creando el producto: %v", err)
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// Respuesta de éxito
	log.Println("Producto creado exitosamente")
	ctx.JSON(201, gin.H{"message": "producto creado exitosamente"})
}
