package resp

type LoginResponse struct {
	Status   int         `json:"status"`
	UserName string      `json:"user_name"`
	Message  string      `json:"message"`
	Data     interface{} `json:"data"`
	Token    string      `json:"token"`
}
