package model

import (
	"context"
	"lecture/WBABEProject-23/protocol"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
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
}

type MenuNum struct {
	MenuID     primitive.ObjectID `bson:"menu_id"`
	Number     int                `bson:"number"`
	IsReviewed bool               `bson:"is_reviewed"`
}

func (m *Model) CheckOrderByID(id primitive.ObjectID) (bool, error) {
	filter := bson.M{"_id": id}
	var result bson.M
	err := m.colOrder.FindOne(context.TODO(), filter).Decode(&result)
	if err == mongo.ErrNoDocuments {
		return false, nil
	} else if err != nil {
		return false, err
	} else {
		return true, nil
	}
}

func (m *Model) CreateOrder(order *Order) *protocol.ApiResponse[any] {
	now := time.Now().UTC()

	start := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)
	end := start.AddDate(0, 0, 1)
	filter := bson.M{"createdat": bson.M{
		"$gte": start,
		"$lt":  end,
	}}
	count, err := m.colOrder.CountDocuments(context.TODO(), filter)
	if err != nil {
		return protocol.Fail(err, protocol.InternalServerError)
	}
	order.OrderID = count + 1
	_, err = m.colOrder.InsertOne(context.TODO(), order)
	if err != nil {
		return protocol.Fail(err, protocol.InternalServerError)
	}
	return protocol.Success(protocol.Created)
}
