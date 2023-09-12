package controller

import (
	"Kelompok-2/dompet-online/model"
	"Kelompok-2/dompet-online/usecase"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	userUC usecase.UserUseCase
	authUC usecase.AuthUseCase
	engine *gin.Engine
}

func (a *AuthController) loginHandler(c *gin.Context) {

}

func (a *AuthController) registerHandler(c *gin.Context) {

}

func (a *AuthController) changePasswordHandler(c *gin.Context) {

}

func (a *AuthController) UpdateUserNameHandler(c *gin.Context) {
	var users model.Users
	if err := c.ShouldBindJSON(&users); err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}
	err := a.userUC.UpdateUsername(users.UserName)
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "successfully update username",
	})

}

func (a *AuthController) Route() {
	rg := a.engine.Group("/api/v1")

	rg.POST("/auth/login", a.loginHandler)
	rg.POST("/auth/register", a.registerHandler)
	rg.POST("/auth/change-password", a.changePasswordHandler)
	rg.PUT("/users", a.UpdateUserNameHandler)
}

func NewAuthController(userUC usecase.UserUseCase, authUC usecase.AuthUseCase, engine *gin.Engine) *AuthController {
	return &AuthController{
		userUC: userUC,
		authUC: authUC,
		engine: engine,
	}
}
