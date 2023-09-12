package usecase

import (
	"Kelompok-2/dompet-online/model"
	"Kelompok-2/dompet-online/model/dto/req"
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
}

type userUseCase struct {
	repo repository.UserRepository
}

func NewUserUseCase(repo repository.UserRepository) UserUseCase {
	return &userUseCase{
		repo: repo,
	}
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
	// validasi password dengan passwordConfirm
	//if payload.Password != payload.PasswordConfirm {
	//	return fmt.Errorf("password and password confirmation do not match")
	//}

	user := model.Users{
		Id:              common.GenerateID(),
		FullName:        payload.FullName,
		Email:           payload.Email,
		PhoneNumber:     payload.PhoneNumber,
		UserName:        payload.UserName,
		Password:        hashPassword,
		PasswordConfirm: hashPasswordConfirm,
		CreatedAt:       time.Now(),
		//UpdatedAt:       time.Now(),
		//DeleteAt:        time.Time{},
	}

	err = u.repo.Save(user)
	if err != nil {
		return fmt.Errorf("failed save user: %v", err.Error())
	}
	return nil
}
