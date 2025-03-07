package application

import (
	"fmt"
	"github.com/JosephAntony37900/API-Hexagonal-1-Productor/Order/domain/repository"
)

type DeleteOrder struct {
	repo repository.OrderRepository
}

func NewDeleteOrder(repo repository.OrderRepository) *DeleteOrder{
	return &DeleteOrder{repo: repo}
}

func (do *DeleteOrder) Run(id int) error {
	_, err := do.repo.FindByID(id)
	if err != nil {
		return fmt.Errorf("Pedido no encontrado: %w", err)
	}
	if err := do.repo.Delete(id); err != nil {
		return fmt.Errorf("Error eliminando el pedido: %w", err)
	}
	return nil
}