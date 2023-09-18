package req

type AuthRegisterRequest struct {
	FullName        string `json:"fullName" validate:"required,min=3,max=50"`
	Email           string `json:"email" validate:"required,email"`
	PhoneNumber     string `json:"phoneNumber" validate:"required,min=10,max=15"`
	UserName        string `json:"userName" validate:"required,min=3,max=30"`
	Password        string `json:"password" validate:"required,"`
	PasswordConfirm string `json:"passwordConfirm" validate:"required,eqfield=Password"`
}
