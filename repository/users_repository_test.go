package repository

import (
	"Kelompok-2/dompet-online/model"
	"Kelompok-2/dompet-online/model/dto/req"
	"database/sql"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"regexp"
	"testing"
	"time"
)

type UserRepoTestSuite struct {
	suite.Suite
	mockDB  *sql.DB
	mockSQL sqlmock.Sqlmock
	repo    UserRepository
}

func (suite *UserRepoTestSuite) SetupTest() {
	db, mock, err := sqlmock.New()
	assert.NoError(suite.T(), err)
	suite.mockDB = db
	suite.mockSQL = mock
	suite.repo = NewUserRepository(suite.mockDB)
}

func TestUserRepoTestSuite(t *testing.T) {
	suite.Run(t, new(UserRepoTestSuite))
}

func (suite *UserRepoTestSuite) TestSave_Success() {
	mockData := model.Users{
		Id:              "1",
		FullName:        "Ags",
		Email:           "agus@mail.com",
		PhoneNumber:     "081256789102",
		UserName:        "agus",
		Password:        "password",
		PasswordConfirm: "password",
		IsActive:        true,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
		DisableAt:       time.Now(),
	}
	expectQuery := `INSERT INTO users`
	suite.mockSQL.ExpectExec(expectQuery).WithArgs(
		mockData.Id,
		mockData.FullName,
		mockData.UserName,
		mockData.Email,
		mockData.PhoneNumber,
		mockData.Password,
		mockData.PasswordConfirm,
		mockData.IsActive,
		mockData.CreatedAt,
		mockData.UpdatedAt,
		mockData.DisableAt).WillReturnResult(sqlmock.NewResult(1, 1))

	got := suite.repo.Save(mockData)
	assert.Nil(suite.T(), got)
	assert.NoError(suite.T(), got)
}

func (suite *UserRepoTestSuite) TestSave_Fail() {
	mockData := model.Users{
		Id:              "1",
		FullName:        "Ags",
		Email:           "agus@mail.com",
		PhoneNumber:     "081256789102",
		UserName:        "agus",
		Password:        "password",
		PasswordConfirm: "password",
		IsActive:        true,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
		DisableAt:       time.Now(),
	}
	expectQuery := `INSERT INTO users`
	suite.mockSQL.ExpectExec(expectQuery).WithArgs(
		mockData.Id,
		mockData.FullName,
		mockData.UserName,
		mockData.Email,
		mockData.PhoneNumber,
		mockData.Password,
		mockData.PasswordConfirm,
		mockData.IsActive,
		mockData.CreatedAt,
		mockData.UpdatedAt,
		mockData.DisableAt).WillReturnError(errors.New("error"))

	got := suite.repo.Save(mockData)
	assert.Error(suite.T(), got)
}

func (suite *UserRepoTestSuite) TestFindByUserName_Success() {
	mockData := model.Users{
		Id:       "1",
		UserName: "agus",
		Password: "password",
	}
	rows := sqlmock.NewRows([]string{"id", "user_name", "password"})
	rows.AddRow(mockData.Id, mockData.UserName, mockData.Password)

	expectedQuery := `SELECT id, user_name, password FROM users WHERE user_name = $1`
	suite.mockSQL.ExpectQuery(regexp.QuoteMeta(expectedQuery)).WithArgs(mockData.UserName).WillReturnRows(rows)
	gotUser, gotError := suite.repo.FindByUserName(mockData.UserName)
	assert.Nil(suite.T(), gotError)
	assert.NoError(suite.T(), gotError)
	assert.Equal(suite.T(), mockData, gotUser)
}

func (suite *UserRepoTestSuite) TestFindByUserName_Fail() {
	expectedQuery := `SELECT id, user_name, password FROM users WHERE user_name = $1`
	suite.mockSQL.ExpectQuery(regexp.QuoteMeta(expectedQuery)).WithArgs("XX").WillReturnError(errors.New("error"))
	gotUser, gotError := suite.repo.FindByUserName("XX")
	assert.Error(suite.T(), gotError)
	assert.Equal(suite.T(), model.Users{}, gotUser)
}

func (suite *UserRepoTestSuite) TestFindById_Success() {
	mockData := model.Users{
		Id: "1",
	}
	rows := sqlmock.NewRows([]string{"id"})
	rows.AddRow(mockData.Id)

	expectedQuery := `SELECT id FROM users WHERE id = $1`
	suite.mockSQL.ExpectQuery(regexp.QuoteMeta(expectedQuery)).WithArgs(mockData.Id).WillReturnRows(rows)
	gotUser, gotError := suite.repo.FindById(mockData.Id)
	assert.Nil(suite.T(), gotError)
	assert.NoError(suite.T(), gotError)
	assert.Equal(suite.T(), mockData, gotUser)
}

