package usecase

import (
	"Kelompok-2/dompet-online/exception"
	"Kelompok-2/dompet-online/model"
	"Kelompok-2/dompet-online/model/dto/req"
	"Kelompok-2/dompet-online/model/dto/resp"
	"Kelompok-2/dompet-online/repository"
	"Kelompok-2/dompet-online/util/common"
	"Kelompok-2/dompet-online/util/security"
	"fmt"
	"github.com/go-playground/validator/v10"
	"time"
)

type UserUseCase interface {
	FindByUserName(username string) (model.Users, error)
	FindAll() ([]model.Users, error)
	Register(payload req.AuthRegisterRequest) error
	UpdateAccount(payload req.UpdateAccountRequest) error
	DeleteById(id string) error
	FindByPhoneNumber(phoneNumber string) (model.Users, error)
	Login(payload req.AuthLoginRequest) (resp.LoginResponse, error)
	ChangePassword(payload req.UpdatePasswordRequest) error
	FindById(Id string) (model.Users, error)
	FindByUsernameEmailPhoneNumber(identifier string) (model.Users, error)
	DisableUserId(id string) (model.Users, error)
}

type userUseCase struct {
	repo     repository.UserRepository
	walletUc WalletUseCase
}

func NewUserUseCase(repo repository.UserRepository, walletUc WalletUseCase) UserUseCase {
	return &userUseCase{
		repo:     repo,
		walletUc: walletUc,
	}
}

func (u *userUseCase) FindById(id string) (model.Users, error) {
	return u.repo.FindById(id)
}

func (u *userUseCase) DeleteById(id string) error {
	user, err := u.repo.FindById(id)
	if err != nil {
		return err
	}

	err = u.repo.DeleteById(user.Id)
	if err != nil {
		return fmt.Errorf("failed to delete users: %v", err)
	}
	return nil
}

// UpdateUsername implements UserUseCase.
func (u *userUseCase) UpdateAccount(payload req.UpdateAccountRequest) error {
	validate := validator.New()
	err := validate.Struct(payload)
	if err != nil {
		return err
	}

	if err := u.repo.UpdateAccount(payload); err != nil {
		return fmt.Errorf("failed update username: %v", err.Error())
	}
	return nil

}

// FindByUsername implements UserUseCase.
func (u *userUseCase) FindByUserName(username string) (model.Users, error) {
	return u.repo.FindByUserName(username)
}

// FindAll implements UserUseCase.
func (u *userUseCase) FindAll() ([]model.Users, error) {
	users, err := u.repo.FindAll()
	if err != nil {
		return nil, fmt.Errorf("failed to get all data: %v", err)
	}
	return users, nil
}

// Register implements UserUseCase.
func (u *userUseCase) Register(payload req.AuthRegisterRequest) error {
	// Validate the payload
	validate := validator.New()
	err := validate.Struct(payload)
	if err != nil {
		return err
	}

	hashPassword, err := security.HashPassword(payload.Password)
	if err != nil {
		return err
	}

	hashPasswordConfirm, err := security.HashPassword(payload.PasswordConfirm)
	if err != nil {
		return err
	}

	user := model.Users{
		Id:              common.GenerateID(),
		FullName:        payload.FullName,
		Email:           payload.Email,
		PhoneNumber:     payload.PhoneNumber,
		UserName:        payload.UserName,
		Password:        hashPassword,
		PasswordConfirm: hashPasswordConfirm,
		IsActive:        true,
		CreatedAt:       time.Now(),
	}

	err = u.repo.Save(user)
	if err != nil {
		return fmt.Errorf("failed save user: %v", err.Error())
	}

	wallet := model.Wallet{
		Id:           common.GenerateID(),
		UserId:       user.Id,
		RekeningUser: common.GenerateRandomRekeningNumber(10),
		Balance:      0,
	}

	// Panggil use case wallet untuk menyimpan wallet
	err = u.walletUc.CreateWallet(wallet)
	if err != nil {
		return fmt.Errorf("failed create wallet: %v", err.Error())
	}

	return nil
}

// FindByPhoneNumber UserUseCase
func (u *userUseCase) FindByPhoneNumber(phoneNumber string) (model.Users, error) {
	byPhoneNumber, err := u.repo.FindByPhoneNumber(phoneNumber)
	if err != nil {
		return model.Users{}, exception.NewNotFoundError(err.Error())
	}
	return byPhoneNumber, nil
}

// Login implements UserUseCase.
// Login implements UserUseCase.
func (u *userUseCase) Login(payload req.AuthLoginRequest) (resp.LoginResponse, error) {
	// Validasi payload menggunakan struct
	validate := validator.New()
	err := validate.Struct(payload)
	if err != nil {
		return resp.LoginResponse{}, err
	}

	var user model.Users

	// Determine which identifier is provided and use the appropriate method to find the user
	if payload.UserName != "" {
		user, err = u.FindByUsernameEmailPhoneNumber(payload.UserName)
	} else if payload.Email != "" {
		user, err = u.FindByUsernameEmailPhoneNumber(payload.Email)
	} else if payload.PhoneNumber != "" {
		user, err = u.FindByUsernameEmailPhoneNumber(payload.PhoneNumber)
	} else {
		return resp.LoginResponse{}, fmt.Errorf("invalid login request: no identifier provided")
	}

	if err != nil {
		return resp.LoginResponse{}, fmt.Errorf("unauthorized: invalid credential")
	}

	// Validasi Password
	err = security.VerifyPassword(user.Password, payload.Password)
	if err != nil {
		return resp.LoginResponse{}, fmt.Errorf("unauthorized: invalid credential")
	}

	// Generate Token
	token, err := security.GenerateJwtToken(user)
	if err != nil {
		return resp.LoginResponse{}, err
	}

	return resp.LoginResponse{
		UserName: user.UserName,
		Token:    token,
	}, nil
}

// ChangePassword implements UserUseCase.
func (u *userUseCase) ChangePassword(payload req.UpdatePasswordRequest) error {
	// Validasi payload menggunakan struct
	validate := validator.New()
	err := validate.Struct(payload)
	if err != nil {
		return err
	}

	// read Username di db
	user, err := u.repo.FindByUserName(payload.UserName)
	if err != nil {
		return err
	}

	// Validasi password saat ini
	err = security.VerifyPassword(user.Password, payload.CurrentPassword)
	if err != nil {
		return fmt.Errorf("update password failed: invalid current password")
	}

	// Hash new password and password confirmation
	hashedNewPassword, err := security.HashPassword(payload.NewPassword)
	if err != nil {
		return err
	}

	hashedNewPasswordConfirm, err := security.HashPassword(payload.NewPasswordConfirm)
	if err != nil {
		return err
	}

	// update password dan password confirm
	err = u.repo.UpdatePassword(user.UserName, hashedNewPassword, hashedNewPasswordConfirm)
	if err != nil {
		return fmt.Errorf("failed save user: %v", err.Error())
	}
	return nil
}

func (u *userUseCase) FindByUsernameEmailPhoneNumber(identifier string) (model.Users, error) {
	user, err := u.repo.FindByUsernameEmailPhoneNumber(identifier)
	if err != nil {
		return model.Users{}, err
	}
	return user, nil
}

func (u *userUseCase) DisableUserId(id string) (model.Users, error) {
	user, err := u.repo.FindById(id)
	if err != nil {
		return model.Users{}, err
	}

	_, err = u.repo.DisableUserId(user.Id)
	if err != nil {
		return model.Users{}, fmt.Errorf("failed to disable users: %v", err)
	}
	return user, nil
}
