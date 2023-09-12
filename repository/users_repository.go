package repository

import (
	"Kelompok-2/dompet-online/model"
	"database/sql"
)

type UserRepository interface {
	Save(user model.Users) error
	FindByUserName(username string) (model.Users, error)
	FindByPhoneNumber(phoneNumber string) (model.Users, error)
	UpdatePassword(username string, newPassword string, newPasswordConfirm string) error
	UpdateUserName(username string) error
	FindAll() ([]model.Users, error)
	DeleteById(id string) error
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

// Save implements UserRepository.
func (u *userRepository) Save(user model.Users) error {
	_, err := u.db.Exec(`INSERT INTO users(id, full_name, user_name, email, phone_number, password, password_confirm, created_at, updated_at, deleted_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`,
		user.Id,
		user.FullName,
		user.UserName,
		user.Email,
		user.PhoneNumber,
		user.Password,
		user.PasswordConfirm,
		user.CreatedAt,
		user.UpdatedAt,
		user.DeleteAt,
	)
	if err != nil {
		return err
	}
	return nil
}

// FindByUsername implements UserRepository.
func (u *userRepository) FindByUserName(username string) (model.Users, error) {

	row := u.db.QueryRow("SELECT id, user_name, password FROM users WHERE user_name = $1", username)
	var user model.Users
	err := row.Scan(
		&user.Id,
		&user.UserName,
		&user.Password,
	)
	if err != nil {
		return model.Users{}, err
	}
	return user, nil
}

func (u *userRepository) UpdatePassword(username string, newPassword string, newPasswordConfirm string) error {
	_, err := u.db.Exec("UPDATE users SET password = $1, password_confirm = $2 WHERE user_name = $3", newPassword, newPasswordConfirm, username)
	if err != nil {
		return err
	}
	return nil
}

func (u *userRepository) UpdateUserName(username string) error {
	_, err := u.db.Exec("UPADTE users SET username =$1", username)
	if err != nil {
		return err
	}
	return nil

}

func (u *userRepository) FindAll() ([]model.Users, error) {
	//TODO implement me
	panic("implement me")
}

func (u *userRepository) DeleteById(id string) error {
	//TODO implement me
	panic("implement me")
}

func (u *userRepository) FindByPhoneNumber(phoneNumber string) (model.Users, error) {
	//TODO implement me
	panic("implement me")
}
