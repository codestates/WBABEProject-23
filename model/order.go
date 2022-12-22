package model

const (
	Receipting = iota
	ReceiptCancled
	ReceiptComplete
	AdditionalReceipt
	ReceiptCooking
	Delivering
	DeliverComplete
)

type Order struct {
	Orderer User `bson:"orderer"`
	Status  int  `bson:"status"`
}
