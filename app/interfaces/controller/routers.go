package controller

import (
	"api/app/application/use_cases/user"
	"api/app/infrastructure/config"
	infraRepo "api/app/infrastructure/repository"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.Default()

	userRepo := infraRepo.NewUserRepository(config.DB)
	getUsersUC := user.NewGetUsersUseCase(userRepo)
	createUserUC := user.NewCreateUserUseCase(userRepo)
	loginUserUC := user.NewLoginUserUseCase(userRepo)
	userController := NewUserController(getUsersUC, createUserUC, loginUserUC)

	userGroup := router.Group("/users")
	userController.RegisterRoutes(userGroup)

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	return router
}
