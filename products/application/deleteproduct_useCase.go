package application

import (
	"fmt"

	"github.com/JosephAntony37900/API-Hexagonal-1-Productor/products/domain/repository"
)

type DeleteProduct struct {
	repo repository.ProductRepository
}

func NewDeleteProduct(repo repository.ProductRepository) *DeleteProduct {
	return &DeleteProduct{repo: repo}
}

func (dp *DeleteProduct) Run(id int) error {
	_, err := dp.repo.FindByID(id)
	if err != nil {
		return fmt.Errorf("product no encontrado: %w", err)
	}

	if err := dp.repo.Delete(id); err != nil {
		return fmt.Errorf("error eliminando el producto: %w", err)
	}

	return nil
}
