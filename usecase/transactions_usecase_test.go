package usecase

import (
	"Kelompok-2/dompet-online/mock/repository_mock"
	"Kelompok-2/dompet-online/mock/usecase_mock"
	"Kelompok-2/dompet-online/model"
	"Kelompok-2/dompet-online/model/dto/req"
	"Kelompok-2/dompet-online/model/dto/resp"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"testing"
)

type TransactionUsecaseTestSuite struct {
	suite.Suite
	transactionRepoMock *repository_mock.TransactionRepositoryMock
	walletUsecaseMock   *usecase_mock.WalletUseCaseMock
	transactionUseCase  TransactionUseCase
}

func (suite *TransactionUsecaseTestSuite) SetupTest() {
	suite.transactionRepoMock = new(repository_mock.TransactionRepositoryMock)
	suite.walletUsecaseMock = new(usecase_mock.WalletUseCaseMock)
	suite.transactionUseCase = NewTransactionUseCase(suite.transactionRepoMock, suite.walletUsecaseMock)
}

func TestTransactionUsecaseTestSuite(t *testing.T) {
	suite.Run(t, new(TransactionUsecaseTestSuite))
}

func (suite *TransactionUsecaseTestSuite) TestGetHistoriesTransactions() {
	// Arrange
	userID := "user123"
	expectedTransactions := []resp.GetTransactionsResponse{
		// Define expected response data here
	}
	suite.transactionRepoMock.On("FindAll", userID).Return(expectedTransactions, nil)

	// Act
	transactions, err := suite.transactionUseCase.GetHistoriesTransactions(userID)

	// Assert
	suite.NoError(err)
	suite.Equal(expectedTransactions, transactions)
	suite.transactionRepoMock.AssertExpectations(suite.T())
}

func (suite *TransactionUsecaseTestSuite) TestTopUp() {
	// Arrange
	payload := req.TopUpRequest{
		// Define payload data here
	}
	expectedTransaction := model.Transactions{
		// Define expected transaction data here
	}
	suite.transactionRepoMock.On("CreateTransaction", mock.AnythingOfType("model.Transactions")).Return(expectedTransaction, nil)
	suite.walletUsecaseMock.On("UpdateWalletBalance", payload.WalletID, payload.Amount).Return(nil)

	// Act
	transaction, err := suite.transactionUseCase.TopUp(payload)

	// Assert
	suite.NoError(err)
	suite.Equal(expectedTransaction, transaction)
	suite.transactionRepoMock.AssertExpectations(suite.T())
	suite.walletUsecaseMock.AssertExpectations(suite.T())
}

func (suite *TransactionUsecaseTestSuite) TestTransfer() {
	// Arrange
	payload := req.TransferRequest{
		// Define payload data here
	}
	expectedTransaction := model.Transactions{
		// Define expected transaction data here
	}
	suite.transactionRepoMock.On("CreateTransaction", mock.AnythingOfType("model.Transactions")).Return(expectedTransaction, nil)
	suite.walletUsecaseMock.On("UpdateWalletBalance", payload.SourceWalletID, -payload.Amount).Return(nil)
	suite.walletUsecaseMock.On("UpdateWalletBalance", payload.DestinationWalletID, payload.Amount).Return(nil)

	// Act
	transaction, err := suite.transactionUseCase.Transfer(payload)

	// Assert
	suite.NoError(err)
	suite.Equal(expectedTransaction, transaction)
	suite.transactionRepoMock.AssertExpectations(suite.T())
	suite.walletUsecaseMock.AssertExpectations(suite.T())
}

func (suite *TransactionUsecaseTestSuite) TestCountTransaction() {
	// Arrange
	userID := "user123"
	expectedCount := 5
	suite.transactionRepoMock.On("Count", userID).Return(expectedCount, nil)

	// Act
	count, err := suite.transactionUseCase.CountTransaction(userID)

	// Assert
	suite.NoError(err)
	suite.Equal(expectedCount, count)
	suite.transactionRepoMock.AssertExpectations(suite.T())
}
