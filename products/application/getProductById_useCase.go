package application

import (
	"github.com/JosephAntony37900/API-Hexagonal-1-Productor/products/domain/entities"
	"github.com/JosephAntony37900/API-Hexagonal-1-Productor/products/domain/repository"
)

type GetProductByID struct {
	repo repository.ProductRepository
}

func NewGetProductByID(repo repository.ProductRepository) *GetProductByID {
	return &GetProductByID{repo: repo}
}

func (gp *GetProductByID) Run(id int) (*entities.Product, error) {
	product, err := gp.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	return product, nil
}