func (suite *UserRepoTestSuite) TestFindById_Fail() {
	expectedQuery := `SELECT id FROM users WHERE id = $1`
	suite.mockSQL.ExpectQuery(regexp.QuoteMeta(expectedQuery)).WithArgs("XX").WillReturnError(errors.New("error"))
	gotUser, gotError := suite.repo.FindById("XX")
	assert.Error(suite.T(), gotError)
	assert.Equal(suite.T(), model.Users{}, gotUser)
}

func (suite *UserRepoTestSuite) TestFindByPhoneNumber_Success() {
	mockData := model.Users{
		Id:              "1",
		FullName:        "Ags",
		Email:           "agus@mail.com",
		PhoneNumber:     "081256789102",
		UserName:        "agus",
		Password:        "password",
		PasswordConfirm: "password",
		IsActive:        true,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
		DisableAt:       time.Now(),
	}
	rows := sqlmock.NewRows([]string{"id", "full_name", "user_name", "email", "phone_number", "password", "password_confirm", "is_active", "created_at", "updated_at", "disable_at"})
	rows.AddRow(mockData.Id, mockData.FullName, mockData.UserName, mockData.Email, mockData.PhoneNumber, mockData.Password, mockData.PasswordConfirm, mockData.IsActive, mockData.CreatedAt, mockData.UpdatedAt, mockData.DisableAt)

	expectedQuery := `SELECT id, full_name, user_name, email, phone_number, password, password_confirm, is_active, created_at, updated_at, disable_at FROM users WHERE phone_number = $1`
	suite.mockSQL.ExpectQuery(regexp.QuoteMeta(expectedQuery)).WithArgs(mockData.PhoneNumber).WillReturnRows(rows)
	gotUser, gotError := suite.repo.FindByPhoneNumber(mockData.PhoneNumber)
	assert.Nil(suite.T(), gotError)
	assert.NoError(suite.T(), gotError)
	assert.Equal(suite.T(), mockData, gotUser)
}

func (suite *UserRepoTestSuite) TestFindPhoneNumber_Fail() {
	expectedQuery := `SELECT id, full_name, user_name, email, phone_number, password, password_confirm, is_active, created_at, updated_at, disable_at FROM users WHERE phone_number = $1`
	suite.mockSQL.ExpectQuery(regexp.QuoteMeta(expectedQuery)).WithArgs("XX").WillReturnError(errors.New("error"))
	gotUser, gotError := suite.repo.FindByPhoneNumber("XX")
	assert.Error(suite.T(), gotError)
	assert.Equal(suite.T(), model.Users{}, gotUser)
}

func (suite *UserRepoTestSuite) TestUpdatePassword_Success() {
	username := "agus"
	newPassword := "new_password"
	newPasswordConfirm := "new_password_confirm"

	expectQuery := `UPDATE users SET password = ?, password_confirm = ? WHERE user_name = ?`
	suite.mockSQL.ExpectExec(expectQuery).WithArgs(newPassword, newPasswordConfirm, username).WillReturnResult(sqlmock.NewResult(1, 1))

	err := suite.repo.UpdatePassword(username, newPassword, newPasswordConfirm)
	assert.Nil(suite.T(), err) // Check if error is nil
}

func (suite *UserRepoTestSuite) TestUpdatePassword_Fail() {
	username := "agus"
	newPassword := "new_password"
	newPasswordConfirm := "new_password_confirm"

	expectQuery := `UPDATE users SET password = ?, password_confirm = ? WHERE user_name = ?`
	suite.mockSQL.ExpectExec(expectQuery).WithArgs(newPassword, newPasswordConfirm, username).WillReturnError(errors.New("error"))

	err := suite.repo.UpdatePassword(username, newPassword, newPasswordConfirm)
	assert.Error(suite.T(), err) // Check if an error occurred
}

