package model

import "time"

type Users struct {
	Id              string
	FullName        string
	Email           string
	PhoneNumber     string
	UserName        string
	Password        string
	PasswordConfirm string
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeleteAt        time.Time
}
