package model

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// menu state
const (
	Ready    = 1
	NotReady = 2
)

type Menu struct {
	ID         primitive.ObjectID `bson:"_id"`
	Name       string             `bson:"name,omitempty"`
	State      int                `bson:"state,omitempty"`
	Price      int                `bson:"price,omitempty"`
	Origin     string             `bson:"origin,omitempty"`
	Score      float32            `bson:"score,omitempty"`
	IsDeleted  bool               `bson:"is_deleted,omitempty"`
	Category   string             `bson:"category,omitempty"`
	BusinessID primitive.M        `bson:"business_id,omitempty"`
}

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

func (m *Model) GetMenuById(id primitive.ObjectID) (*Menu, error) {
	filter := bson.M{"_id": id, "is_deleted": false}
	var result Menu
	err := m.colMenu.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return nil, err
	} else {
		return &result, nil
	}
}
