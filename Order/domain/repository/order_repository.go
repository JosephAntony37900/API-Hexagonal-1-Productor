package repository

import "github.com/JosephAntony37900/API-Hexagonal-1-Productor/Order/domain/entities"

type OrderRepository interface {
	Save(order entities.Order) error
	FindByID(id int) (*entities.Order, error)
	FindByUserID(usuario_id int) ([]entities.Order, error)
	//FindAll() ([]entities.Order, error)
	Update(order entities.Order) error
	Delete(id int) error
}