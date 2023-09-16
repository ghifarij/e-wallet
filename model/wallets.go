package model

import "time"

type Wallet struct {
	Id           string
	UserId       string
	RekeningUser string
	Balance      int
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
