package domain

type User struct {
	UserName string `json:"username"`
	Password string `json:"password"`
	Token    string `json:"token"`
}
