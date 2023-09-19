package usecase

import (
	"Kelompok-2/dompet-online/mock/repository_mock"
	"Kelompok-2/dompet-online/mock/usecase_mock"
	"Kelompok-2/dompet-online/model"
	"Kelompok-2/dompet-online/model/dto/req"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

type UserUseCaseTestSuite struct {
	suite.Suite
	userRepoMock   *repository_mock.UserRepositoryMock
	userWalletMock *usecase_mock.WalletUseCaseMock
	userUseCase    UserUseCase
}

func (suite *UserUseCaseTestSuite) SetupTest() {
	suite.userRepoMock = new(repository_mock.UserRepositoryMock)
	suite.userWalletMock = new(usecase_mock.WalletUseCaseMock)
	suite.userUseCase = NewUserUseCase(suite.userRepoMock, suite.userWalletMock)
}

func TestUserUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(UserUseCaseTestSuite))
}

func (suite *UserUseCaseTestSuite) TestLogin_Success() {
	// Prepare a valid login request payload
	validLoginRequest := req.AuthLoginRequest{
		//LoginOption: req.AuthLoginOption{
		//	Email: "valid@example.com", // Provide a valid email or other identifier
		//},
		Password: "validPassword", // Provide a valid password
	}

	// Create a user instance with valid credentials
	validUser := model.Users{
		Email:    "valid@example.com",
		Password: "hashedValidPassword", // Replace with the hashed valid password
		IsActive: true,
		// Other user fields...
	}

	// Mock the UserRepository's FindByUsernameEmailPhoneNumber function to return the valid user
	suite.userRepoMock.On("FindByUsernameEmailPhoneNumber", validLoginRequest.LoginOption.Email).Return(validUser, nil)

	// Mock the GenerateJwtToken function to return a valid token
	suite.userRepoMock.On("GenerateJwtToken", mock.AnythingOfType("model.Users")).Return("validToken", nil)

	// Call the Login function with the valid payload
	response, err := suite.userUseCase.Login(validLoginRequest)

	// Assert that the login was successful
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), "valid@example.com", response.UserName)
	assert.Equal(suite.T(), "validToken", response.Token)

	// Verify that the expectations were met
	suite.userRepoMock.AssertExpectations(suite.T())
}

func (suite *UserUseCaseTestSuite) TestLogin_Failure_InvalidCredential() {
	// Prepare an invalid login request payload
	invalidLoginRequest := req.AuthLoginRequest{
		//LoginOption: req.AuthLoginOption{
		//	Email: "invalid@example.com", // Provide an invalid email or other identifier
		//},
		Password: "invalidPassword", // Provide an invalid password
	}

	// Create a user instance with invalid credentials
	invalidUser := model.Users{
		Email:    "invalid@example.com",
		Password: "hashedInvalidPassword", // Replace with the hashed invalid password
		IsActive: true,
		// Other user fields...
	}

	// Mock the UserRepository's FindByUsernameEmailPhoneNumber function to return the invalid user
	suite.userRepoMock.On("FindByUsernameEmailPhoneNumber", invalidLoginRequest.LoginOption.Email).Return(invalidUser, nil)

	// Call the Login function with the invalid payload
	_, err := suite.userUseCase.Login(invalidLoginRequest)

	// Assert that the login failed due to invalid credentials
	assert.Error(suite.T(), err)
	assert.EqualError(suite.T(), err, "unauthorized: invalid credential")

	// Verify that the expectations were met
	suite.userRepoMock.AssertExpectations(suite.T())
}

func (suite *UserUseCaseTestSuite) TestRegister_Success() {
	registerRequest := req.AuthRegisterRequest{
		FullName:        "agus bagus",
		Email:           "agus@mail.com",
		PhoneNumber:     "08121234567",
		UserName:        "agus",
		Password:        "password123",
		PasswordConfirm: "password123",
	}
	suite.userRepoMock.On("Save", mock.AnythingOfType("model.Users")).Return(nil)
	suite.userWalletMock.On("CreateWallet", mock.AnythingOfType("model.Wallet")).Return(nil)

	err := suite.userUseCase.Register(registerRequest)
	assert.NoError(suite.T(), err)
	suite.userRepoMock.AssertExpectations(suite.T())
	suite.userWalletMock.AssertExpectations(suite.T())
}

func (suite *UserUseCaseTestSuite) TestRegister_Failure() {
	registerRequest := req.AuthRegisterRequest{
		FullName:        "agus bagus",
		Email:           "agus@mail.com",
		PhoneNumber:     "08121234567",
		UserName:        "agus",
		Password:        "password123",
		PasswordConfirm: "password123",
	}

	// Mock a failure response from the user repository
	suite.userRepoMock.On("Save", mock.AnythingOfType("model.Users")).Return(errors.New("registration failed"))

	// Mock a success response from the wallet use case to isolate the registration failure
	suite.userWalletMock.On("CreateWallet", mock.AnythingOfType("model.Wallet")).Return(nil)

	err := suite.userUseCase.Register(registerRequest)
	assert.Error(suite.T(), err)
	assert.EqualError(suite.T(), err, "registration failed")

	// Verify that the expectations were met
	suite.userRepoMock.AssertExpectations(suite.T())
	suite.userWalletMock.AssertExpectations(suite.T())
}

