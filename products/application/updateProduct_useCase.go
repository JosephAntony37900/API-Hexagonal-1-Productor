package application

import (
	"fmt"

	_ "github.com/JosephAntony37900/API-Hexagonal-1-Productor/products/domain/entities"
	"github.com/JosephAntony37900/API-Hexagonal-1-Productor/products/domain/repository"
)

type UpdateProduct struct {
	repo repository.ProductRepository
}

func NewUpdateProduct(repo repository.ProductRepository) *UpdateProduct {
	return &UpdateProduct{repo: repo}
}

func (up *UpdateProduct) Run(id int, nombre string, precio float64, cantidad int) error {
	// Verificar si el producto existe
	product, err := up.repo.FindByID(id)
	if err != nil {
		return fmt.Errorf("producto no encontrado: %w", err)
	}

	// Actualizar los campos del producto
	product.Nombre = nombre
	product.Precio = precio
	product.Cantidad = cantidad

	// Guardar los cambios en el repositorio
	if err := up.repo.Update(*product); err != nil {
		return fmt.Errorf("error actualizando el producto: %w", err)
	}

	return nil
}
