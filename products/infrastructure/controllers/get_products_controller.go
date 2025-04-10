package controllers

import (
	"log"

	"github.com/JosephAntony37900/API-Hexagonal-1-Productor/products/application"
	"github.com/gin-gonic/gin"
)

type GetProductsController struct {
	getProducts *application.GetProducts
}

func NewGetProductsController(getProducts *application.GetProducts) *GetProductsController {
	return &GetProductsController{getProducts: getProducts}
}

func (c *GetProductsController) Handle(ctx *gin.Context) {
	log.Println("Petición para listar todos los productos, recibido")
	products, err := c.getProducts.Run()
	if err != nil {
		log.Printf("Error buscando products: %v", err)
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	log.Printf("Retornando %d products", len(products))
	ctx.JSON(200, products)
}
