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

	rg.POST("/auth/register", a.registerHandler)
	rg.POST("/auth/login", a.loginHandler)
}

func (a *UserController) UsersRoute() {
	rg := a.engine.Group("/api/v1", middleware.AuthMiddleware())

	rg.GET("/users/:phoneNumber", a.findUserByPhoneNumberHandler)
	rg.GET("/users", a.listsUsersHandler)
	rg.PUT("/users", a.updateAccountHandler)
	rg.PATCH("/users", a.changePasswordAccountHandler)
	rg.PUT("/users/:id", a.disableAccountHandler)
}

// Auth //

// UserController godoc
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        Body body req.AuthRegisterRequest  true  "Auth register"
// @Success      201  {object}  resp.RegisterResponse
// @Router       /auth/register [post]
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

// UserController godoc
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        Body body req.AuthLoginRequest  true  "Auth login"
// @Success      200  {object}  resp.LoginResponse
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

// Admin //

// UserController godoc
// @Tags         Admin
// @Accept       json
// @Produce      json
// @Security	 Bearer
// @Success      200  {object}  model.Users
// @Router       /users [get]
func (a *UserController) listsUsersHandler(c *gin.Context) {
	users, err := a.userUC.ListsUsersHandler()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, users)
}

// UserController godoc
// @Tags         Admin
// @Accept       json
// @Produce      json
// @Security	 Bearer
// @Param		 phoneNumber path string true "User PhoneNumber"
// @Success      200  {object}  model.Users
// @Router       /users/{phoneNumber} [get]
func (a *UserController) findUserByPhoneNumberHandler(c *gin.Context) {
	phoneNumber := c.Param("phoneNumber")
	user, err := a.userUC.FindByUserByPhoneNumber(phoneNumber)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, user)
}

// user //

// UserController godoc
// @Tags         User
// @Accept       json
// @Produce      json
// @Security	 Bearer
// @Param        Body body req.UpdateAccountRequest  true  "Update Personal Information"
// @Success      200  {object}  resp.UpdateAccountRespone
// @Router       /users [put]
func (a *UserController) updateAccountHandler(c *gin.Context) {
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

	response := resp.UpdateAccountRespone{
		Status:  http.StatusOK,
		Message: "successfully Update account",
	}

	c.JSON(response.Status, response)
}

// UserController godoc
// @Tags         User
// @Accept       json
// @Produce      json
// @Security	 Bearer
// @Param        Body body req.UpdatePasswordRequest  true  "Change Password"
// @Success      200  {object}  resp.UpdatePasswordResponse
// @Router       /users [patch]
func (a *UserController) changePasswordAccountHandler(c *gin.Context) {
	var payload req.UpdatePasswordRequest
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	err := a.userUC.ChangePasswordAccount(payload)
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

// UserController godoc
// @Tags         User
// @Accept       json
// @Produce      json
// @Security	 Bearer
// @Param		 id path string true "Disable Account"
// @Success      200  {object}  resp.DisableAccountResponse
// @Router       /users/{id} [put]
func (a *UserController) disableAccountHandler(c *gin.Context) {
	UserId := c.Param("id")
	_, err := a.userUC.DisableAccount(UserId)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	response := resp.DisableAccountResponse{
		Status:  http.StatusOK,
		Message: "successfully disable account",
	}

	c.JSON(response.Status, response)
}
