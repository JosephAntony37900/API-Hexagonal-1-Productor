package application

import (
	"fmt"

	"github.com/JosephAntony37900/API-Hexagonal-1-Productor/Order/domain/repository"
)

type UpdateOrder struct {
	repo repository.OrderRepository
}

func NewUpdateOrder(repo repository.OrderRepository) *UpdateOrder {
	return &UpdateOrder{repo: repo}
}

func(uo *UpdateOrder) Run(id int, Producto string, Estado string, Pais string, Entidad_federativa string, Cp string) error {
	order, err := uo.repo.FindByID(id)
	if err != nil {
		return fmt.Errorf("pedido no encontrado: %w", err)
	}

    order.Producto = Producto
	order.Estado = Estado
	order.Pais = Pais
	order.Entidad_federativa = Entidad_federativa
	order.Cp = Cp

	if err := uo.repo.Update(*order); err != nil {
		return fmt.Errorf("Error actualizando el pedido: %w", err)
	}
    return nil
}