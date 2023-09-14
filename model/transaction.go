package model

import "time"

type Transactions struct {
	Id              string
	SourceOfFoundId string
	UserId          string
	Destination     string
	Amount          int
	Description     string
	CreateAt        time.Time
}