func (suite *UserUseCaseTestSuite) TestFindByUserByPhoneNumber_Success() {
	mockData := model.Users{
		PhoneNumber: "081256789102",
	}

	suite.userRepoMock.On("FindByPhoneNumber", mockData.PhoneNumber).Return(mockData, nil)
	got, err := suite.userUseCase.FindByUserByPhoneNumber(mockData.PhoneNumber)
	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), mockData, got)
}

func (suite *UserUseCaseTestSuite) TestFindByUserByPhoneNumber_Failure() {
	phoneNumber := "081256789102"

	// Mock the SQL query for finding a user by phone number and return an error
	suite.userRepoMock.On("FindByPhoneNumber", phoneNumber).Return(model.Users{}, errors.New("database error"))

	got, err := suite.userUseCase.FindByUserByPhoneNumber(phoneNumber)

	// Assert that the error is not nil and contains the expected error message
	assert.Error(suite.T(), err)
	assert.EqualError(suite.T(), err, "database error")

	// Assert that the result (got) is empty or zero-value (model.Users{})
	assert.Equal(suite.T(), model.Users{}, got)

	// Verify that the expectations were met
	suite.userRepoMock.AssertExpectations(suite.T())
}

func (suite *UserUseCaseTestSuite) TestListsUsersHandler_Success() {
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
		},
	}

	suite.userRepoMock.On("FindAll").Return(mockData, nil)
	got, err := suite.userUseCase.ListsUsersHandler()
	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), mockData, got)
	assert.Equal(suite.T(), len(mockData), 1)
}

func (suite *UserUseCaseTestSuite) TestListsUsersHandler_Failure() {
	// Mock the SQL query for finding all users and make it return an error
	suite.userRepoMock.On("FindAll").Return(nil, errors.New("database error"))

	got, err := suite.userUseCase.ListsUsersHandler()
	assert.Error(suite.T(), err) // Expecting an error
	assert.Nil(suite.T(), got)   // Expecting nil result

	// Verify that the expectations were met
	suite.userRepoMock.AssertExpectations(suite.T())
}

func (suite *UserUseCaseTestSuite) TestUpdateAccount_Success() {
	updateAccountRequest := req.UpdateAccountRequest{
		Id:          "1",
		FullName:    "Sanju Zi",
		Username:    "sanzi",
		Email:       "sanzi@mail.com",
		PhoneNumber: "085812345678",
	}
	suite.userRepoMock.On("UpdateAccount", updateAccountRequest).Return(nil)
	err := suite.userUseCase.UpdateAccount(updateAccountRequest)
	assert.NoError(suite.T(), err)
	suite.userRepoMock.AssertExpectations(suite.T())
}

func (suite *UserUseCaseTestSuite) TestUpdateAccount_Fail() {
	updateAccountRequest := req.UpdateAccountRequest{
		Id:          "1",
		FullName:    "Sanju Zi",
		Username:    "sanzi",
		Email:       "sanzi@mail.com",
		PhoneNumber: "085812345678",
	}
	suite.userRepoMock.On("UpdateAccount", updateAccountRequest).Return(errors.New("update failed"))
	err := suite.userUseCase.UpdateAccount(updateAccountRequest)
	assert.Error(suite.T(), err)
	assert.EqualError(suite.T(), err, "failed update username: update failed")
	suite.userRepoMock.AssertExpectations(suite.T())
}

func (suite *UserUseCaseTestSuite) TestChangePasswordAccount_Success() {
	// Define a user with a valid current password
	user := model.Users{
		UserName: "testuser",
		Password: "old_password",
	}

	// Mock the repository to return the user and succeed in updating the password
	suite.userRepoMock.On("FindByUserName", user.UserName).Return(user, nil)
	suite.userRepoMock.On("UpdatePassword", user.UserName, mock.Anything, mock.Anything).Return(nil)

	// Create a request with valid data
	request := req.UpdatePasswordRequest{
		UserName:           user.UserName,
		CurrentPassword:    "old_password",
		NewPassword:        "new_password",
		NewPasswordConfirm: "new_password",
	}

	err := suite.userUseCase.ChangePasswordAccount(request)
	assert.NoError(suite.T(), err) // Expecting no error

	// Verify that the expectations were met
	suite.userRepoMock.AssertExpectations(suite.T())
}

