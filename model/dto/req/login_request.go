package req

type AuthLoginRequest struct {
	UserName    string
	Email       string
	PhoneNumber string
	Password    string `validate:"required"`
}
