package model

import (
	"context"
	"fmt"
	"reflect"
	"time"

	"go.mongodb.org/mongo-driver/mongo"

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
	State        int                `bson:"state"`
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

func (m *Model) ListOrder(userName string, cur bool) []Order {
	var filter bson.M
	if cur {
		filter = bson.M{"orderer": userName, "state": bson.M{"$ne": DeliverComplete}}
	} else {
		filter = bson.M{"orderer": userName, "state": DeliverComplete}
	}
	option := options.Find().SetProjection(bson.M{"orderer": 1, "state": 1, "businessname": 1, "menu": 1, "createdat": 1})
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

func (m *Model) ModifyOrder(orderID primitive.ObjectID, menu []MenuNum) bool {
	filter := bson.M{"_id": orderID}

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
	state := order.State
	switch state {
	case DeliverComplete, Delivering, ReceiptCancled, AdditionalReceiptCancled:
		return false
	case ReceiptCooking, AdditionalReceiptCooking:
		if !addition {
			return false
		}
	}
	update := bson.M{"$set": bson.M{"menu": menu, "state": AdditionalReceipting}}
	result, err := m.colOrder.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
	return true
}

func (m *Model) AdminListOrder(businessName string) []Order {
	filter := bson.M{"businessname": businessName}
	cursor, err := m.colOrder.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer cursor.Close(context.TODO())
	var order []Order
	if err = cursor.All(context.TODO(), &order); err != nil {
		panic(err)
	}
	return order
}

func (m *Model) UpdateOrderState(orderID primitive.ObjectID, state int) bool {
	filter := bson.M{"_id": orderID}
	update := bson.M{"$set": bson.M{"state": state}}

	result, err := m.colOrder.UpdateOne(context.TODO(), filter, update)
	if err == mongo.ErrNoDocuments {
		fmt.Println(err)
		return false
	} else if err != nil {
		panic(err)
	}
	fmt.Println(result)
	return true
}
