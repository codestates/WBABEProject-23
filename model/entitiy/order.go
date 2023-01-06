package entitiy

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	Receipting = iota + 1
	ReceiptCancled
	ReceiptComplete
	AdditionalReceipting
	AdditionalReceiptingComplete
	AdditionalReceiptCancled
	AdditionalReceiptCooking
	ReceiptCooking
	Delivering
	DeliverComplete
)

type Order struct {
	ID        primitive.ObjectID `bson:"_id"`
	OrderID   int64              `bson:"orderid"`
	BID       primitive.ObjectID `bson:"business_id"`
	Orderer   string             `bson:"orderer"`
	State     int                `bson:"state"`
	Menu      []MenuNum          `bson:"menu"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
}

type MenuNum struct {
	MenuID     primitive.ObjectID `bson:"menu_id"`
	Number     int                `bson:"number"`
	IsReviewed bool               `bson:"is_reviewed"`
}
