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
	repo     repository.TransactionRepository
	walletUC WalletUseCase
}

func NewTransactionUseCase(repo repository.TransactionRepository, walletUC WalletUseCase) TransactionUseCase {
	return &transactionUseCase{repo: repo, walletUC: walletUC}
}

func (t *transactionUseCase) GetHistoryTransactions(userId string) ([]resp.GetTransactionsResponse, error) {
	getTransactionsResponses, err := t.repo.FindAll(userId)
	if err != nil {
		return []resp.GetTransactionsResponse{}, err
	}

	return getTransactionsResponses, nil
}

func (t *transactionUseCase) TopUp(payload req.TopUpRequest) (model.Transactions, error) {
	transaction := model.Transactions{
		Id:              common.GenerateID(),
		UserId:          payload.UserId,
		SourceWalletID:  payload.WalletID,
		Destination:     "TopUp",
		Amount:          payload.Amount,
		Description:     "TopUp",
		PaymentMethodID: payload.PaymentMethodId,
		CreateAt:        time.Now(),
	}

	createdTransaction, err := t.repo.CreateTransaction(transaction)
	if err != nil {
		return model.Transactions{}, err
	}

	err = t.walletUC.UpdateWalletBalance(payload.WalletID, payload.Amount)
	if err != nil {
		return model.Transactions{}, err
	}

	return createdTransaction, nil
}

func (t *transactionUseCase) Transfer(payload req.TransferRequest) (model.Transactions, error) {
	transaction := model.Transactions{
		Id:              common.GenerateID(),
		UserId:          payload.UserId,
		SourceWalletID:  payload.SourceWalletID,
		Destination:     payload.DestinationWalletID,
		Amount:          payload.Amount,
		Description:     payload.Description,
		PaymentMethodID: payload.PaymentMethodID,
	}

	createdTransaction, err := t.repo.CreateTransaction(transaction)
	if err != nil {
		return model.Transactions{}, err
	}

	//amount := -payload.Amount
	err = t.walletUC.UpdateWalletBalance(payload.SourceWalletID, -payload.Amount)
	if err != nil {
		return model.Transactions{}, err
	}

	err = t.walletUC.UpdateWalletBalance(payload.DestinationWalletID, payload.Amount)
	if err != nil {
		return model.Transactions{}, err
	}

	return createdTransaction, nil
}

func (t *transactionUseCase) CountTransaction(userId string) (int, error) {
	count, err := t.repo.Count(userId)
	if err != nil {
		return 0, err
	}

	return count, nil

}
