package web

import (
	"fmt"
	"hex_test/user/application/in"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	user in.IUser
}

func NewUserController(ILogin in.IUser) UserController {
	return UserController{
		user: ILogin,
	}
}

func (u *UserController) Login(ginCtx *gin.Context) {
	userResult, err := u.user.Login("testUserName ", "testPassword")

	if err != nil {
		ginCtx.AbortWithError(400, nil)
	}

	fmt.Println(userResult)
}

func (u *UserController) Logout(ginCtx *gin.Context) {
	err := u.user.Logout("token")

	if err != nil {
		ginCtx.AbortWithError(400, nil)
	}
}

func (u *UserController) GetTokenInfo(ginCtx *gin.Context) {
	info, err := u.user.GetTokenInfo("token")

	if err != nil {
		ginCtx.AbortWithError(400, nil)
	}

	fmt.Println(info)
}
