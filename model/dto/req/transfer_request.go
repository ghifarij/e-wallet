package req

type TransferRequest struct {
	Amount       int    `validate:"required"`
	RekeningUser string `validate:"required"`
	Description  string `validate:"required"`
}
