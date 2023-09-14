package manager

import "Kelompok-2/dompet-online/repository"

type RepoManager interface {
	UserRepo() repository.UserRepository
	WalletRepo() repository.WalletRepository
	TransactionRepo() repository.TransactionRepository
}

type repoManager struct {
	infraManager InfraManager
}

func NewRepoManager(infraManager InfraManager) RepoManager {
	return &repoManager{infraManager: infraManager}
}

func (r *repoManager) UserRepo() repository.UserRepository {
	return repository.NewUserRepository(r.infraManager.Conn())
}

func (r *repoManager) WalletRepo() repository.WalletRepository {
	return repository.NewWalletRepository(r.infraManager.Conn())
}

func (r *repoManager) TransactionRepo() repository.TransactionRepository {
	return repository.NewTransactionRepository(r.infraManager.Conn())
}
