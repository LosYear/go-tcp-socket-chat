package shared

const (
	LoginActionName            = "login"
	LogoutActionName           = "logout"
	GetUsersActionName         = "users"
	GetUsersCountActionName    = "userscount"
	SendMessageActionName      = "sendmessage"
	NewMessageNotificationName = "newmessagenotification"
)

type Request struct {
	Name    string `json:"name"`
	Payload string `json:"payload,omitempty"`
}

type Response struct {
	Name    string      `json:"name"`
	Error   bool        `json:"error"`
	Message string      `json:"message"`
	Payload interface{} `json:"payload,omitempty"`
}

type TextMessage struct {
	Username string `json:"username"`
	Text     string `json:"text"`
}
