package usecase

import (
	"Kelompok-2/dompet-online/model"
	"Kelompok-2/dompet-online/repository"
)

type WalletUseCase interface {
	GetWalletByUserId(userId string) (model.Wallet, error)
	CreateWallet(payload model.Wallet) error
	UpdateWalletBalance(walletID string, amount int) error
}

type walletUseCase struct {
	repo repository.WalletRepository
}

func NewWalletUseCase(repo repository.WalletRepository) WalletUseCase {
	return &walletUseCase{
		repo: repo,
	}
}

func (w *walletUseCase) GetWalletByUserId(userId string) (model.Wallet, error) {
	wallet, err := w.repo.FindByUserId(userId)
	if err != nil {
		return model.Wallet{}, err
	}
	return wallet, nil
}

func (w *walletUseCase) CreateWallet(payload model.Wallet) error {
	err := w.repo.Save(payload)
	if err != nil {
		return err
	}

	return nil
}

func (w *walletUseCase) UpdateWalletBalance(walletID string, amount int) error {
	err := w.repo.UpdateWalletBalance(walletID, amount)
	if err != nil {
		return err
	}

	return nil
}
