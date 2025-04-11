package service

import (
	"os"

	"github.com/JosephAntony37900/API-Hexagonal-1-Productor/Users/domain/services"
	"github.com/JosephAntony37900/API-Hexagonal-1-Productor/Users/infraestructure/adapters"
)

// Inicializar el servicio de BCrypt
func InitBcryptService() services.IBcrypService {
	return adapters.NewBcrypt()
}

// Inicializar el Token Manager
func InitTokenManager() services.TokenManager {
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		panic("JWT_SECRET no está configurado en las variables de entorno")
	}
	return &adapters.JWTManager{SecretKey: jwtSecret}
}