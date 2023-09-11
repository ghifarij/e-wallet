package req

type AuthLoginRequest struct {
	UserName string `validate:"required"`
	Password string `validate:"required"`
}
