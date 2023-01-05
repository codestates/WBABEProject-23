package model

import (
	"context"
	"fmt"
	"lecture/WBABEProject-23/protocol"
	"reflect"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (m *Model) UpdateOrder(orderID primitive.ObjectID, menu []MenuNum) *protocol.ApiResponse[any] {
	filter := bson.M{"_id": orderID}

	var order Order
	err := m.colOrder.FindOne(context.TODO(), filter).Decode(&order)
	if err != nil {
		return protocol.Fail(err, protocol.InternalServerError)
	}
	addition := false
	if reflect.DeepEqual(order.Menu, menu) {
		addition = true
	}
	state := order.State
	switch state {
	case DeliverComplete, Delivering, ReceiptCancled, AdditionalReceiptCancled:
		return protocol.FailCustomMessage(nil, "The state is not updatable", protocol.BadRequest)
	case ReceiptCooking, AdditionalReceiptCooking:
		if !addition {
			return protocol.FailCustomMessage(nil, "The state is not updatable", protocol.BadRequest)
		}
	}
	update := bson.M{"$set": bson.M{"menu": menu, "state": AdditionalReceipting, "updated_at": time.Now()}}
	result, err := m.colOrder.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return protocol.Fail(err, protocol.InternalServerError)
	}
	fmt.Println(result)
	return protocol.Success(protocol.OK)
}

func (m *Model) UpdateOrderState(orderID primitive.ObjectID, state int) *protocol.ApiResponse[any] {
	filter := bson.M{"_id": orderID}
	update := bson.M{"$set": bson.M{"state": state}}

	result, err := m.colOrder.UpdateOne(context.TODO(), filter, update)
	if err == mongo.ErrNoDocuments {
		return protocol.Fail(err, protocol.BadRequest)
	} else if err != nil {
		return protocol.Fail(err, protocol.InternalServerError)
	}
	fmt.Println(result)
	return protocol.Success(protocol.OK)
}
