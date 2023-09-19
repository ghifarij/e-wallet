package usecase_mock

import (
	"Kelompok-2/dompet-online/model"
	"Kelompok-2/dompet-online/model/dto/req"
	"Kelompok-2/dompet-online/model/dto/resp"
	"github.com/stretchr/testify/mock"
)

type TransactionUseCaseMock struct {
	mock.Mock
}

func (t *TransactionUseCaseMock) GetHistoriesTransactions(userId string) ([]resp.GetTransactionsResponse, error) {
	args := t.Called(userId)
	if args.Get(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]resp.GetTransactionsResponse), nil
}

func (t *TransactionUseCaseMock) TopUp(payload req.TopUpRequest) (model.Transactions, error) {
	args := t.Called(payload)
	if args.Get(1) != nil {
		return model.Transactions{}, args.Error(1)
	}
	return args.Get(0).(model.Transactions), nil
}

func (t *TransactionUseCaseMock) Transfer(payload req.TransferRequest) (model.Transactions, error) {
	args := t.Called(payload)
	if args.Get(1) != nil {
		return model.Transactions{}, args.Error(1)
	}
	return args.Get(0).(model.Transactions), nil
}

func (t *TransactionUseCaseMock) CountTransaction(userId string) (int, error) {
	args := t.Called(userId)
	if args.Get(1) != nil {
		return 0, args.Error(1)
	}
	return 0, nil
}
