package model

import "go.mongodb.org/mongo-driver/bson/primitive"

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
	Id      primitive.ObjectID `bson:"_id"`
	Orderer User               `bson:"orderer"`
	Status  int                `bson:"status"`
}
