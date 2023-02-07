package application

import (
	"hex_test/user/adapter/out/persistence"
	"hex_test/user/application/in"
	"hex_test/user/application/out"
	"hex_test/user/domain"
)

type user struct {
	outLogin out.IUser
}

func NewUser(databaseImp string) in.IUser {
	return &user{
		outLogin: persistence.NewUserPersistence(databaseImp),
	}
}

func (u *user) Login(username, password string) (domain.User, error) {
	token, err := u.outLogin.Login(username, password)

	if err != nil {
		return domain.User{}, err
	}

	return domain.User{
		Token: token,
	}, nil
}

func (u *user) Logout(token string) error {
	return nil
}

func (u *user) GetTokenInfo(token string) (string, error) {
	return "", nil
}
