package controller

import (
	"Kelompok-2/dompet-online/delivery/middleware"
	"Kelompok-2/dompet-online/exception"
	"Kelompok-2/dompet-online/model/dto/req"
	"Kelompok-2/dompet-online/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
)

type WalletController struct {
	walletUC usecase.WalletUseCase
	userUC   usecase.UserUseCase
	engine   *gin.Engine
}

func NewWalletController(walletUC usecase.WalletUseCase, userUC usecase.UserUseCase, engine *gin.Engine) *WalletController {
	return &WalletController{
		walletUC: walletUC,
		userUC:   userUC,
		engine:   engine,
	}
}

func (w *WalletController) createWallet(c *gin.Context) {
	var wallet req.WalletRequestBody
	if err := c.ShouldBindJSON(&wallet); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	err := w.walletUC.CreateWallet(wallet)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, wallet)
}

func (w *WalletController) getWalletByUserId(c *gin.Context) {
	userId := c.Param("userId")

	wallet, err := w.walletUC.FindByUserId(userId)
	if err != nil {
		exception.ErrorHandling(c, err)
		return
	}
	c.JSON(http.StatusOK, wallet)
}
func (w *WalletController) Route() {
	rg := w.engine.Group("/api/v1", middleware.AuthMiddleware())

	rg.POST("/wallets", w.createWallet)
	rg.GET("/wallets/:userId", w.getWalletByUserId)
}
