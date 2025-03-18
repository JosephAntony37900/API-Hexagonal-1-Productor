package application

import (
	"fmt"
	"github.com/JosephAntony37900/API-Hexagonal-1-Productor/Users/domain/entities"
	"github.com/JosephAntony37900/API-Hexagonal-1-Productor/Users/domain/repository"
	helpers "github.com/JosephAntony37900/API-Hexagonal-1-Productor/helpers"
)

type CreateUsers struct {
	repo repository.UserRepository
}

func NewCreateUser(repo repository.UserRepository) *CreateUsers {
	return &CreateUsers{repo: repo}
}

func (cu *CreateUsers) Run(nombre string, email string, contraseña string) error {
	hashedPassword, err := helpers.HashPassword(contraseña)
	if err != nil {
		return fmt.Errorf("error al encriptar la contraseña: %w", err)
	}

	user := entities.Users{Nombre: nombre, Email: email, Contraseña: hashedPassword}
	if err := cu.repo.Save(user); err != nil {
		return fmt.Errorf("error al guardar el usuario: %w", err)
	}
	return nil
}
