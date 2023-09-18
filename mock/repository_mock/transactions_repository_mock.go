package repository_mock

import (
	"Kelompok-2/dompet-online/model"
	"Kelompok-2/dompet-online/model/dto/resp"
	"github.com/stretchr/testify/mock"
)

type TransactionRepositoryMock struct {
	mock.Mock
}

func (t *TransactionRepositoryMock) FindAll(userId string) ([]resp.GetTransactionsResponse, error) {
	args := t.Called(userId)
	if args.Get(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]resp.GetTransactionsResponse), nil
}

func (t *TransactionRepositoryMock) Count(userId string) (int, error) {
	args := t.Called(userId)
	if args.Get(1) != nil {
		return 0, args.Error(1)
	}
	return 0, nil
}

func (t *TransactionRepositoryMock) FindWalletByUserID(userID string) (model.Wallet, error) {
	args := t.Called(userID)
	if args.Get(1) != nil {
		return model.Wallet{}, args.Error(1)
	}
	return args.Get(0).(model.Wallet), nil
}

func (t *TransactionRepositoryMock) UpdateWalletBalance(walletID string, amount int) error {
	return t.Called(walletID, amount).Error(0)
}

func (t *TransactionRepositoryMock) FindWalletByRekening(rekening string) (model.Wallet, error) {
	args := t.Called(rekening)
	if args.Get(1) != nil {
		return model.Wallet{}, args.Error(1)
	}
	return args.Get(0).(model.Wallet), nil
}

func (t *TransactionRepositoryMock) CreateTransaction(transaction model.Transactions) (model.Transactions, error) {
	args := t.Called(transaction)
	if args.Get(1) != nil {
		return model.Transactions{}, args.Error(1)
	}
	return args.Get(0).(model.Transactions), nil
}
