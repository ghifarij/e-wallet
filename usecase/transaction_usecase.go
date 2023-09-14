package usecase

import (
	"Kelompok-2/dompet-online/model"
	"Kelompok-2/dompet-online/model/dto/req"
	"Kelompok-2/dompet-online/model/dto/resp"
	"Kelompok-2/dompet-online/repository"
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
	//TODO implement me
	panic("implement me")
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
