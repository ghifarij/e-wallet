package req

type AuthLoginRequest struct {
	LoginOption loginOption
	Password    string `json:"password" validate:"required"`
}

type loginOption struct {
	UserName    string `json:"userName,omitempty"`
	Email       string `json:"email,omitempty"`
	PhoneNumber string `json:"phoneNumber,omitempty"`
}
