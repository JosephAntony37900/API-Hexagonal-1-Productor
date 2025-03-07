package repository

import "github.com/JosephAntony37900/API-Hexagonal-1-Productor/Users/domain/entities"

type UserRepository interface {
	Save(user entities.Users) error
	FindByID(id int) (*entities.Users, error)
	FindAll() ([]entities.Users, error)
	Update(user entities.Users) error
	Delete(id int)error
}