package application

import (
	"github.com/JosephAntony37900/API-Hexagonal-1-Productor/products/domain/entities"
	"github.com/JosephAntony37900/API-Hexagonal-1-Productor/products/domain/repository"
)

type GetProductsByMinPrice struct {
	productRepo repository.ProductRepository
}

func NewGetProductsByMinPrice(productRepo repository.ProductRepository) *GetProductsByMinPrice {
	return &GetProductsByMinPrice{productRepo: productRepo}
}

func (uc *GetProductsByMinPrice) Execute(minPrice float64) ([]entities.Product, error) {
	return uc.productRepo.FindByMinimumPrice(minPrice)
}
