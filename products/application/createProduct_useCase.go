package application

import (
	"fmt"

	"github.com/JosephAntony37900/API-Hexagonal-1-Productor/products/domain/entities"
	"github.com/JosephAntony37900/API-Hexagonal-1-Productor/products/domain/repository"
)

type CreateProduct struct {
	repo repository.ProductRepository
}

func NewCreateProduct(repo repository.ProductRepository) *CreateProduct {
	return &CreateProduct{repo: repo}
}

func (cp *CreateProduct) Run(nombre string, precio float64, cantidad int) error {
	product := entities.Product{Nombre: nombre, Precio: precio, Cantidad: cantidad}
	if err := cp.repo.Save(product); err != nil {
		return fmt.Errorf("error guardando el producto: %w", err)
	}
	return nil
}
