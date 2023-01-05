package model

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Business struct {
	ID    primitive.ObjectID `bson:"_id,omitempty"`
	Name  string             `bson:"name,omitempty"`
	Admin primitive.ObjectID `bson:"admin,omitempty"`
	Menu  []Menu             `bson:"menu,omitempty"`
}

func (m *Model) CheckBusinessID(id primitive.ObjectID) (bool, error) {
	filter := bson.M{"_id": id}
	var result bson.M
	err := m.colBusiness.FindOne(context.TODO(), filter).Decode(&result)
	if err == mongo.ErrNoDocuments {
		return false, nil
	} else if err != nil {
		return false, err
	} else {
		return true, nil
	}
}
