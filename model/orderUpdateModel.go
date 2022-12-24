package model

import (
	"context"
	"fmt"
	"reflect"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (m *Model) ModifyOrder(orderID primitive.ObjectID, menu []MenuNum) bool {
	filter := bson.M{"_id": orderID}

	var order Order
	err := m.colOrder.FindOne(context.TODO(), filter).Decode(&order)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	addition := false
	if reflect.DeepEqual(order.Menu, menu) {
		addition = true
	}
	state := order.State
	switch state {
	case DeliverComplete, Delivering, ReceiptCancled, AdditionalReceiptCancled:
		return false
	case ReceiptCooking, AdditionalReceiptCooking:
		if !addition {
			return false
		}
	}
	update := bson.M{"$set": bson.M{"menu": menu, "state": AdditionalReceipting}}
	result, err := m.colOrder.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
	return true
}

func (m *Model) UpdateOrderState(orderID primitive.ObjectID, state int) bool {
	filter := bson.M{"_id": orderID}
	update := bson.M{"$set": bson.M{"state": state}}

	result, err := m.colOrder.UpdateOne(context.TODO(), filter, update)
	if err == mongo.ErrNoDocuments {
		fmt.Println(err)
		return false
	} else if err != nil {
		panic(err)
	}
	fmt.Println(result)
	return true
}
