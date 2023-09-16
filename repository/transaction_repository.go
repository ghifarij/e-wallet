package repository

import (
	"Kelompok-2/dompet-online/model"
	"Kelompok-2/dompet-online/model/dto/resp"
	"database/sql"
)

type TransactionRepository interface {
	FindAll(userId string) ([]resp.GetTransactionsResponse, error)
	Save(transaction model.Transactions) error
	Count(userId string) (int, error)
	FindWalletByUserID(userID string) (model.Wallet, error)
	UpdateWalletBalance(wallet model.Wallet) error
	FindWalletByRekening(rekening string) (model.Wallet, error)
	CreateTransaction(transaction model.Transactions) (model.Transactions, error)
}

type transactionRepository struct {
	db *sql.DB
}

func NewTransactionRepository(db *sql.DB) TransactionRepository {
	return &transactionRepository{
		db: db,
	}
}

func (t *transactionRepository) Save(transaction model.Transactions) error {
	tx, err := t.db.Begin()
	if err != nil {
		return err
	}
	_, err = tx.Exec(`INSERT INTO transaction(id, source_of_funds_id, user_id, destination, amount, description, created_at) VALUES ($1, $2, $3, $4, $5, $6, $7)`,
		transaction.Id,
		transaction.SourceOfFoundId,
		transaction.UserId,
		transaction.Destination,
		transaction.Amount,
		transaction.Description,
		transaction.CreateAt,
	)
	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}

func (t *transactionRepository) FindAll(userId string) ([]resp.GetTransactionsResponse, error) {
	var getTransactionsResponses []resp.GetTransactionsResponse

	query := `
        SELECT
            t.id,
            s.id,
            s.name AS source_of_funds,
            u.id AS user_id,
            u.full_name,
            t.destination,
            w.rekening_user,
            w.balance,
            t.amount,
            t.description,
            t.created_at
        FROM transactions AS t 
        INNER JOIN users AS u ON t.user_id = u.id 
        INNER JOIN wallets AS w ON t.user_id = w.user_id 
        INNER JOIN source_of_funds AS s ON t.source_of_funds_id = s.id;`
	rows, err := t.db.Query(query)
	if err != nil {
		return []resp.GetTransactionsResponse{}, err
	}

	for rows.Next() {
		var getTransactionResponse resp.GetTransactionsResponse

		err := rows.Scan(
			&getTransactionResponse.Id,
			&getTransactionResponse.SourceOfFound.Id,
			&getTransactionResponse.SourceOfFound.Name,
			&getTransactionResponse.User.UserId,
			&getTransactionResponse.User.FullName,
			&getTransactionResponse.DestinationId,
			&getTransactionResponse.Wallet.RekeningUser,
			&getTransactionResponse.Wallet.Balance,
			&getTransactionResponse.Amount,
			&getTransactionResponse.Description,
			&getTransactionResponse.CreateAt,
		)
		if err != nil {
			return nil, err
		}

		getTransactionsResponses = append(getTransactionsResponses, getTransactionResponse)
	}

	// Check if no results were found and return an empty slice
	if len(getTransactionsResponses) == 0 {
		return []resp.GetTransactionsResponse{}, nil
	}

	return getTransactionsResponses, nil
}

func (t *transactionRepository) Count(userId string) (int, error) {
	rows, err := t.db.Query("SELECT COUNT(*) FROM transactions")
	if err != nil {
		return 0, err
	}

	var count int
	for rows.Next() {
		if err := rows.Scan(&count); err != nil {
			return 0, err
		}
	}

	return count, nil
}

func (r *transactionRepository) FindWalletByUserID(userID string) (model.Wallet, error) {
	query := "SELECT id, user_id, rekening_user, balance FROM wallets WHERE user_id = $1"
	var wallet model.Wallet

	err := r.db.QueryRow(query, userID).Scan(&wallet.Id, &wallet.UserId, &wallet.RekeningUser, &wallet.Balance)
	if err != nil {
		return model.Wallet{}, err
	}

	return wallet, nil
}

func (r *transactionRepository) UpdateWalletBalance(wallet model.Wallet) error {
	query := "UPDATE wallets SET balance = $1 WHERE id = $2"

	_, err := r.db.Exec(query, wallet.Balance, wallet.Id)
	if err != nil {
		return err
	}

	return nil
}

func (r *transactionRepository) FindWalletByRekening(rekening string) (model.Wallet, error) {
	query := "SELECT id, user_id, rekening_user, balance FROM wallets WHERE rekening_user = $1"
	var wallet model.Wallet

	err := r.db.QueryRow(query, rekening).Scan(&wallet.Id, &wallet.UserId, &wallet.RekeningUser, &wallet.Balance)
	if err != nil {
		return model.Wallet{}, err
	}

	return wallet, nil
}

func (r *transactionRepository) CreateTransaction(transaction model.Transactions) (model.Transactions, error) {
	query := `
        INSERT INTO transactions (id, source_of_funds_id, user_id, destination, amount, description, created_at)
        VALUES ($1, $2, $3, $4, $5, $6, $7)
    `
	_, err := r.db.Exec(
		query,
		transaction.Id,
		transaction.SourceOfFoundId,
		transaction.UserId,
		transaction.Destination,
		transaction.Amount,
		transaction.Description,
		transaction.CreateAt,
	)
	if err != nil {
		return model.Transactions{}, err
	}

	// Return the created transaction
	return transaction, nil
}
