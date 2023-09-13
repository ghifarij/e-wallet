package req

type WalletRequestBody struct {
	UserId string `validate:"required" json:"userId"`
}
