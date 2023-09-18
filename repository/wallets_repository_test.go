package repository

import (
	"Kelompok-2/dompet-online/model"
	"database/sql"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"regexp"
	"testing"
	"time"
)

type WalletRepoTestSuite struct {
	suite.Suite
	mockDB  *sql.DB
	mockSQL sqlmock.Sqlmock
	repo    WalletRepository
}

func (suite *WalletRepoTestSuite) SetupTest() {
	db, mock, err := sqlmock.New()
	assert.NoError(suite.T(), err)
	suite.mockDB = db
	suite.mockSQL = mock
	suite.repo = NewWalletRepository(suite.mockDB)
}

func TestWalletRepoTestSuite(t *testing.T) {
	suite.Run(t, new(WalletRepoTestSuite))
}

func (suite *WalletRepoTestSuite) TestFindByUserId_Success() {
	mockData := model.Wallet{
		Id:           "1",
		UserId:       "1",
		RekeningUser: "1234",
		Balance:      10000,
	}
	rows := sqlmock.NewRows([]string{"id", "user_id", "rekening_user", "balance"})
	rows.AddRow(mockData.Id, mockData.UserId, mockData.RekeningUser, mockData.Balance)

	expectedQuery := "SELECT id, user_id, rekening_user, balance FROM wallets WHERE user_id = $1"
	suite.mockSQL.ExpectQuery(regexp.QuoteMeta(expectedQuery)).WithArgs(mockData.UserId).WillReturnRows(rows)
	gotWallet, gotError := suite.repo.FindByUserId(mockData.Id)
	assert.Nil(suite.T(), gotError)
	assert.NoError(suite.T(), gotError)
	assert.Equal(suite.T(), mockData, gotWallet)
}

func (suite *WalletRepoTestSuite) TestFindByUserId_Fail() {
	expectedQuery := "SELECT id, user_id, rekening_user, balance FROM wallets WHERE user_id = $1"
	suite.mockSQL.ExpectQuery(regexp.QuoteMeta(expectedQuery)).WithArgs("xx").WillReturnError(errors.New("error"))
	gotWallet, gotError := suite.repo.FindByUserId("xx")
	assert.Error(suite.T(), gotError)
	assert.Equal(suite.T(), model.Wallet{}, gotWallet)
}

//func (suite *WalletRepoTestSuite) TestFindByRekeningUser_Success() {
//	mockData := model.Wallet{
//		Id:           "1",
//		UserId:       "1",
//		RekeningUser: "1234",
//		Balance:      10000,
//	}
//	rows := sqlmock.NewRows([]string{"id", "user_id", "rekening_user", "balance"})
//	rows.AddRow(mockData.Id, mockData.UserId, mockData.RekeningUser, mockData.Balance)
//	expectedQuery := `SELECT id, user_id, rekening_user, balance FROM wallets WHERE rekening_user = $1`
//	suite.mockSQL.ExpectQuery(regexp.QuoteMeta(expectedQuery)).WithArgs(mockData.RekeningUser).WillReturnRows(rows)
//	gotWallet, gotError := suite.repo.FindByRekeningUser(mockData.RekeningUser)
//	assert.Nil(suite.T(), gotError)
//	assert.NoError(suite.T(), gotError)
//	assert.Equal(suite.T(), mockData, gotWallet)
//}

//func (suite *WalletRepoTestSuite) TestFindByRekeningUser_Fail() {
//	expectedQuery := `SELECT id, user_id, rekening_user, balance FROM wallets WHERE rekening_user = $1`
//	suite.mockSQL.ExpectQuery(regexp.QuoteMeta(expectedQuery)).WithArgs("xx").WillReturnError(errors.New("error"))
//	gotWallet, gotError := suite.repo.FindByRekeningUser("xx")
//	assert.Error(suite.T(), gotError)
//	assert.Equal(suite.T(), model.Wallet{}, gotWallet)
//}

func (suite *WalletRepoTestSuite) TestSave_Success() {
	mockData := model.Wallet{
		Id:           "1",
		UserId:       "1",
		RekeningUser: "1234",
		Balance:      10000,
		CreatedAt:    time.Time{},
		UpdatedAt:    time.Time{},
	}
	expectQuery := `INSERT INTO wallets`
	suite.mockSQL.ExpectExec(expectQuery).WithArgs(
		mockData.Id,
		mockData.UserId,
		mockData.RekeningUser,
		mockData.Balance,
		mockData.CreatedAt,
		mockData.UpdatedAt,
	).WillReturnResult(sqlmock.NewResult(1, 1))

	got := suite.repo.Save(mockData)
	assert.Nil(suite.T(), got)
	assert.NoError(suite.T(), got)
}

func (suite *WalletRepoTestSuite) TestSave_fail() {
	mockData := model.Wallet{
		Id:           "1",
		UserId:       "1",
		RekeningUser: "1234",
		Balance:      10000,
		CreatedAt:    time.Time{},
		UpdatedAt:    time.Time{},
	}
	expectQuery := `INSERT INTO wallets`
	suite.mockSQL.ExpectExec(expectQuery).WithArgs(
		mockData.Id,
		mockData.UserId,
		mockData.RekeningUser,
		mockData.Balance,
		mockData.CreatedAt,
		mockData.UpdatedAt,
	).WillReturnError(errors.New("error"))
	got := suite.repo.Save(mockData)
	assert.Error(suite.T(), got)
}

//func (suite *WalletRepoTestSuite) TestUpdateWalletBalance_success() {
//	mockData := model.Wallet{
//		Id:      "1",
//		Balance: 0,
//	}
//	expectQuery := `UPDATE wallets
//        SET balance = balance + $1
//        WHERE id = $2`
//	suite.mockSQL.ExpectExec(expectQuery).WithArgs(
//		mockData.Id,
//		mockData.Balance).WillReturnResult(sqlmock.NewResult(1, 1))
//	got := suite.repo.UpdateWalletBalance(mockData.Id, mockData.Balance)
//	assert.Nil(suite.T(), got)
//	assert.NoError(suite.T(), got)
//}
