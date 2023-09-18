package repository_mock

import (
	"Kelompok-2/dompet-online/model"
	"github.com/stretchr/testify/mock"
)

type WalletRepositoryMock struct {
	mock.Mock
}

func (w *WalletRepositoryMock) FindByUserId(userid string) (model.Wallet, error) {
	args := w.Called(userid)
	if args.Get(1) != nil {
		return model.Wallet{}, args.Error(1)
	}
	return args.Get(0).(model.Wallet), nil
}

func (w *WalletRepositoryMock) FindByRekeningUser(number string) (model.Wallet, error) {
	args := w.Called(number)
	if args.Get(1) != nil {
		return model.Wallet{}, args.Error(1)
	}
	return args.Get(0).(model.Wallet), nil
}

func (w *WalletRepositoryMock) Save(wallet model.Wallet) error {
	return w.Called(wallet).Error(0)
}

func (w *WalletRepositoryMock) UpdateWalletBalance(walletID string, amount int) error {
	return w.Called(walletID, amount).Error(0)
}
