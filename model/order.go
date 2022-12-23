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
	OrderID      int64              `bson:"orderid`
	Orderer      string             `bson:"orderer"`
	Status       int                `bson:"status"`
	BusinessName string             `bson:"businessname`
	Menu         []MenuNum          `bson:"menu`
	CreatedAt    time.Time          `bson:"createdat`
}

type MenuNum struct {
	MenuName string `bson:"menuname"`
	Number   int    `bson:"number"`
}

func (m *Model) MakeOrder(order Order) {
	now := time.Now().UTC()

	// Set the start and end of the range to the start and end of the desired day
	start := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)
	end := start.AddDate(0, 0, 1)
	filter := bson.M{"createdat": bson.M{
		"$gte": start,
		"$lt":  end,
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
