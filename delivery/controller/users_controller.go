package controller

import (
	"Kelompok-2/dompet-online/delivery/middleware"
	"Kelompok-2/dompet-online/exception"
	"Kelompok-2/dompet-online/model/dto/req"
	"Kelompok-2/dompet-online/model/dto/resp"
	"Kelompok-2/dompet-online/usecase"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userUC   usecase.UserUseCase
	walletUC usecase.WalletUseCase
	engine   *gin.Engine
}

func NewUserController(userUC usecase.UserUseCase, walletUC usecase.WalletUseCase, engine *gin.Engine) *UserController {
	return &UserController{
		userUC:   userUC,
		walletUC: walletUC,
		engine:   engine,
	}
}

// Route
func (a *UserController) AuthRoute() {
	rg := a.engine.Group("/api/v1")

	rg.POST("/auth/login", a.loginHandler)
	rg.POST("/auth/register", a.registerHandler)
}

func (a *UserController) UsersRoute() {
	rg := a.engine.Group("/api/v1", middleware.AuthMiddleware())

	rg.GET("/users/:phoneNumber", a.findByPhoneNumber)
	rg.GET("/users", a.listHandler)
	rg.PUT("/users/update-username", a.updateUsername)
	rg.DELETE("/users/:id", a.updateUsername)
	rg.PATCH("/users/change-password", a.changePasswordHandler)
	rg.GET("/wallets/:userId", a.getWalletByUserId)

}

// Auth
func (a *UserController) loginHandler(c *gin.Context) {
	var payload req.AuthLoginRequest
	if err := c.ShouldBindJSON(&payload); err != nil {
		exception.ErrorHandling(c, err)
		return
	}

	authResponse, err := a.userUC.Login(payload)
	if err != nil {
		exception.ErrorHandling(c, err)
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

func (a *UserController) registerHandler(c *gin.Context) {
	var payload req.AuthRegisterRequest
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	err := a.userUC.Register(payload)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{
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

// Users
func (a *UserController) listHandler(c *gin.Context) {
	users, err := a.userUC.FindAll()
	if err != nil {
		exception.ErrorHandling(c, err)
		return
	}
	c.JSON(200, users)
}

func (a *UserController) findByPhoneNumber(c *gin.Context) {
	phoneNumber := c.Param("phoneNumber")
	user, err := a.userUC.FindByPhoneNumber(phoneNumber)

	if err != nil {
		exception.ErrorHandling(c, err)
		return
	}

	c.JSON(http.StatusOK, user)
}

func (a *UserController) updateUsername(c *gin.Context) {
	var payload req.UpdateUserNameRequest
	if err := c.ShouldBindJSON(&payload); err != nil {
		exception.ErrorHandling(c, err)
		return
	}

	err := a.userUC.UpdateUsername(payload)
	if err != nil {
		exception.ErrorHandling(c, err)
		return
	}

	response := resp.UpdateUserNameRespone{
		Status:  http.StatusOK,
		Message: "successfully Update username",
	}

	c.JSON(response.Status, response)
}

func (a *UserController) deleteById(c *gin.Context) {
	id := c.Param("id")

	err := a.userUC.DeleteById(id)
	if err != nil {
		exception.ErrorHandling(c, err)
		return
	}

	message := fmt.Sprintf("successfully delete user with id %s", id)
	c.JSON(200, gin.H{
		"message": message,
	})
}

func (a *UserController) changePasswordHandler(c *gin.Context) {
	var payload req.UpdatePasswordRequest
	if err := c.ShouldBindJSON(&payload); err != nil {
		exception.ErrorHandling(c, err)
		return
	}

	err := a.userUC.ChangePassword(payload)
	if err != nil {
		exception.ErrorHandling(c, err)
		return
	}

	response := resp.UpdatePasswordResponse{
		Status:  http.StatusOK,
		Message: "successfully change password",
	}

	c.JSON(response.Status, response)
}

func (a *UserController) getWalletByUserId(c *gin.Context) {
	userId := c.Param("userId")

	wallet, err := a.walletUC.GetWalletByUserId(userId)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, wallet)
}
