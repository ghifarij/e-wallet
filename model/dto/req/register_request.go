package req

type AuthRegisterRequest struct {
	FullName        string `validate:"required,min=3,max=50"`
	Email           string `validate:"required,email"`
	PhoneNumber     string `validate:"required,min=10,max=15"`
	UserName        string `validate:"required,min=3,max=30"`
	Password        string `validate:"required"`
	PasswordConfirm string `validate:"required,eqfield=Password"`
}
