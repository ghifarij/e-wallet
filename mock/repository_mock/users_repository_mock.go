package repository_mock

import (
	"Kelompok-2/dompet-online/model"
	"Kelompok-2/dompet-online/model/dto/req"
	"github.com/stretchr/testify/mock"
	"time"
)

type UserRepositoryMock struct {
	mock.Mock
}

func (u *UserRepositoryMock) Save(user model.Users) error {
	return u.Called(user).Error(0)
}

func (u *UserRepositoryMock) FindByUserName(username string) (model.Users, error) {
	args := u.Called(username)
	if args.Get(1) != nil {
		return model.Users{}, args.Error(1)
	}
	return args.Get(0).(model.Users), nil
}

func (u *UserRepositoryMock) FindById(id string) (model.Users, error) {
	args := u.Called(id)
	if args.Get(1) != nil {
		return model.Users{}, args.Error(1)
	}
	return args.Get(0).(model.Users), nil
}

func (u *UserRepositoryMock) FindByPhoneNumber(phoneNumber string) (model.Users, error) {
	args := u.Called(phoneNumber)
	if args.Get(1) != nil {
		return model.Users{}, args.Error(1)
	}
	return args.Get(0).(model.Users), nil
}

func (u *UserRepositoryMock) FindByUsernameEmailPhoneNumber(identifier string) (model.Users, error) {
	args := u.Called(identifier)
	if args.Get(1) != nil {
		return model.Users{}, args.Error(1)
	}
	return args.Get(0).(model.Users), nil
}

func (u *UserRepositoryMock) UpdatePassword(username string, newPassword string, newPasswordConfirm string) error {
	return u.Called(username, newPassword, newPassword, newPasswordConfirm).Error(0)
}

func (u *UserRepositoryMock) UpdateAccount(payload req.UpdateAccountRequest) error {
	return u.Called(payload).Error(0)
}

func (u *UserRepositoryMock) FindAll() ([]model.Users, error) {
	args := u.Called()
	if args.Get(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]model.Users), nil
}

func (u *UserRepositoryMock) DisableUserId(id string, disableTime time.Time) (model.Users, error) {
	args := u.Called(id, disableTime)
	if args.Get(1) != nil {
		return model.Users{}, args.Error(1)
	}
	return args.Get(0).(model.Users), nil
}
