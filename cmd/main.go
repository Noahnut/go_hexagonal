package main

import (
	transfer_web "hex_test/transfer/adapter/in/web"
	transfer_application "hex_test/transfer/application"
	user_web "hex_test/user/adapter/in/web"
	user_application "hex_test/user/application"

	"github.com/gin-gonic/gin"
)

func main() {

	userController := user_web.NewUserController(user_application.NewUser("DBImp"))
	transferController := transfer_web.NewTransferController(transfer_application.NewTransfer("DBImp"))

	var router *gin.Engine

	router.GET("/login", userController.Login)
	router.GET("/logout", userController.Logout)
	router.GET("/token_info", userController.GetTokenInfo)

	router.GET("/transfer", transferController.TransferMoney)

}
