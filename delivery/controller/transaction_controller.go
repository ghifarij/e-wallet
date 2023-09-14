package controller

import (
	"Kelompok-2/dompet-online/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
)

type TransactionController struct {
	transactionUC usecase.TransactionUseCase
	engine        *gin.Engine
}

func NewTransactionController(transactionUC usecase.TransactionUseCase, engine *gin.Engine) *TransactionController {
	return &TransactionController{
		transactionUC: transactionUC,
		engine:        engine,
	}
}

func (t *TransactionController) Route() {
	rg := t.engine.Group("/api/v1")

	rg.GET("/transactions-history/:userId", t.getTransactionsHistory)
}

func (t *TransactionController) getTransactionsHistory(c *gin.Context) {
	userId := c.Param("userId")

	getTransactionhistory, err := t.transactionUC.GetHistoryTransactions(userId)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, getTransactionhistory)
}
