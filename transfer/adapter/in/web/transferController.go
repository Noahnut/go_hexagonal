package web

import (
	"hex_test/transfer/application/in"

	"github.com/gin-gonic/gin"
)

type TransferController struct {
	transfer in.ITransfer
}

func NewTransferController(transfer in.ITransfer) TransferController {
	return TransferController{
		transfer: transfer,
	}
}

func (t *TransferController) TransferMoney(ginCtx *gin.Context) {
	t.transfer.TransferMoney("account", 123)
}
