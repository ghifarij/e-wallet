package req

type TopUpRequest struct {
	UserId          string `json:"your_userId" validate:"required"`
	WalletID        string `json:"your_wallet_id" validate:"required"`
	Amount          int    `json:"topUp_amount" validate:"required,min=10000"`
	PaymentMethodId string `validate:"required"`
}
