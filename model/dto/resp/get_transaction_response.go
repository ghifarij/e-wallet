package resp

import (
	"Kelompok-2/dompet-online/model"
	"time"
)

type GetTransactionsResponse struct {
	Id            string               `json:"id_transaction"`
	SourceOfFound model.SourceOfFounds `json:"source_of_found"`
	User          user                 `json:"user"`
	DestinationId string               `json:"destination_id"`
	Wallet        wallet               `json:"wallet"`
	Amount        int                  `json:"amount"`
	Description   string               `json:"description"`
	CreateAt      time.Time            `json:"time_of_transaction"`
}

type user struct {
	UserId   string
	FullName string
}

type wallet struct {
	RekeningUser string
	Balance      int
}
