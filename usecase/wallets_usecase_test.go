package usecase

import (
	"Kelompok-2/dompet-online/mock/repository_mock"
	"Kelompok-2/dompet-online/model"
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

type WalletUsecaseTestSuite struct {
	suite.Suite
	walletRepoMock *repository_mock.WalletRepositoryMock
	walletUseCase  WalletUseCase
}

func (suite *WalletUsecaseTestSuite) SetupTest() {
	suite.walletRepoMock = new(repository_mock.WalletRepositoryMock)
	suite.walletUseCase = NewWalletUseCase(suite.walletRepoMock)
}

func TestWalletUsecaseTestSuite(t *testing.T) {
	suite.Run(t, new(WalletUsecaseTestSuite))
}

func (suite *WalletUsecaseTestSuite) TestGetWalletByUserId_Success() {
	expected := model.Wallet{
		Id:           "1",
		UserId:       "001",
		RekeningUser: "1234",
		Balance:      10000,
		CreatedAt:    time.Time{},
		UpdatedAt:    time.Time{},
	}

	suite.walletRepoMock.On("FindByUserId", expected.UserId).Return(expected, nil)

	got, err := suite.walletUseCase.GetWalletByUserId(expected.UserId)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), expected, got)
}

func (suite *WalletUsecaseTestSuite) TestGetWalletByUserId_fail() {
	mockUserId := "0"
	expectedError := fmt.Errorf("user with ID %v not found", mockUserId)
	suite.walletRepoMock.On("FindByUserId", mockUserId).Return(model.Wallet{}, expectedError)
	_, err := suite.walletUseCase.GetWalletByUserId(mockUserId)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), expectedError.Error(), err.Error())
}

func (suite *WalletUsecaseTestSuite) TestCreateWallet_Success() {
	mockData := model.Wallet{
		Id:           "1",
		UserId:       "001",
		RekeningUser: "1234",
		Balance:      10000,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Time{},
	}
	suite.walletRepoMock.On("Save", mockData).Return(nil)
	got := suite.walletUseCase.CreateWallet(mockData)
	assert.Nil(suite.T(), got)
	assert.NoError(suite.T(), got)
}

func (suite *WalletUsecaseTestSuite) TestCreateWallet_Failed() {
	mockData := model.Wallet{
		Id:           "1",
		UserId:       "001",
		RekeningUser: "1234",
		Balance:      10000,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Time{},
	}
	suite.walletRepoMock.On("Save", mockData).Return(errors.New("failed"))
	got := suite.walletUseCase.CreateWallet(mockData)
	assert.Error(suite.T(), got)
}

func (suite *WalletUsecaseTestSuite) TestUpdateWalletBalance_Success() {
	walletID := "1"
	amount := 5000

	suite.walletRepoMock.On("UpdateWalletBalance", walletID, amount).Return(nil)
	err := suite.walletUseCase.UpdateWalletBalance(walletID, amount)

	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
}

func (suite *WalletUsecaseTestSuite) TestUpdateWalletBalance_Failed() {
	walletID := "1"
	amount := 5000

	suite.walletRepoMock.On("UpdateWalletBalance", walletID, amount).Return(errors.New("failed"))
	err := suite.walletUseCase.UpdateWalletBalance(walletID, amount)

	assert.Error(suite.T(), err)
}
