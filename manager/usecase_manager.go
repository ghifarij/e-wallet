package manager

import "Kelompok-2/dompet-online/usecase"

type UseCaseManager interface {
	UserUseCase() usecase.UserUseCase
	WalletUseCase() usecase.WalletUseCase
	TransactionUseCase() usecase.TransactionUseCase
}

type useCaseManager struct {
	repoManager RepoManager
}

func NewUseCaseManager(repoManager RepoManager) UseCaseManager {
	return &useCaseManager{
		repoManager: repoManager,
	}
}

func (u *useCaseManager) UserUseCase() usecase.UserUseCase {
	return usecase.NewUserUseCase(u.repoManager.UserRepo(), u.WalletUseCase())
}

func (u *useCaseManager) WalletUseCase() usecase.WalletUseCase {
	return usecase.NewWalletUseCase(u.repoManager.WalletRepo())
}

func (u *useCaseManager) TransactionUseCase() usecase.TransactionUseCase {
	return usecase.NewTransactionUseCase(u.repoManager.TransactionRepo(), u.WalletUseCase())
}
