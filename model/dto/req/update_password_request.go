package req

type UpdatePasswordRequest struct {
	UserName           string `validate:"required,min=3,max=50" json:"userName"`
	CurrentPassword    string `validate:"required" json:"currentPassword"`
	NewPassword        string `validate:"required" json:"newPassword"`
	NewPasswordConfirm string `validate:"required,eqfield=NewPassword" json:"newPasswordConfirm"`
}
