package infraestructure

import (
	"database/sql"
	"github.com/gin-gonic/gin"
    "os"
	userApp "github.com/JosephAntony37900/API-Hexagonal-1-Productor/Users/application"
    userController "github.com/JosephAntony37900/API-Hexagonal-1-Productor/Users/infraestructure/controllers"
    userRepo "github.com/JosephAntony37900/API-Hexagonal-1-Productor/Users/infraestructure/repository"
    userRoutes "github.com/JosephAntony37900/API-Hexagonal-1-Productor/Users/infraestructure/routes"
    "github.com/JosephAntony37900/API-Hexagonal-1-Productor/Users/infraestructure/services"
)

func InitUsersDependencies(Engine *gin.Engine, db *sql.DB){
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		panic("JWT_SECRET no está configurado en las variables de entorno")
	}

	bcryptService := service.InitBcryptService()
    jwtManager := service.InitTokenManager()

    
	userRepository := userRepo.NewCreateUserRepoMySQL(db)

	createUser := userApp.NewCreateUser(userRepository, bcryptService)
    getUsers := userApp.NewGetUsers(userRepository)
    deleteUsers := userApp.NewDeleteUser(userRepository)
    updateUsers := userApp.NewUpdateUser(userRepository)
    loginUser := userApp.NewLoginUser(userRepository, jwtManager, bcryptService)     

	createUserController := userController.NewCreateUserController(createUser)
    getUserController := userController.NewUsersController(getUsers)
    deleteUserController := userController.NewDeleteUserController(deleteUsers)
    updateUserController := userController.NewUpdateUserController(updateUsers)
    loginUserController := userController.NewLoginUserController(loginUser)

	userRoutes.SetupUserRoutes(Engine, createUserController, loginUserController, getUserController, deleteUserController, updateUserController)

}