func (suite *UserUseCaseTestSuite) TestChangePasswordAccount_Failure() {
	// Define a user with a valid current password
	user := model.Users{
		UserName: "testuser",
		Password: "old_password",
	}

	// Mock the repository to return the user
	suite.userRepoMock.On("FindByUserName", user.UserName).Return(user, nil)

	// Create a request with an incorrect current password
	request := req.UpdatePasswordRequest{
		UserName:           user.UserName,
		CurrentPassword:    "wrong_password", // This is incorrect
		NewPassword:        "new_password",
		NewPasswordConfirm: "new_password",
	}

	err := suite.userUseCase.ChangePasswordAccount(request)
	assert.Error(suite.T(), err) // Expecting an error

	// Verify that the expectations were met
	suite.userRepoMock.AssertExpectations(suite.T())
}

func (suite *UserUseCaseTestSuite) TestDisableAccount_Success() {
	userID := "12345"

	user := model.Users{
		Id:          userID,
		FullName:    "komarudin",
		Email:       "komar@main.com",
		PhoneNumber: "081212122121",
		UserName:    "komarudin",
		DisableAt:   time.Time{},
	}

	suite.userRepoMock.On("FindById", userID).Return(user, nil)

	disableTime := time.Now()
	user.DisableAt = disableTime

	suite.userRepoMock.On("DisableUserId", userID, disableTime).Return(int64(1), nil).Once()

	disabledUser, err := suite.userUseCase.DisableAccount(userID, disableTime)
	assert.NoError(suite.T(), err)
	assert.NotEqual(suite.T(), time.Time{}, disabledUser.DisableAt)

	suite.userRepoMock.AssertExpectations(suite.T())
}

func (suite *UserUseCaseTestSuite) TestDisableAccount_NotFound() {
	userID := "XX"

	suite.userRepoMock.On("FindById", userID).Return(model.Users{}, errors.New("user not found"))

	_, err := suite.userUseCase.DisableAccount(userID, time.Now())

	assert.Error(suite.T(), err)
	assert.EqualError(suite.T(), err, "user not found")

	// Ensure that the expected repository methods were called
	suite.userRepoMock.AssertExpectations(suite.T())
}

func (suite *UserUseCaseTestSuite) TestFindByUserName_Success() {
	mockData := model.Users{
		UserName: "Ags",
	}

	suite.userRepoMock.On("FindByUserName", mockData.UserName).Return(mockData, nil)
	got, err := suite.userUseCase.FindByUserName(mockData.UserName)
	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), mockData, got)
}

func (suite *UserUseCaseTestSuite) TestFindByUserName_Failure() {
	// Define the expected error and a username for the test case
	expectedError := errors.New("error: user not found")
	username := "NonExistentUser"

	// Mock the SQL query for finding a user by username and make it return an error
	suite.userRepoMock.On("FindByUserName", username).Return(model.Users{}, expectedError)

	// Call the function under test
	_, err := suite.userUseCase.FindByUserName(username)

	// Assert that the error matches the expected error
	assert.Error(suite.T(), err)
	assert.EqualError(suite.T(), err, expectedError.Error())

	// Verify that the expectations were met
	suite.userRepoMock.AssertExpectations(suite.T())
}

func (suite *UserUseCaseTestSuite) TestFindById_Success() {
	mockData := model.Users{
		Id: "1",
	}

	suite.userRepoMock.On("FindById", mockData.Id).Return(mockData, nil)
	got, err := suite.userUseCase.FindById(mockData.Id)
	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), mockData, got)
}

func (suite *UserUseCaseTestSuite) TestFindById_Failure() {
	userID := "1"

	// Modify the mock behavior to return an error
	suite.userRepoMock.On("FindById", userID).Return(model.Users{}, errors.New("user not found"))

	_, err := suite.userUseCase.FindById(userID)

	// Assert that an error is returned
	assert.Error(suite.T(), err)
	assert.EqualError(suite.T(), err, "user not found")

	// Verify that the expectations were met
	suite.userRepoMock.AssertExpectations(suite.T())
}

func (suite *UserUseCaseTestSuite) TestFindByUsernameEmailPhoneNumber_Success() {
	identifier := "komarudin@mail.com"

	user := model.Users{
		Id:          "12345",
		FullName:    "komar",
		Email:       "komarudin@mail.com",
		PhoneNumber: "085812345678",
		UserName:    "komar",
	}

	suite.userRepoMock.On("FindByUsernameEmailPhoneNumber", identifier).Return(user, nil)
	resultUser, err := suite.userUseCase.FindByUsernameEmailPhoneNumber(identifier)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), user, resultUser)
	suite.userRepoMock.AssertExpectations(suite.T())
}

func (suite *UserUseCaseTestSuite) TestFindByUsernameEmailPhoneNumber_UserNotFound() {
	identifier := "komarudin@mail.com"
	suite.userRepoMock.On("FindByUsernameEmailPhoneNumber", identifier).Return(model.Users{}, errors.New("user not found"))
	_, err := suite.userUseCase.FindByUsernameEmailPhoneNumber(identifier)
	assert.Error(suite.T(), err)
	assert.EqualError(suite.T(), err, "user not found")
	suite.userRepoMock.AssertExpectations(suite.T())
}
