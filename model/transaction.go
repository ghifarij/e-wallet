package model

import "time"

type Transactions struct {
	Id              string
	SourceWalletID  string
	UserId          string
	PaymentMethodID string
	Destination     string
	Amount          int
	Description     string
	CreateAt        time.Time
}
