package controller

import (
	"Kelompok-2/dompet-online/delivery/middleware"
	"Kelompok-2/dompet-online/model/dto/req"
	"Kelompok-2/dompet-online/model/dto/resp"
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
	rg := t.engine.Group("/api/v1", middleware.AuthMiddleware())

	rg.GET("/transactions/:userId", t.getHistoriesTransactionsHandler)
	rg.PUT("/transactions/transfer", t.transferTransaction)
	rg.PUT("/transactions/topUp", t.topUpTransaction)
	rg.GET("/transactions/count/:userId", t.CountTransaction)
}

// TransactionController godoc
// @Tags         Transaction
// @Accept       json
// @Produce      json
// @Security     Bearer
// @Param        userId path string true "Get History Transaction"
// @Success      200  {object}  resp.GetTransactionsResponse
// @Router       /transactions/{userId} [get]
func (t *TransactionController) getHistoriesTransactionsHandler(c *gin.Context) {
	userId := c.Param("userId")

	getHistoryTransaction, err := t.transactionUC.GetHistoriesTransactions(userId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	// Create a response object
	var response []resp.GetTransactionsResponse

	// Populate the response object with data from getHistoryTransaction
	for _, transaction := range getHistoryTransaction {
		transactionResponse := resp.GetTransactionsResponse{
			Id:            transaction.Id,
			Destination:   transaction.Destination,
			Amount:        transaction.Amount,
			Description:   transaction.Description,
			CreateAt:      transaction.CreateAt,
			User:          transaction.User,
			Wallet:        transaction.Wallet,
			PaymentMethod: transaction.PaymentMethod,
		}

		response = append(response, transactionResponse)
	}

	c.JSON(http.StatusOK, response)
}

func (t *TransactionController) transferTransaction(c *gin.Context) {
	var transferRequest req.TransferRequest
	if err := c.ShouldBindJSON(&transferRequest); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	transaction, err := t.transactionUC.Transfer(transferRequest)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, transaction)
}

func (t *TransactionController) topUpTransaction(c *gin.Context) {
	var topUpRequest req.TopUpRequest
	if err := c.ShouldBindJSON(&topUpRequest); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	transaction, err := t.transactionUC.TopUp(topUpRequest)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, transaction)
}

func (t *TransactionController) CountTransaction(c *gin.Context) {
	userId := c.Param("userId")

	countTransaction, err := t.transactionUC.CountTransaction(userId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, countTransaction)
}
