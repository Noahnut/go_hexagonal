package in

type ITransfer interface {
	TransferMoney(toAccount string, number int64) error
}
