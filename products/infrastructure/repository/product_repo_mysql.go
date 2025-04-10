package repository

import (
	"database/sql"
	"fmt"

	"github.com/JosephAntony37900/API-Hexagonal-1-Productor/products/domain/entities"
)

type ProductRepoMySQL struct {
	db *sql.DB
}

func NewProductRepoMySQL(db *sql.DB) *ProductRepoMySQL {
	return &ProductRepoMySQL{db: db}
}

func (r *ProductRepoMySQL) Save(product entities.Product) error {
	query := "INSERT INTO productos (Nombre, Precio, Cantidad) VALUES (?, ?, ?)"
	_, err := r.db.Exec(query, product.Nombre, product.Precio, product.Cantidad)
	if err != nil {
		return fmt.Errorf("error insertando el producto: %w", err)
	}
	return nil
}

func (r *ProductRepoMySQL) FindByID(id int) (*entities.Product, error) {
	query := "SELECT Id, Nombre, Precio, Cantidad FROM productos WHERE Id = ?"
	row := r.db.QueryRow(query, id)

	var product entities.Product
	if err := row.Scan(&product.Id, &product.Nombre, &product.Precio, &product.Cantidad); err != nil {
		return nil, fmt.Errorf("error buscando el producto: %w", err)
	}
	return &product, nil
}

func (r *ProductRepoMySQL) FindAll() ([]entities.Product, error) {
	query := "SELECT Id, Nombre, Precio, Cantidad FROM productos"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error buscando el producto: %w", err)
	}
	defer rows.Close()

	var products []entities.Product
	for rows.Next() {
		var product entities.Product
		if err := rows.Scan(&product.Id, &product.Nombre, &product.Precio, &product.Cantidad); err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}

func (r *ProductRepoMySQL) Update(product entities.Product) error {
	query := "UPDATE productos SET Nombre = ?, Precio = ? WHERE Id = ?"
	_, err := r.db.Exec(query, product.Nombre, product.Precio, product.Id, product.Cantidad)
	if err != nil {
		return fmt.Errorf("error actualizando el producto: %w", err)
	}
	return nil
}

func (r *ProductRepoMySQL) Delete(id int) error {
	query := "DELETE FROM productos WHERE Id = ?"
	_, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("error eliminando el producto: %w", err)
	}
	return nil
}

func (r *ProductRepoMySQL) FindByMinimumPrice(minPrice float64) ([]entities.Product, error) {
	query := "SELECT Id, Nombre, Precio, Cantidad FROM productos WHERE Precio >= ?"
	rows, err := r.db.Query(query, minPrice)
	if err != nil {
		return nil, fmt.Errorf("error buscando los productos con precio minimo: %w", err)
	}
	defer rows.Close()

	var products []entities.Product
	for rows.Next() {
		var product entities.Product
		if err := rows.Scan(&product.Id, &product.Nombre, &product.Precio, &product.Cantidad); err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}
