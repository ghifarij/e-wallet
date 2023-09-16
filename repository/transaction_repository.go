package repository

import (
	"Kelompok-2/dompet-online/model"
	"Kelompok-2/dompet-online/model/dto/resp"
	"database/sql"
)

type TransactionRepository interface {
	FindAll(userId string) ([]resp.GetTransactionsResponse, error)
	CreateTransaction(transaction model.Transactions) (model.Transactions, error)
	Count(userId string) (int, error)
}

type transactionRepository struct {
	db *sql.DB
}

func NewTransactionRepository(db *sql.DB) TransactionRepository {
	return &transactionRepository{
		db: db,
	}
}

func (t *transactionRepository) FindAll(userId string) ([]resp.GetTransactionsResponse, error) {
	var getTransactionsResponses []resp.GetTransactionsResponse

	query := `
        SELECT
            t.id,
            t.destination,
            t.amount,
            t.description,
            t.created_at,
            u.user_name,
            w.rekening_user,
            w.balance,
            p.name,
            p.description
        FROM transactions AS t 
        JOIN users AS u ON t.user_id = u.id
		JOIN wallets AS w ON t.source_wallet_id = w.id
		JOIN payment_method AS p ON t.payment_method_id = p.id;`
	rows, err := t.db.Query(query)
	if err != nil {
		return []resp.GetTransactionsResponse{}, err
	}

	for rows.Next() {
		var getTransactionResponse resp.GetTransactionsResponse

		err := rows.Scan(
			&getTransactionResponse.Id,
			&getTransactionResponse.Destination,
			&getTransactionResponse.Amount,
			&getTransactionResponse.Description,
			&getTransactionResponse.CreateAt,
			&getTransactionResponse.User.UserName,
			&getTransactionResponse.Wallet.RekeningUser,
			&getTransactionResponse.Wallet.Balance,
			&getTransactionResponse.PaymentMethod.Name,
			&getTransactionResponse.PaymentMethod.Description,
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

func (r *transactionRepository) CreateTransaction(transaction model.Transactions) (model.Transactions, error) {
	query := `
        INSERT INTO transactions (id, user_id, source_wallet_id, destination, amount, description, payment_method_id, created_at)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
        RETURNING id, user_id, source_wallet_id, destination, amount, description, payment_method_id, created_at
    `

	var createdTransaction model.Transactions
	err := r.db.QueryRow(
		query,
		transaction.Id,
		transaction.UserId,
		transaction.SourceWalletID,
		transaction.Destination,
		transaction.Amount,
		transaction.Description,
		transaction.PaymentMethodID,
		transaction.CreateAt,
	).Scan(
		&createdTransaction.Id,
		&createdTransaction.UserId,
		&createdTransaction.SourceWalletID,
		&createdTransaction.Destination,
		&createdTransaction.Amount,
		&createdTransaction.Description,
		&createdTransaction.PaymentMethodID,
		&createdTransaction.CreateAt,
	)

	if err != nil {
		return model.Transactions{}, err
	}

	return createdTransaction, nil
}

func (t *transactionRepository) Count(userId string) (int, error) {
	query := `
        SELECT COUNT(*) FROM transactions WHERE user_id = $1
    `

	var count int
	err := t.db.QueryRow(query, userId).Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}
