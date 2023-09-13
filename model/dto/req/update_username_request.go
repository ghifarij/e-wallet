package req

type UpdateUserNameRequest struct {
	Id       string `validate:"required"`
	Username string `validate:"required"`
}
