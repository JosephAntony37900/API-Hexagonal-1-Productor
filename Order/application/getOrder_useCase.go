package application

import (
	"github.com/JosephAntony37900/API-Hexagonal-1-Productor/Order/domain/entities"
	"github.com/JosephAntony37900/API-Hexagonal-1-Productor/Order/domain/repository"
)

type GetOrders struct {
	repo repository.OrderRepository
}

func NewGetOrders(repo repository.OrderRepository) *GetOrders{
	return &GetOrders{repo: repo}
}

func (geto *GetOrders) Run(usuarioID int) ([]entities.Order, error) {
	orders, err := geto.repo.FindByUserID(usuarioID)
	if err != nil {
		return nil, err
	}
	return orders, nil
}
