package req

type GetTransactionRequest struct {
	UserId       string `validate:"required"`
	RekeningUser string `validate:"required"`
}
