package model

import (
	"context"
	"lecture/WBABEProject-23/protocol"

	"go.mongodb.org/mongo-driver/bson"
)

func (m *Model) CreateMenu(newMenu Menu) (res *protocol.ApiResponse[any]) {
	_, err := m.colMenu.InsertOne(context.TODO(), newMenu)
	if err != nil {
		return protocol.Fail(err, protocol.InternalServerError)
	} else {
		return protocol.Success(protocol.Created)
	}
}

func (m *Model) UpdateMenu(menu *Menu) *protocol.ApiResponse[any] {

	filter := bson.M{"_id": menu.ID}
	update := bson.M{"$set": menu}

	_, err := m.colMenu.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return protocol.Fail(err, protocol.InternalServerError)
	}
	return protocol.Success(protocol.Created)
}
