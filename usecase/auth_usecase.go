package usecase

import (
	"Kelompok-2/dompet-online/model/dto/req"
	"Kelompok-2/dompet-online/model/dto/resp"
	"Kelompok-2/dompet-online/repository"
	"Kelompok-2/dompet-online/util/security"
	"fmt"
	"github.com/go-playground/validator/v10"
)

type AuthUseCase interface {
	Login(payload req.AuthLoginRequest) (resp.LoginResponse, error)
	ChangePassword(payload req.UpdatePasswordRequest) error
}

type authUseCase struct {
	repo repository.UserRepository
}

func NewAuthUseCase(repo repository.UserRepository) AuthUseCase {
	return &authUseCase{repo: repo}
}

// Login implements AuthUseCase.
func (a *authUseCase) Login(payload req.AuthLoginRequest) (resp.LoginResponse, error) {
	// Validasi payload menggunakan struct
	validate := validator.New()
	err := validate.Struct(payload)
	if err != nil {
		return resp.LoginResponse{}, err
	}

	// read Username di db
	user, err := a.repo.FindByUserName(payload.UserName)
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

func (a *authUseCase) ChangePassword(payload req.UpdatePasswordRequest) error {
	// Validasi payload menggunakan struct
	validate := validator.New()
	err := validate.Struct(payload)
	if err != nil {
		return err
	}

	// read Username di db
	user, err := a.repo.FindByUserName(payload.UserName)
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
	err = a.repo.UpdatePassword(user.UserName, hashedNewPassword, hashedNewPasswordConfirm)
	if err != nil {
		return fmt.Errorf("failed save user: %v", err.Error())
	}
	return nil
}
