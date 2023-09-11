package manager

import "Kelompok-2/dompet-online/usecase"

type UseCaseManager interface {
	UserUseCase() usecase.UserUseCase
	AuthUseCase() usecase.AuthUseCase
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

func (u *useCaseManager) AuthUseCase() usecase.AuthUseCase {
	return usecase.NewAuthUseCase(u.repoManager.UserRepo())
}