package persistence

import "hex_test/transfer/application/out"

type transferPersistence struct {
	DBImp string
}

func NewtransferPersistence(DBImp string) out.ITransfer {
	return &transferPersistence{
		DBImp: DBImp,
	}
}

func (t *transferPersistence) TransferMoney(toAccount string, number int64) error {
	return nil
}
