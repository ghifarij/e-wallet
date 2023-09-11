package controller

import (
	"Kelompok-2/dompet-online/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
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

func (a *AuthController) findByPhoneNumber(c *gin.Context) {
	phoneNumber := c.Param("phoneNumber")
	user, err := a.userUC.FindByPhoneNumber(phoneNumber)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (a *AuthController) Route() {
	rg := a.engine.Group("/api/v1")

	rg.POST("/auth/login", a.loginHandler)
	rg.POST("/auth/register", a.registerHandler)
	rg.POST("/auth/change-password", a.changePasswordHandler)
	rg.GET("/users/:phoneNumber", a.findByPhoneNumber)
}

func NewAuthController(userUC usecase.UserUseCase, authUC usecase.AuthUseCase, engine *gin.Engine) *AuthController {
	return &AuthController{
		userUC: userUC,
		authUC: authUC,
		engine: engine,
	}
}
