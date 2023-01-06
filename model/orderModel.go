package model

import (
	"context"
	"lecture/WBABEProject-23/model/entitiy"
	"lecture/WBABEProject-23/protocol"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

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

func (m *Model) CheckOrderReviewable(review *entitiy.Review) *protocol.ApiResponse[any] {
	filter := bson.M{"_id": review.OrderID, "orderer": review.Orderer, "state": entitiy.DeliverComplete}
	projection := bson.M{"menu": bson.M{"$elemMatch": bson.M{"menu_id": review.MenuID, "is_reviewed": false}}}
	findOption := options.FindOne().SetProjection(projection)
	isIn := m.colOrder.FindOne(context.TODO(), filter, findOption)
	if isIn.Err() == mongo.ErrNoDocuments {
		return protocol.Fail(isIn.Err(), protocol.BadRequest)
	} else if isIn.Err() != nil {
		return protocol.Fail(isIn.Err(), protocol.InternalServerError)
	} else {
		return nil
	}
}

func (m *Model) CreateOrder(order *entitiy.Order) *protocol.ApiResponse[any] {
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
