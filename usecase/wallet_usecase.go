package usecase

import (
	"Kelompok-2/dompet-online/model"
	"Kelompok-2/dompet-online/repository"
)

type WalletUseCase interface {
	GetWalletByUserId(userId string) (model.Wallet, error)
	GetWalletByRekeningUser(number string) (model.Wallet, error)
	CreateWallet(payload model.Wallet) error
}

type walletUseCase struct {
	walletRepository repository.WalletRepository
}

func NewWalletUseCase(walletRepository repository.WalletRepository) WalletUseCase {
	return &walletUseCase{
		walletRepository: walletRepository,
	}
}

func (w *walletUseCase) GetWalletByUserId(userId string) (model.Wallet, error) {
	wallet, err := w.walletRepository.FindByUserId(userId)
	if err != nil {
		return model.Wallet{}, err
	}
	return wallet, nil
}

func (w *walletUseCase) GetWalletByRekeningUser(number string) (model.Wallet, error) {
	wallet, err := w.walletRepository.FindByRekeningUser(number)
	if err != nil {
		return model.Wallet{}, err
	}
	return wallet, nil
}

func (w *walletUseCase) CreateWallet(payload model.Wallet) error {
	err := w.walletRepository.Save(payload)
	if err != nil {
		return err
	}

	return nil
}
