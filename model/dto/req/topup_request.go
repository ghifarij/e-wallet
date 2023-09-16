package req

type TopUpRequest struct {
	UserId        string `validate:"required"`
	RekeningUser  string `validate:"required"`
	Amount        int    `validate:"required,min=10000"`
	PaymentMethod string `validate:"required,oneof=BRI BCA"`
}
