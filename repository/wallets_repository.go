package repository

import (
	"Kelompok-2/dompet-online/model"
	"database/sql"
)

type WalletRepository interface {
	FindByUserId(userid string) (model.Wallet, error)
	FindByRekeningUser(number string) (model.Wallet, error)
	Save(wallet model.Wallet) error
}

type walletRepository struct {
	db *sql.DB
}

func (w *walletRepository) FindByUserId(userid string) (model.Wallet, error) {
	row := w.db.QueryRow("SELECT id, user_id, rekening_user, balance FROM wallets WHERE user_id = $1", userid)
	var wallet model.Wallet
	err := row.Scan(
		&wallet.Id,
		&wallet.UserId,
		&wallet.RekeningUser,
		&wallet.Balance,
	)
	if err != nil {
		return model.Wallet{}, err
	}
	return wallet, nil
}

func (w *walletRepository) FindByRekeningUser(number string) (model.Wallet, error) {
	var wallet model.Wallet
	err := w.db.QueryRow(`SELECT id, user_id, rekening_user, balance FROM wallets WHERE rekening_user = $1`, number).Scan(
		&wallet.Id,
		&wallet.UserId,
		&wallet.RekeningUser,
		&wallet.Balance,
	)
	if err != nil {
		return wallet, err
	}
	return wallet, nil
}

func (w *walletRepository) Save(wallet model.Wallet) error {
	_, err := w.db.Exec(`INSERT INTO wallets (id, user_id, rekening_user, balance, created_at, updated_at) VALUES($1, $2, $3, $4, $5, $6)`,
		wallet.Id,
		wallet.UserId,
		wallet.RekeningUser,
		wallet.Balance,
		wallet.CreatedAt,
		wallet.UpdatedAt,
	)
	if err != nil {
		return err
	}
	return nil
}

func NewWalletRepository(db *sql.DB) WalletRepository {
	return &walletRepository{
		db: db,
	}
}
