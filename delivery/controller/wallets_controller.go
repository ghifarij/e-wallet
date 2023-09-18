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

	rg.GET("/wallets/:userId", w.getWalletByUserIdHandler)
}

// WalletController godoc
// @Tags         Wallet
// @Accept       json
// @Produce      json
// @Security	 Bearer
// @Param		 userId path string true "Get Wallet"
// @Success      200  {object}  model.Wallet
// @Router       /wallets/{userId} [get]
func (w *WalletController) getWalletByUserIdHandler(c *gin.Context) {
	userId := c.Param("userId")

	wallet, err := w.walletUC.GetWalletByUserId(userId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, wallet)
}
