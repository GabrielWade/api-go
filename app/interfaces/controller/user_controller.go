package controller

import (
	dtoUser "api/app/application/dto/user"
	"api/app/application/use_cases/user"
	"api/app/domain/models"
	"api/app/interfaces/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	getUsersUseCase   *user.GetUsersUseCase
	createUserUseCase *user.CreateUserUseCase
	loginUserUseCase  *user.LoginUserUseCase
}

func NewUserController(
	getUsersUC *user.GetUsersUseCase,
	createUserUC *user.CreateUserUseCase,
	loginUserUC *user.LoginUserUseCase,
) *UserController {
	return &UserController{
		getUsersUseCase:   getUsersUC,
		createUserUseCase: createUserUC,
		loginUserUseCase:  loginUserUC,
	}
}

func (ctrl *UserController) RegisterRoutes(rg *gin.RouterGroup) {
	rg.GET("/", middleware.CheckAuth, ctrl.GetUsers)
	rg.POST("/", ctrl.CreateUser)
	rg.POST("/login", ctrl.LoginUser)
}

func (ctrl *UserController) GetUsers(c *gin.Context) {
	users, err := ctrl.getUsersUseCase.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"users": users})
}

func (ctrl *UserController) CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := ctrl.createUserUseCase.Execute(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, user)
}

func (ctrl *UserController) LoginUser(c *gin.Context) {
	var req dtoUser.LoginRequestDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	accessToken, refreshToken, err := ctrl.loginUserUseCase.Execute(req.Username, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"access_token": accessToken, "refresh_token": refreshToken})
}
