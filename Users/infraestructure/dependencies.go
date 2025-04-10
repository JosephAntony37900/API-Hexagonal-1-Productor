package infraestructure

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	userApp "github.com/JosephAntony37900/API-Hexagonal-1-Productor/Users/application"
    userController "github.com/JosephAntony37900/API-Hexagonal-1-Productor/Users/infraestructure/controllers"
    userRepo "github.com/JosephAntony37900/API-Hexagonal-1-Productor/Users/infraestructure/repository"
    userRoutes "github.com/JosephAntony37900/API-Hexagonal-1-Productor/Users/infraestructure/routes"
)

func InitUsersDependencies(Engine *gin.Engine, db *sql.DB){
	userRepository := userRepo.NewCreateUserRepoMySQL(db)

	createUser := userApp.NewCreateUser(userRepository)
    getUsers := userApp.NewGetUsers(userRepository)
    deleteUsers := userApp.NewDeleteUser(userRepository)
    updateUsers := userApp.NewUpdateUser(userRepository)
    loginUser := userApp.NewLoginUser(userRepository)     

	createUserController := userController.NewCreateUserController(createUser)
    getUserController := userController.NewUsersController(getUsers)
    deleteUserController := userController.NewDeleteUserController(deleteUsers)
    updateUserController := userController.NewUpdateUserController(updateUsers)
    loginUserController := userController.NewLoginUserController(loginUser)

	userRoutes.SetupUserRoutes(Engine, createUserController, loginUserController, getUserController, deleteUserController, updateUserController)

}