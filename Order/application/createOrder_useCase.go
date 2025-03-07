package application

import (
	"fmt"

	"github.com/JosephAntony37900/API-Hexagonal-1-Productor/Order/domain/entities"
	"github.com/JosephAntony37900/API-Hexagonal-1-Productor/Order/domain/repository"
)

type CreateOrder struct {
	repo repository.OrderRepository
}

// Constructor del caso de uso
func NewCreateOrder(repo repository.OrderRepository) *CreateOrder {
	return &CreateOrder{repo: repo}
}

// Método para ejecutar la lógica de creación de pedido
func (co *CreateOrder) Run(Usuario_id int, Producto string, Pais string, Entidad_federativa string, Cp string) error {
	// Se crea la orden con estado "Pendiente"
	order := entities.Order{
		Usuario_id:        Usuario_id,
		Producto:          Producto,
		Estado:            "Pendiente",
		Pais:             Pais,
		Entidad_federativa: Entidad_federativa,
		Cp:               Cp,
	}

	// Guardar en el repositorio
	if err := co.repo.Save(order); err != nil {
		return fmt.Errorf("error guardando el pedido: %w", err)
	}

	return nil
}
