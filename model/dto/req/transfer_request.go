package req

type TransferRequest struct {
	UserId      string `json:"user_id" binding:"required"`
	Destination string `json:"destination" binding:"required"`
	Amount      int    `json:"amount" binding:"required"`
	Description string `json:"description"`
}
