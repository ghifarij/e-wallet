package controller

import (
	"Kelompok-2/dompet-online/delivery/middleware"
	"Kelompok-2/dompet-online/model/dto/req"
	"Kelompok-2/dompet-online/model/dto/resp"
	"Kelompok-2/dompet-online/usecase"
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
	rg.PUT("/users/update-account", a.updateAccount)
	rg.PATCH("/users/change-password", a.changePasswordHandler)
	rg.PUT("/users/:userId", a.DisableUserID)
}

// Auth

// UserController godoc
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        Body body req.AuthLoginRequest  true  "Auth login"
// @Success      200  {object}  req.AuthLoginRequest
// @Router       /auth/login [post]
func (a *UserController) loginHandler(c *gin.Context) {
	var payload req.AuthLoginRequest
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	authResponse, err := a.userUC.Login(payload)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
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

func (a *UserController) registerHandler(c *gin.Context) {
	var payload req.AuthRegisterRequest
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	err := a.userUC.Register(payload)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
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
	users, err := a.userUC.ListsHandler()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, users)
}

func (a *UserController) findByPhoneNumber(c *gin.Context) {
	phoneNumber := c.Param("phoneNumber")
	user, err := a.userUC.FindByPhoneNumber(phoneNumber)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (a *UserController) updateAccount(c *gin.Context) {
	var payload req.UpdateAccountRequest
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	err := a.userUC.UpdateAccount(payload)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	response := resp.UpdateUserNameRespone{
		Status:  http.StatusOK,
		Message: "successfully Update account",
	}

	c.JSON(response.Status, response)
}

func (a *UserController) changePasswordHandler(c *gin.Context) {
	var payload req.UpdatePasswordRequest
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	err := a.userUC.ChangePassword(payload)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
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

func (a *UserController) DisableUserID(c *gin.Context) {
	DisabeleUserId := c.Param("userId")
	_, err := a.userUC.DisableUserId(DisabeleUserId)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "successfully disable account"})
}
