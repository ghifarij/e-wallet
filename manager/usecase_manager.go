package manager

import "Kelompok-2/dompet-online/usecase"

type UseCaseManager interface {
	UserUseCase() usecase.UserUseCase
	WalletUseCase() usecase.WalletUseCase
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
	return usecase.NewUserUseCase(u.repoManager.UserRepo())
}

func (u *useCaseManager) WalletUseCase() usecase.WalletUseCase {
	return usecase.NewWalletUseCase(u.UserUseCase(), u.repoManager.WalletRepo())
}
