package controllers

import (
	"net/http"
	"strconv"

	"github.com/JosephAntony37900/API-Hexagonal-1-Productor/products/application"
	"github.com/gin-gonic/gin"
)

type GetProductsByMinPriceController struct {
	getProductsByMinPriceUseCase *application.GetProductsByMinPrice
}

func NewGetProductsByMinPriceController(getProductsByMinPriceUseCase *application.GetProductsByMinPrice) *GetProductsByMinPriceController {
	return &GetProductsByMinPriceController{getProductsByMinPriceUseCase: getProductsByMinPriceUseCase}
}

func (ctrl *GetProductsByMinPriceController) Handle(c *gin.Context) {
	minPriceStr := c.Param("minPrice")
	minPrice, err := strconv.ParseFloat(minPriceStr, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "precio minimo invalido we"})
		return
	}

	products, err := ctrl.getProductsByMinPriceUseCase.Execute(minPrice)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, products)
}
