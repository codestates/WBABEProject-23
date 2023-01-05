package model

import (
	"context"
	"lecture/WBABEProject-23/protocol"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (m *Model) ListOrder(userName string, cur bool) *protocol.ApiResponse[any] {
	var filter bson.M
	if cur {
		filter = bson.M{"orderer": userName, "state": bson.M{"$ne": DeliverComplete}}
	} else {
		filter = bson.M{"orderer": userName, "state": DeliverComplete}
	}
	option := options.Find().SetProjection(bson.M{"orderer": 1, "state": 1, "businessname": 1, "menu": 1, "createdat": 1})
	cursor, err := m.colOrder.Find(context.TODO(), filter, option)
	if err != nil {
		return protocol.Fail(err, protocol.InternalServerError)
	}
	defer cursor.Close(context.TODO())

	var orders []Order
	if err = cursor.All(context.TODO(), &orders); err != nil {
		return protocol.Fail(err, protocol.InternalServerError)
	}
	return protocol.SuccessData(orders, protocol.OK)
}

func (m *Model) AdminListOrder(id primitive.ObjectID) *protocol.ApiResponse[any] {
	filter := bson.M{"business_id": id}
	cursor, err := m.colOrder.Find(context.TODO(), filter)
	if err != nil {
		return protocol.Fail(err, protocol.InternalServerError)
	}
	defer cursor.Close(context.TODO())
	var order []Order
	if err = cursor.All(context.TODO(), &order); err != nil {
		return protocol.Fail(err, protocol.InternalServerError)
	}
	return protocol.SuccessData(order, protocol.OK)
}
