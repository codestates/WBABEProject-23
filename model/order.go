package model

import (
	"context"
	"fmt"
	"reflect"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
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
	Id           primitive.ObjectID `bson:"_id"`
	OrderID      int64              `bson:"orderid`
	Orderer      string             `bson:"orderer"`
	Status       int                `bson:"status"`
	BusinessName string             `bson:"businessname`
	Menu         []MenuNum          `bson:"menu`
	CreatedAt    time.Time          `bson:"createdat`
}

type MenuNum struct {
	MenuName   string `bson:"menuname"`
	Number     int    `bson:"number"`
	IsReviewed bool   `bson:"isreviewed"`
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

func (m *Model) ListOrder(userName string) []Order {
	filter := bson.M{"orderer": userName}
	option := options.Find().SetProjection(bson.M{"orderer": 1, "status": 1, "businessname": 1, "menu": 1, "createdat": 1})
	cursor, err := m.colOrder.Find(context.TODO(), filter, option)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer cursor.Close(context.TODO())

	var orders []Order
	if err = cursor.All(context.TODO(), &orders); err != nil {
		fmt.Println(err)
		panic(err)
	}
	return orders
}

func (m *Model) ModifyOrder(business primitive.ObjectID, menu []MenuNum) bool {
	filter := bson.M{"_id": business}

	var order Order
	err := m.colOrder.FindOne(context.TODO(), filter).Decode(&order)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	addition := false
	if reflect.DeepEqual(order.Menu, menu) {
		addition = true
	}
	status := order.Status
	switch status {
	case DeliverComplete, Delivering, ReceiptCancled, AdditionalReceiptCancled:
		return false
	case ReceiptCooking, AdditionalReceiptCooking:
		if !addition {
			return false
		}
	}
	update := bson.M{"$set": bson.M{"menu": menu, "status": AdditionalReceipting}}
	result, err := m.colOrder.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
	return true
}
