package usecase

import (
	"Kelompok-2/dompet-online/model"
	"Kelompok-2/dompet-online/model/dto/req"
	"Kelompok-2/dompet-online/model/dto/resp"
	"Kelompok-2/dompet-online/repository"
	"Kelompok-2/dompet-online/util/common"
	"fmt"
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
	//
	panic("")
}

func (t *transactionUseCase) Transfer(payload req.TransferRequest) (model.Transactions, error) {
	// Fetch the current balances of the sender and receiver from the database
	senderWallet, err := t.repo.FindWalletByUserID(payload.UserId)
	if err != nil {
		return model.Transactions{}, err
	}

	receiverWallet, err := t.repo.FindWalletByRekening(payload.Destination)
	if err != nil {
		return model.Transactions{}, err
	}

	// Check if the sender has enough balance for the transfer
	if senderWallet.Balance < payload.Amount {
		return model.Transactions{}, fmt.Errorf("insufficient balance for the transfer")
	}

	// Create a transfer transaction record in the database
	transaction := model.Transactions{
		Id:              common.GenerateID(),
		UserId:          payload.UserId,
		SourceOfFoundId: "Wallet",
		Destination:     payload.Destination,
		Amount:          payload.Amount,
		Description:     payload.Description,
	}

	createdTransaction, err := t.repo.CreateTransaction(transaction)
	if err != nil {
		return model.Transactions{}, err
	}

	// Update sender's wallet balance
	senderWallet.Balance -= payload.Amount
	err = t.repo.UpdateWalletBalance(senderWallet)
	if err != nil {
		return model.Transactions{}, err
	}

	// Update receiver's wallet balance
	receiverWallet.Balance += payload.Amount
	err = t.repo.UpdateWalletBalance(receiverWallet)
	if err != nil {
		return model.Transactions{}, err
	}

	return createdTransaction, nil

}

func (t *transactionUseCase) CountTransaction(userId string) (int, error) {
	//TODO implement me
	panic("implement me")
}

func NewTransactionUseCase(repo repository.TransactionRepository) TransactionUseCase {
	return &transactionUseCase{repo: repo}
}
