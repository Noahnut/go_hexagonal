package persistence

import "hex_test/user/application/out"

type userPersistence struct {
	DBImp string
}

func NewUserPersistence(DBImp string) out.IUser {
	return &userPersistence{
		DBImp: DBImp,
	}
}

func (u *userPersistence) Login(username, password string) (string, error) {
	return "token", nil
}

func (u *userPersistence) Logout(token string) error {
	return nil
}

func (u *userPersistence) GetTokenInfo(token string) (string, error) {
	return "", nil
}
