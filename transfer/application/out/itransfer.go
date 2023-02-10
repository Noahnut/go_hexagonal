package out

type ITransfer interface {
	TransferMoneyName(toAccount string, number int64) error
}
