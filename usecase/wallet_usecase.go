package usecase

import (
	"Kelompok-2/dompet-online/model"
	"Kelompok-2/dompet-online/model/dto/req"
	"Kelompok-2/dompet-online/repository"
	"Kelompok-2/dompet-online/util/common"
	"fmt"
	"github.com/go-playground/validator/v10"
)

type WalletUseCase interface {
	GetWalletByUserId(payload req.WalletRequestBody) (model.Wallet, error)
	GetWalletByRekeningUser(number string) (model.Wallet, error)
	CreateWallet(payload req.WalletRequestBody) error
	FindByUserId(id string) (model.Users, error)
}

type walletUseCase struct {
	userUC           UserUseCase
	walletRepository repository.WalletRepository
}

func (w *walletUseCase) FindByUserId(id string) (model.Users, error) {

	byId, err := w.userUC.FindById(id)
	if err != nil {
		return model.Users{}, fmt.Errorf("user not found")
	}
	return byId, nil
}

func (w *walletUseCase) GetWalletByUserId(payload req.WalletRequestBody) (model.Wallet, error) {
	wallet, err := w.walletRepository.FindByUserId(payload.UserId)
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

func (w *walletUseCase) CreateWallet(payload req.WalletRequestBody) error {
	// Validate the payload
	validate := validator.New()
	err := validate.Struct(payload)
	if err != nil {
		return err
	}
	user, err := w.userUC.FindById(payload.UserId)
	if err != nil {
		return fmt.Errorf("id di user")
	}

	wallet, err := w.walletRepository.FindByUserId(user.Id)
	if err != nil {
		return fmt.Errorf("user_id di wallet")
	}

	wallet.UserId = user.Id
	wallet.RekeningUser = common.GenerateRandomRekeningNumber(user.Id)
	wallet.Balance = 0

	err = w.walletRepository.Save(wallet)
	if err != nil {
		return err
	}

	return nil
}

func NewWalletUseCase(userUC UserUseCase, walletRepository repository.WalletRepository) WalletUseCase {
	return &walletUseCase{
		userUC:           userUC,
		walletRepository: walletRepository,
	}
}
