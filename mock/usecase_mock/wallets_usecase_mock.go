package usecase_mock

import (
	"Kelompok-2/dompet-online/model"
	"github.com/stretchr/testify/mock"
)

type WalletUseCaseMock struct {
	mock.Mock
}

func (w *WalletUseCaseMock) UpdateWalletBalance(walletID string, amount int) error {
	return w.Called(walletID, amount).Error(0)
}

func (w *WalletUseCaseMock) GetWalletByUserId(userId string) (model.Wallet, error) {
	args := w.Called(userId)
	if args.Get(1) != nil {
		return model.Wallet{}, args.Error(1)
	}
	return args.Get(0).(model.Wallet), nil
}

func (w *WalletUseCaseMock) GetWalletByRekeningUser(number string) (model.Wallet, error) {
	args := w.Called(number)
	if args.Get(1) != nil {
		return model.Wallet{}, args.Error(1)
	}
	return args.Get(0).(model.Wallet), nil
}

func (w *WalletUseCaseMock) CreateWallet(payload model.Wallet) error {
	args := w.Called(payload)
	if args.Get(0) != nil {
		return args.Error(0)
	}
	return nil
}
