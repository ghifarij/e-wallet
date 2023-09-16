package controller

import (
	"Kelompok-2/dompet-online/delivery/middleware"
	"Kelompok-2/dompet-online/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
)

type WalletController struct {
	walletUC usecase.WalletUseCase
	engine   *gin.Engine
}

func NewWalletController(walletUC usecase.WalletUseCase, engine *gin.Engine) *WalletController {
	return &WalletController{
		walletUC: walletUC,
		engine:   engine,
	}
}

func (w *WalletController) Route() {
	rg := w.engine.Group("/api/v1", middleware.AuthMiddleware())

	rg.GET("/wallets/:userId", w.getWalletByUserId)
}

func (w *WalletController) getWalletByUserId(c *gin.Context) {
	userId := c.Param("userId")

	wallet, err := w.walletUC.GetWalletByUserId(userId)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, wallet)
}