func (suite *UserRepoTestSuite) TestUpdateAccount_Success() {
	payload := req.UpdateAccountRequest{
		Id:          "1",
		FullName:    "Updated Name",
		Username:    "updated_username",
		Email:       "updated_email@example.com",
		PhoneNumber: "updated_phone_number",
	}

	expectQuery := `UPDATE users SET full_name = ?, user_name = ?, email = ?, phone_number = ? WHERE id = ?`
	suite.mockSQL.ExpectExec(expectQuery).WithArgs(
		payload.FullName,
		payload.Username,
		payload.Email,
		payload.PhoneNumber,
		payload.Id,
	).WillReturnResult(sqlmock.NewResult(1, 1))

	err := suite.repo.UpdateAccount(payload)
	assert.Nil(suite.T(), err) // Check if error is nil
}

func (suite *UserRepoTestSuite) TestUpdateAccount_Fail() {
	payload := req.UpdateAccountRequest{
		Id:          "1",
		FullName:    "Updated Name",
		Username:    "updated_username",
		Email:       "updated_email@example.com",
		PhoneNumber: "updated_phone_number",
	}

	expectQuery := `UPDATE users SET full_name = ?, user_name = ?, email = ?, phone_number = ? WHERE id = ?`
	suite.mockSQL.ExpectExec(expectQuery).WithArgs(
		payload.FullName,
		payload.Username,
		payload.Email,
		payload.PhoneNumber,
		payload.Id,
	).WillReturnError(errors.New("error"))

	err := suite.repo.UpdateAccount(payload)
	assert.Error(suite.T(), err) // Check if an error occurred
}

func (suite *UserRepoTestSuite) TestFindAll_Success() {
	mockData := []model.Users{
		{
			Id:              "1",
			FullName:        "Ags",
			Email:           "agus@mail.com",
			PhoneNumber:     "081256789102",
			UserName:        "agus",
			Password:        "password",
			PasswordConfirm: "password",
			IsActive:        true,
			CreatedAt:       time.Now(),
			UpdatedAt:       time.Now(),
			DisableAt:       time.Now(),
		}, {
			Id:              "2",
			FullName:        "Sui",
			Email:           "suuui@mail.com",
			PhoneNumber:     "081256789111",
			UserName:        "suuui",
			Password:        "password",
			PasswordConfirm: "password",
			IsActive:        true,
			CreatedAt:       time.Now(),
			UpdatedAt:       time.Now(),
			DisableAt:       time.Now(),
		},
	}
	rows := sqlmock.NewRows([]string{"id", "full_name", "user_name", "email", "phone_number", "password", "password_confirm", "is_active", "created_at", "updated_at", "disable_at"})
	for _, u := range mockData {
		rows.AddRow(u.Id, u.FullName, u.UserName, u.Email, u.PhoneNumber, u.Password, u.PasswordConfirm, u.IsActive, u.CreatedAt, u.UpdatedAt, u.DisableAt)
	}

	expectedQuery := `SELECT id, full_name, user_name, email, phone_number, password, password_confirm, is_active, created_at, updated_at, disable_at FROM users`
	suite.mockSQL.ExpectQuery(regexp.QuoteMeta(expectedQuery)).WillReturnRows(rows)
	_, gotError := suite.repo.FindAll()
	assert.Nil(suite.T(), gotError)
}

func (suite *UserRepoTestSuite) TestFindAll_Fail() {
	expectedQuery := `SELECT id, full_name, user_name, email, phone_number, password, password_confirm, is_active, created_at, updated_at, disable_at FROM users`
	suite.mockSQL.ExpectQuery(regexp.QuoteMeta(expectedQuery)).WillReturnError(errors.New("error"))
	gotProduct, gotError := suite.repo.FindAll()
	assert.Error(suite.T(), gotError)
	assert.Nil(suite.T(), gotProduct)
}

func (suite *UserRepoTestSuite) TestDisableUserId_Success() {
	userId := "1"
	disableTime := time.Now()

	expectQuery := `UPDATE users SET is_active = false, disable_at = ? WHERE id = ?`
	suite.mockSQL.ExpectExec(expectQuery).WithArgs(disableTime, userId).WillReturnResult(sqlmock.NewResult(1, 1))

	_, err := suite.repo.DisableUserId(userId, disableTime)
	assert.Nil(suite.T(), err) // Check if error is nil
}

func (suite *UserRepoTestSuite) TestDisableUserId_Fail() {
	userId := "1"
	disableTime := time.Now()

	expectQuery := `UPDATE users SET is_active = false, disable_at = ? WHERE id = ?`
	suite.mockSQL.ExpectExec(expectQuery).WithArgs(disableTime, userId).WillReturnError(errors.New("error"))

	_, err := suite.repo.DisableUserId(userId, disableTime)
	assert.Error(suite.T(), err) // Check if an error occurred
}
