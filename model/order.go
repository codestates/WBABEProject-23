package model

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	Receipting = iota + 1
	ReceiptCancled
	ReceiptComplete
	AdditionalReceipt
	ReceiptCooking
	Delivering
	DeliverComplete
)

type Order struct {
	Id           primitive.ObjectID `bson:"_id"`
	OrderID      int64              `bson:"orderId`
	Orderer      primitive.ObjectID `bson:"orderer"`
	Status       int                `bson:"status"`
	BusinessName string             `bson:"businessName`
	Menu         []MenuNum          `bson:"menu`
	CreatedAt    time.Time          `bson:"createdAt`
}

type MenuNum struct {
	MenuName string `bson:"menu_name"`
	Number   int    `bson:"number"`
}

func (m *Model) MakeOrder(order Order) {
	year, month, day := order.CreatedAt.Date()
	filter := bson.M{"created-at": bson.M{
		"$gte": bson.M{"$dateFromParts": bson.M{"year": year, "month": month, "day": day}},
		"$lt":  bson.M{"$dateFromParts": bson.M{"year": year, "month": month, "day": day + 1}},
	}}
	count, err := m.colOrder.CountDocuments(context.TODO(), filter)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	order.OrderID = count + 1
	fmt.Println(order)
	result, err := m.colOrder.InsertOne(context.TODO(), order)
	if err != nil {
		panic(err)
	}
	fmt.Println(result.InsertedID)
}
