package application

import (
	"github.com/JosephAntony37900/API-Hexagonal-1-Productor/Order/domain/repository"
	"github.com/JosephAntony37900/API-Hexagonal-1-Productor/Order/domain/entities"
)

type GetByIdOrder struct {
	repo repository.OrderRepository
}

func NewGetByIdOrder(repo repository.OrderRepository) *GetByIdOrder{
	return &GetByIdOrder{repo: repo}
}

func (do *GetByIdOrder) Run(id int) (*entities.Order, error) {
	orders, err := do.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	return orders, nil
}