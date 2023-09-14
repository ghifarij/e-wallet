package req

type UpdateAccountRequest struct {
	Id          string `validate:"required"`
	FullName    string `validate:"required,min=3,max=50"`
	Username    string `validate:"required,min=3,max=30"`
	Email       string `validate:"required,email"`
	PhoneNumber string `validate:"required,min=10,max=15"`
}
