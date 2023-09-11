package req

type UpdatePasswordRequest struct {
	UserName           string `validate:"required,min=3,max=50"`
	CurrentPassword    string `validate:"required"`
	NewPassword        string `validate:"required"`
	NewPasswordConfirm string `validate:"required,eqfield=NewPassword"`
}
