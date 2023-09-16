package resp

import (
	"time"
)

type GetTransactionsResponse struct {
	Id            string    `json:"id_transaction"`
	Destination   string    `json:"destination"`
	Amount        int       `json:"amount"`
	Description   string    `json:"description"`
	CreateAt      time.Time `json:"time_of_transaction"`
	User          user      `json:"user"`
	Wallet        wallet    `json:"wallet"`
	PaymentMethod paymentMethod
}

type paymentMethod struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type user struct {
	UserName string `json:"user_name"`
}

type wallet struct {
	RekeningUser string `json:"rekening_user"`
	Balance      int    `json:"balance"`
}
