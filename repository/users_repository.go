package repository

import (
	"Kelompok-2/dompet-online/model"
	"Kelompok-2/dompet-online/model/dto/req"
	"database/sql"
	"time"
)

type UserRepository interface {
	Save(user model.Users) error
	FindByUserName(username string) (model.Users, error)
	FindById(id string) (model.Users, error)
	FindByPhoneNumber(phoneNumber string) (model.Users, error)
	FindByUsernameEmailPhoneNumber(identifier string) (model.Users, error)
	UpdatePassword(username string, newPassword string, newPasswordConfirm string) error
	UpdateAccount(payload req.UpdateAccountRequest) error
	FindAll() ([]model.Users, error)
	DisableUserId(id string, disableTime time.Time) (model.Users, error)
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
	_, err := u.db.Exec(`INSERT INTO users(id, full_name, user_name, email, phone_number, password, password_confirm, is_active, created_at, updated_at, disable_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`,
		user.Id,
		user.FullName,
		user.UserName,
		user.Email,
		user.PhoneNumber,
		user.Password,
		user.PasswordConfirm,
		user.IsActive,
		user.CreatedAt,
		user.UpdatedAt,
		user.DisableAt,
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

func (u *userRepository) UpdateAccount(payload req.UpdateAccountRequest) error {
	_, err := u.db.Exec("UPDATE users SET full_name =$2, user_name =$3, email =$4, phone_number =$5 WHERE id =$1", payload.Id, payload.FullName, payload.Username, payload.Email, payload.PhoneNumber)
	if err != nil {
		return err
	}
	return nil

}

func (u *userRepository) FindAll() ([]model.Users, error) {
	rows, err := u.db.Query("SELECT id, full_name, user_name, email, phone_number, password, password_confirm, is_active, created_at, updated_at, disable_at FROM users")
	if err != nil {
		return nil, err
	}
	var users []model.Users
	for rows.Next() {
		var user model.Users
		err := rows.Scan(
			&user.Id,
			&user.FullName,
			&user.UserName,
			&user.Email,
			&user.PhoneNumber,
			&user.Password,
			&user.PasswordConfirm,
			&user.IsActive,
			&user.CreatedAt,
			&user.UpdatedAt,
			&user.DisableAt,
		)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (u *userRepository) FindByPhoneNumber(phoneNumber string) (model.Users, error) {
	row := u.db.QueryRow("SELECT id, full_name, user_name, email, phone_number, password, password_confirm, is_active, created_at, updated_at, disable_at FROM users WHERE phone_number = $1", phoneNumber)
	var user model.Users
	err := row.Scan(
		&user.Id,
		&user.FullName,
		&user.UserName,
		&user.Email,
		&user.PhoneNumber,
		&user.Password,
		&user.PasswordConfirm,
		&user.IsActive,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.DisableAt,
	)
	if err != nil {
		return model.Users{}, err
	}
	return user, nil
}

func (u *userRepository) FindById(id string) (model.Users, error) {

	row := u.db.QueryRow("SELECT id FROM users WHERE id = $1", id)
	var users model.Users
	err := row.Scan(&users.Id)
	if err != nil {
		return model.Users{}, err
	}
	return users, nil
}

func (u *userRepository) FindByUsernameEmailPhoneNumber(identifier string) (model.Users, error) {
	row := u.db.QueryRow("SELECT id, full_name, user_name, email, phone_number, password, password_confirm, is_active, created_at, updated_at, disable_at FROM users WHERE user_name = $1 OR email = $2 OR phone_number = $3", identifier, identifier, identifier)
	var user model.Users
	err := row.Scan(
		&user.Id,
		&user.FullName,
		&user.UserName,
		&user.Email,
		&user.PhoneNumber,
		&user.Password,
		&user.PasswordConfirm,
		&user.IsActive,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.DisableAt,
	)
	if err != nil {
		return model.Users{}, err
	}
	return user, nil
}

func (u *userRepository) DisableUserId(id string, disableTime time.Time) (model.Users, error) {
	_, err := u.db.Exec("UPDATE users SET is_active = false, disable_at = $2 WHERE id = $1", id, disableTime)
	if err != nil {
		return model.Users{}, err
	}
	return model.Users{}, nil
}
