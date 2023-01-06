package model

import (
	"context"
	"lecture/WBABEProject-23/model/entitiy"
	"lecture/WBABEProject-23/protocol"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (m *Model) CheckMenuID(id primitive.ObjectID) (bool, error) {
	filter := bson.M{"_id": id}
	var result bson.M
	err := m.colMenu.FindOne(context.TODO(), filter).Decode(&result)
	if err == mongo.ErrNoDocuments {
		return false, nil
	} else if err != nil {
		return false, err
	} else {
		return true, nil
	}
}

func (m *Model) CheckMenuFieldExists(field string) (bool, error) {
	filter := bson.M{field: bson.M{"$exists": true}}
	var result bson.M
	err := m.colMenu.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return false, err
	} else if len(result) == 0 {
		return false, nil
	} else {
		return true, nil
	}
}

func (m *Model) GetMenuById(id primitive.ObjectID) (*entitiy.Menu, error) {
	filter := bson.M{"_id": id, "is_deleted": false}
	var result entitiy.Menu
	err := m.colMenu.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return nil, err
	} else {
		return &result, nil
	}
}

func (m *Model) CreateMenu(newMenu entitiy.Menu) (res *protocol.ApiResponse[any]) {
	_, err := m.colMenu.InsertOne(context.TODO(), newMenu)
	if err != nil {
		return protocol.Fail(err, protocol.InternalServerError)
	} else {
		return protocol.Success(protocol.Created)
	}
}

func (m *Model) UpdateMenu(menu *entitiy.Menu) *protocol.ApiResponse[any] {

	filter := bson.M{"_id": menu.ID}
	update := bson.M{"$set": menu}

	_, err := m.colMenu.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return protocol.Fail(err, protocol.InternalServerError)
	}
	return protocol.Success(protocol.Created)
}

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
