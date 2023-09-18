package usecase_mock

import (
	"Kelompok-2/dompet-online/model"
	"Kelompok-2/dompet-online/model/dto/req"
	"Kelompok-2/dompet-online/model/dto/resp"
	"github.com/stretchr/testify/mock"
)

type UserUseCaseMock struct {
	mock.Mock
}

func (u *UserUseCaseMock) FindById(id string) (model.Users, error) {
	args := u.Called(id)
	if args.Get(1) != nil {
		return model.Users{}, args.Error(1)
	}
	return args.Get(0).(model.Users), nil
}

func (u *UserUseCaseMock) DeleteById(id string) error {
	return u.Called(id).Error(0)
}

// UpdateUsername implements UserUseCase.
func (u *UserUseCaseMock) UpdateAccount(payload req.UpdateAccountRequest) error {
	return u.Called(payload).Error(0)
}

// FindByUsername implements UserUseCase.
func (u *UserUseCaseMock) FindByUserName(username string) (model.Users, error) {
	args := u.Called(username)
	if args.Get(1) != nil {
		return model.Users{}, args.Error(1)
	}
	return args.Get(0).(model.Users), nil
}

// FindAll implements UserUseCase.
func (u *UserUseCaseMock) FindAll() ([]model.Users, error) {
	args := u.Called()
	if args.Get(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]model.Users), nil
}

// Register implements UserUseCase.
func (u *UserUseCaseMock) Register(payload req.AuthRegisterRequest) error {
	return u.Called(payload).Error(0)
}

// FindByPhoneNumber UserUseCase
func (u *UserUseCaseMock) FindByPhoneNumber(phoneNumber string) (model.Users, error) {
	args := u.Called(phoneNumber)
	if args.Get(1) != nil {
		return model.Users{}, args.Error(1)
	}
	return args.Get(0).(model.Users), nil
}

// Login implements UserUseCase.
func (u *UserUseCaseMock) Login(payload req.AuthLoginRequest) (resp.LoginResponse, error) {
	args := u.Called(payload)
	if args.Get(1) != nil {
		return resp.LoginResponse{}, args.Error(1)
	}
	return args.Get(0).(resp.LoginResponse), nil
}

func (u *UserUseCaseMock) ChangePassword(payload req.UpdatePasswordRequest) error {
	return u.Called(payload).Error(0)
}

func (u *UserUseCaseMock) FindByUsernameEmailPhoneNumber(identifier string) (model.Users, error) {
	args := u.Called(identifier)
	if args.Get(1) != nil {
		return model.Users{}, args.Error(1)
	}
	return args.Get(0).(model.Users), nil
}

func (u *UserUseCaseMock) DisableUserId(id string) (model.Users, error) {
	args := u.Called(id)
	if args.Get(1) != nil {
		return model.Users{}, args.Error(1)
	}
	return args.Get(0).(model.Users), nil
}
