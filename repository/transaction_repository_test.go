package repository

import (
	"Kelompok-2/dompet-online/model"
	"database/sql"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

type TransactionRepoTestSuite struct {
	suite.Suite
	mockDb  *sql.DB
	mockSQL sqlmock.Sqlmock
	repo    TransactionRepository
}

func (suite *TransactionRepoTestSuite) SetupTest() {
	db, mock, err := sqlmock.New()
	assert.NoError(suite.T(), err)
	suite.mockDb = db
	suite.mockSQL = mock
	suite.repo = NewTransactionRepository(suite.mockDb)
}

func TestTransactionRepoTestSuite(t *testing.T) {
	suite.Run(t, new(TransactionRepoTestSuite))
}

func (suite *TransactionRepoTestSuite) TestCreateTransaction_Success() {
	mockData := model.Transactions{
		Id:              "1",
		UserId:          "1",
		SourceWalletID:  "001",
		Destination:     "pln",
		Amount:          50000,
		Description:     "tagihan listrik",
		PaymentMethodID: "dana",
		CreateAt:        time.Time{},
	}
	rows := sqlmock.NewRows([]string{"id", "user_id", "source_wallet_id", "destination", "amount", "description", "payment_method_id", "created_at"})
	rows.AddRow(mockData.Id, mockData.UserId, mockData.SourceWalletID, mockData.Destination, mockData.Amount, mockData.Description, mockData.PaymentMethodID, mockData.CreateAt)

	suite.mockSQL.ExpectExec("INSERT INTO transactions").
		WithArgs(mockData.Id, mockData.UserId, mockData.SourceWalletID, mockData.Destination, mockData.Amount, mockData.Description, mockData.PaymentMethodID, mockData.CreateAt).
		WillReturnResult(sqlmock.NewResult(1, 1))
	result, err := suite.repo.CreateTransaction(mockData)
	assert.NoError(suite.T(), err)
	assert.NotNil(suite.T(), result)
}

func (suite *TransactionRepoTestSuite) TestCreateTransaction_Fail() {
	mockData := model.Transactions{
		Id:              "1",
		UserId:          "1",
		SourceWalletID:  "001",
		Destination:     "pln",
		Amount:          50000,
		Description:     "tagihan listrik",
		PaymentMethodID: "dana",
		CreateAt:        time.Time{},
	}

	// TODO create transaction failed
	rows := sqlmock.NewRows([]string{"id", "source_wallet_ID", "user_Id", "Payment_Method_id", "destination", "amount", "description", "created_at"})
	rows.AddRow(mockData.Id, mockData.UserId, mockData.SourceWalletID, mockData.Destination, mockData.Amount, mockData.Description, mockData.PaymentMethodID, mockData.CreateAt)
	suite.mockSQL.ExpectExec("INSERT INTO transactions").
		WithArgs(mockData.Id, mockData.UserId, mockData.SourceWalletID, mockData.Destination, mockData.Amount, mockData.Description, mockData.PaymentMethodID, mockData.CreateAt).
		WillReturnError(errors.New("Insertion failed"))
	result, err := suite.repo.CreateTransaction(mockData)
	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), result)
	assert.NoError(suite.T(), suite.mockSQL.ExpectationsWereMet())
}

func (suite *TransactionRepoTestSuite) TestFindAll_Success() {
	userId := "1"

	// Define the expected rows returned by the SQL query
	expectedRows := sqlmock.NewRows([]string{
		"transaction_id",
		"destination",
		"amount",
		"transaction_description",
		"transaction_created_at",
		"user_name",
		"rekening_user",
		"balance",
		"payment_method_name",
		"payment_method_description",
	}).AddRow(
		"1",
		"pln",
		50000,
		"tagihan listrik",
		time.Now(),
		"John Doe",
		"123456",
		100000,
		"dana",
		"Dana Payment Method",
	)

	// Set up expectations for the SQL mock
	suite.mockSQL.ExpectQuery(`SELECT .* FROM transactions`).WithArgs(userId).WillReturnRows(expectedRows)

	// Call the FindAll method
	result, err := suite.repo.FindAll(userId)
	assert.NoError(suite.T(), err)
	assert.NotNil(suite.T(), result)
	assert.Equal(suite.T(), 1, len(result))

	// Validate the first returned result
	assert.Equal(suite.T(), "1", result[0].Id)
	assert.Equal(suite.T(), "pln", result[0].Destination)
	assert.Equal(suite.T(), 50000, result[0].Amount)
	assert.Equal(suite.T(), "tagihan listrik", result[0].Description)
	assert.NotNil(suite.T(), result[0].CreateAt)
	assert.Equal(suite.T(), "John Doe", result[0].User.UserName)
	assert.Equal(suite.T(), "123456", result[0].Wallet.RekeningUser)
	assert.Equal(suite.T(), 100000, result[0].Wallet.Balance)
	assert.Equal(suite.T(), "dana", result[0].PaymentMethod.Name)
	assert.Equal(suite.T(), "Dana Payment Method", result[0].PaymentMethod.Description)

	// Ensure all expectations were met
	assert.NoError(suite.T(), suite.mockSQL.ExpectationsWereMet())
}

func (suite *TransactionRepoTestSuite) TestFindAll_Fail() {
	userId := "1"

	// Set up expectations for the SQL mock to return an error
	suite.mockSQL.ExpectQuery(`SELECT .* FROM transactions`).WithArgs(userId).WillReturnError(errors.New("Database error"))

	// Call the FindAll method
	result, err := suite.repo.FindAll(userId)

	// Assert that an error occurred
	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), result)

	// Ensure all expectations were met
	assert.NoError(suite.T(), suite.mockSQL.ExpectationsWereMet())
}

func (suite *TransactionRepoTestSuite) TestCount_Success() {
	mockData := []model.Transactions{
		{
			Id:              "1",
			UserId:          "1",
			SourceWalletID:  "001",
			Destination:     "pln",
			Amount:          50000,
			Description:     "tagihan listrik",
			PaymentMethodID: "dana",
			CreateAt:        time.Time{},
		},
		{
			Id:              "2",
			UserId:          "2",
			SourceWalletID:  "002",
			Destination:     "pln",
			Amount:          75000,
			Description:     "tagihan listrik",
			PaymentMethodID: "dana",
			CreateAt:        time.Time{},
		},
	}
	suite.mockSQL.ExpectQuery(`SELECT COUNT(id) FROM transactions WHERE user_id = $1`).
		WithArgs("1").
		WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(len(mockData)))
	count, err := suite.repo.Count("1")
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), len(mockData), count)
}

func (suite *TransactionRepoTestSuite) TestCount_Fail() {
	// Set up expectations for the count query to return an error
	suite.mockSQL.ExpectQuery(`SELECT COUNT(id) FROM transactions WHERE user_id = $1`).
		WithArgs("1").
		WillReturnError(errors.New("Count query failed"))

	count, err := suite.repo.Count("1")
	assert.Error(suite.T(), err)
	assert.Zero(suite.T(), count) // Assuming the count should be zero in case of failure
}
