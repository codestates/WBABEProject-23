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
	/* 
	주문을 수정할 수 없는 상태를 const를 사용하니 가독성에 좋습니다. 좋은 코드네요.
	Enum을 활용한다면 조금 더 용이할 수 있습니다. 참고 바랍니다.
	*/
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
