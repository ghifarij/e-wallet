package req

type TransferRequest struct {
	UserId              string `json:"source_user_id"`
	SourceWalletID      string `json:"source_wallet_id"`
	DestinationWalletID string `json:"destination_wallet_id"`
	Amount              int    `json:"amount"`
	Description         string `json:"description"`
	PaymentMethodID     string `json:"payment_method_id"`
}
