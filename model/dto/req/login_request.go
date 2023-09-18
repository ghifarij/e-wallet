package req

type AuthLoginRequest struct {
	LoginOption string `json:"login_option"`
	UserName    string `json:"userName,omitempty"`
	Email       string `json:"email,omitempty"`
	PhoneNumber string `json:"phoneNumber,omitempty"`
	Password    string `json:"password" validate:"required"`
}
