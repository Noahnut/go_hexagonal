package in

import "hex_test/user/domain"

type IUser interface {
	Login(username, password string) (domain.User, error)
	Logout(token string) error
	GetTokenInfo(token string) (string, error)
}
