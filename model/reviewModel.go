package model

import (
	"context"
	"encoding/json"
	"fmt"
	"lecture/WBABEProject-23/model/entitiy"
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

func (m *Model) CreateReview(review *entitiy.Review) *protocol.ApiResponse[any] {
	result, err := m.colReview.InsertOne(context.TODO(), review)
	if err != nil {
		return protocol.Fail(err, protocol.InternalServerError)
	}
	if res := m.updateOrderReviewd(review, result); res != nil {
		return res
	}
	avg, res := m.calAvgScore(review)
	if res != nil {
		return res
	}
	res = m.menuReviewUpdate(review, avg)
	return res
}
func (m *Model) updateOrderReviewd(review *entitiy.Review, result *mongo.InsertOneResult) *protocol.ApiResponse[any] {
	orderFilter := bson.M{"_id": review.OrderID, "orderer": review.Orderer, "state": entitiy.DeliverComplete}
	orderUpdate := bson.M{"$set": bson.M{
		"menu.$[i].is_reviewed": true,
		"menu.$[i].review":      bson.M{"$ref": "review", "$id": result.InsertedID},
	}}
	orderArrayFilters := options.ArrayFilters{
		Filters: []interface{}{bson.M{"i.menu_id": review.MenuID}},
	}
	orderUpdateOption := options.Update().SetArrayFilters(orderArrayFilters)
	_, err := m.colOrder.UpdateOne(context.TODO(), orderFilter, orderUpdate, orderUpdateOption)
	if err != nil {
		return protocol.Fail(err, protocol.InternalServerError)
	}
	return nil
}

func (m *Model) calAvgScore(review *entitiy.Review) (float32, *protocol.ApiResponse[any]) {
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
		return 0, protocol.Fail(err, protocol.InternalServerError)
	}
	var sum struct {
		TotalScore float32 `bson:"totalScore"`
		Count      int     `bson:"count"`
	}
	if cursor.Next(context.TODO()) {
		if err := cursor.Decode(&sum); err != nil {
			return 0, protocol.Fail(err, protocol.InternalServerError)
		}
	} else {
		fmt.Println("No documents found")
	}
	return sum.TotalScore / float32(sum.Count), nil
}

func (m *Model) menuReviewUpdate(review *entitiy.Review, avg float32) *protocol.ApiResponse[any] {
	menuFilter := bson.M{"_id": review.MenuID}
	menuUpdate := bson.M{"$set": bson.M{"score": avg}}
	_, err := m.colMenu.UpdateOne(context.TODO(), menuFilter, menuUpdate)
	if err != nil {
		return protocol.Fail(err, protocol.InternalServerError)
	}
	return protocol.Success(protocol.Created)
}

func (m *Model) ReadReview(toRead primitive.ObjectID) *protocol.ApiResponse[any] {
	pipeline := []bson.M{
		{"$match": bson.M{"menu_id": toRead}},
		{"$lookup": bson.M{
			"from":         "menu",
			"localField":   "menu_id",
			"foreignField": "_id",
			"as":           "menu_name",
		}},
		{"$unwind": "$menu_name"},
	}
	cursor, err := m.colReview.Aggregate(context.TODO(), pipeline)
	if err != nil {
		return protocol.Fail(err, protocol.InternalServerError)
	}
	defer cursor.Close(context.TODO())
	results := []bson.M{}
	err = cursor.All(context.TODO(), &results)
	if err != nil {
		return protocol.Fail(err, protocol.InternalServerError)
	}
	for _, result := range results {
		res, _ := json.Marshal(result)
		fmt.Println(res)
	}
	return protocol.SuccessData(results, protocol.OK)
}
