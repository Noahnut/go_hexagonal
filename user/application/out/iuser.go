package out

type IUser interface {
	Login(username, password string) (string, error)
	Logout(token string) error
	GetTokenInfo(token string) (string, error)
}
