package controller

import (
	"Kelompok-2/dompet-online/model/dto/req"
	"Kelompok-2/dompet-online/model/dto/resp"
	"Kelompok-2/dompet-online/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	userUC usecase.UserUseCase
	authUC usecase.AuthUseCase
	engine *gin.Engine
}

func (a *AuthController) loginHandler(c *gin.Context) {
	var payload req.AuthLoginRequest
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	authResponse, err := a.authUC.Login(payload)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	response := resp.LoginResponse{
		Status:   http.StatusOK,
		UserName: authResponse.UserName,
		Message:  "successfully login",
		Token:    authResponse.Token,
	}

	c.JSON(response.Status, response)
}

func (a *AuthController) registerHandler(c *gin.Context) {
	var payload req.AuthRegisterRequest
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	err := a.userUC.Register(payload)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	response := resp.RegisterResponse{
		Status:  http.StatusCreated,
		Message: "successfully register",
	}

	c.JSON(response.Status, response)
}

func (a *AuthController) changePasswordHandler(c *gin.Context) {
	var payload req.UpdatePasswordRequest
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	err := a.authUC.ChangePassword(payload)
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	response := resp.UpdatePasswordResponse{
		Status:  http.StatusOK,
		Message: "successfully change password",
	}

	c.JSON(response.Status, response)
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

func (a *AuthController) updateUsername(c *gin.Context) {
	var payload req.UpdateUserNameRequest
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	err := a.userUC.UpdateUsername(payload)
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	response := resp.UpdateUserNameRespone{
		Status:  http.StatusOK,
		Message: "successfully Update username",
	}

	c.JSON(response.Status, response)
}

func (a *AuthController) Route() {
	rg := a.engine.Group("/api/v1")

	rg.POST("/auth/login", a.loginHandler)
	rg.POST("/auth/register", a.registerHandler)
	rg.PATCH("/auth/change-password", a.changePasswordHandler)
	rg.GET("/users/:phoneNumber", a.findByPhoneNumber)
	rg.PUT("/users/update-username", a.updateUsername)
}

func NewAuthController(userUC usecase.UserUseCase, authUC usecase.AuthUseCase, engine *gin.Engine) *AuthController {
	return &AuthController{
		userUC: userUC,
		authUC: authUC,
		engine: engine,
	}
}
