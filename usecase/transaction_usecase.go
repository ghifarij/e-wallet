package usecase

import (
	"Kelompok-2/dompet-online/model"
	"Kelompok-2/dompet-online/model/dto/req"
	"Kelompok-2/dompet-online/model/dto/resp"
	"Kelompok-2/dompet-online/repository"
	"Kelompok-2/dompet-online/util/common"
	"time"
)

type TransactionUseCase interface {
	GetHistoryTransactions(userId string) ([]resp.GetTransactionsResponse, error)
	TopUp(payload req.TopUpRequest) (model.Transactions, error)
	Transfer(payload req.TransferRequest) (model.Transactions, error)
	CountTransaction(userId string) (int, error)
}

type transactionUseCase struct {
	repo repository.TransactionRepository
}

func (t *transactionUseCase) GetHistoryTransactions(userId string) ([]resp.GetTransactionsResponse, error) {
	getTransactionsResponses, err := t.repo.FindAll(userId)
	if err != nil {
		return []resp.GetTransactionsResponse{}, err
	}

	return getTransactionsResponses, nil
}

func (t *transactionUseCase) TopUp(payload req.TopUpRequest) (model.Transactions, error) {
	// Create a new transaction record
	transaction := model.Transactions{
		Id:              common.GenerateID(),
		UserId:          payload.UserId,
		SourceWalletID:  payload.WalletID,
		Destination:     "TopUp", // TopUp transaction does not have a destination
		Amount:          payload.Amount,
		Description:     "TopUp",
		PaymentMethodID: payload.PaymentMethodId,
		CreateAt:        time.Now(),
	}

	// Create the transaction in the database
	createdTransaction, err := t.repo.CreateTransaction(transaction)
	if err != nil {
		return model.Transactions{}, err
	}

	// Update the user's wallet balance
	err = t.repo.UpdateWalletBalance(payload.WalletID, payload.Amount)
	if err != nil {
		return model.Transactions{}, err
	}

	return createdTransaction, nil
}

func (t *transactionUseCase) Transfer(payload req.TransferRequest) (model.Transactions, error) {
	//TODO implement me
	panic("implement me")
}

func (t *transactionUseCase) CountTransaction(userId string) (int, error) {
	//TODO implement me
	panic("implement me")
}

func NewTransactionUseCase(repo repository.TransactionRepository) TransactionUseCase {
	return &transactionUseCase{repo: repo}
}
