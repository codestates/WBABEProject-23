package model

import (
	"context"
	"encoding/json"
	"fmt"
	"lecture/WBABEProject-23/protocol"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"

	"go.mongodb.org/mongo-driver/bson"
)

func (m *Model) ReadMenu(id primitive.ObjectID, sortBy string, sortOrder int) *protocol.ApiResponse[any] {
	filter := bson.M{"business_id": bson.M{"$ref": "business", "$id": id}, "is_deleted": false}
	option := options.Find().SetSort(bson.M{sortBy: sortOrder}).SetProjection(bson.M{"name": 1, "price": 1, "origin": 1, "score": 1, "category": 1})
	cursor, err := m.colMenu.Find(context.TODO(), filter, option)
	if err != nil {
		return protocol.Fail(err, protocol.InternalServerError)
	}
	defer cursor.Close(context.TODO())
	var results []bson.M
	if err = cursor.All(context.TODO(), &results); err != nil {
		return protocol.Fail(err, protocol.InternalServerError)
	}
	return protocol.SuccessData(results, protocol.OK)
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
