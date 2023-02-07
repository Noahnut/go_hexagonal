package application

import (
	"hex_test/transfer/adapter/out/persistence"
	"hex_test/transfer/application/in"
	"hex_test/transfer/application/out"
)

type transfer struct {
	outTransfer out.ITransfer
}

func NewTransfer(databaseImp string) in.ITransfer {
	return &transfer{
		outTransfer: persistence.NewtransferPersistence(databaseImp),
	}
}

func (t *transfer) TransferMoney(toAccount string, number int64) error {
	return nil
}
