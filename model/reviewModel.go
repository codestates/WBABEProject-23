package model

import (
	"context"
	"fmt"
	"lecture/WBABEProject-23/protocol"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Review struct {
	OrderID primitive.ObjectID `bson:"order_id" json:"order_id"`
	MenuID  primitive.ObjectID `bson:"menu_id" json:"menu_id"`
	Orderer string             `bson:"orderer"`
	Content string             `bson:"content"`
	Score   float32            `bson:"score"`
}

func (m *Model) CreateReview(review *Review) *protocol.ApiResponse[any] {
	orderFilter := bson.M{"_id": review.OrderID, "orderer": review.Orderer, "state": DeliverComplete}
	orderProjection := bson.M{"menu": bson.M{"$elemMatch": bson.M{"menu_id": review.MenuID, "is_reviewed": false}}}
	orderFindOption := options.FindOne().SetProjection(orderProjection)
	isIn := m.colOrder.FindOne(context.TODO(), orderFilter, orderFindOption)
	if isIn.Err() == mongo.ErrNoDocuments {
		return protocol.Fail(isIn.Err(), protocol.BadRequest)
	} else if isIn.Err() != nil {
		return protocol.Fail(isIn.Err(), protocol.InternalServerError)
	}
	result, err := m.colReview.InsertOne(context.TODO(), review)
	if err != nil {
		return protocol.Fail(err, protocol.InternalServerError)
	}
	orderUpdate := bson.M{"$set": bson.M{
		"menu.$[i].is_reviewed": true,
		"menu.$[i].review":      bson.M{"$ref": "review", "$id": result.InsertedID},
	}}
	orderArrayFilters := options.ArrayFilters{
		Filters: []interface{}{bson.M{"i.menu_id": review.MenuID}},
	}
	orderUpdateOption := options.Update().SetArrayFilters(orderArrayFilters)
	orderUpdateResult, err := m.colOrder.UpdateOne(context.TODO(), orderFilter, orderUpdate, orderUpdateOption)
	if err != nil {
		return protocol.Fail(err, protocol.InternalServerError)
	}
	fmt.Println(orderUpdateResult)
	//////////////////////////////////////////////////////////////////////////////////////
	pipeline := []bson.M{
		{
			"$match": bson.M{
				"menu_id": review.MenuID,
			},
		},
		{
			"$group": bson.M{
				"_id": bson.M{
					"menu_id": "$menu_id",
				},
				"totalScore": bson.M{
					"$sum": "$score",
				},
				"count": bson.M{
					"$sum": 1,
				},
			},
		},
	}
	cursor, err := m.colReview.Aggregate(context.TODO(), pipeline)
	if err != nil {
		return protocol.Fail(err, protocol.InternalServerError)
	}
	var sum struct {
		TotalScore float32 `bson:"totalScore"`
		Count      int     `bson:"count`
	}
	if cursor.Next(context.TODO()) {
		if err := cursor.Decode(&sum); err != nil {
			return protocol.Fail(err, protocol.InternalServerError)
		}
	} else {
		fmt.Println("No documents found")
	}
	avg := sum.TotalScore / float32(sum.Count)
	//////////////////////////////////////////////////////////////////////////
	menuFilter := bson.M{"_id": review.MenuID}
	menuUpdate := bson.M{"$set": bson.M{"score": avg}}
	menuUpdateResult, err := m.colMenu.UpdateOne(context.TODO(), menuFilter, menuUpdate)
	if err != nil {
		return protocol.Fail(err, protocol.InternalServerError)
	}
	fmt.Println(menuUpdateResult)
	return protocol.Success(protocol.Created)
}
