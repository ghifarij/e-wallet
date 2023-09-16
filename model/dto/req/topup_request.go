package req

type TopUpRequest struct {
	UserId          string `validate:"required"`
	WalletID        string `validate:"required"`
	Amount          int    `validate:"required,min=10000"`
	PaymentMethodId string `validate:"required"`
}
