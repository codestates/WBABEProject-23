package model

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (m *Model) ListOrder(userName string, cur bool) []Order {
	var filter bson.M
	if cur {
		filter = bson.M{"orderer": userName, "state": bson.M{"$ne": DeliverComplete}}
	} else {
		filter = bson.M{"orderer": userName, "state": DeliverComplete}
	}
	option := options.Find().SetProjection(bson.M{"orderer": 1, "state": 1, "businessname": 1, "menu": 1, "createdat": 1})
	cursor, err := m.colOrder.Find(context.TODO(), filter, option)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer cursor.Close(context.TODO())

	var orders []Order
	if err = cursor.All(context.TODO(), &orders); err != nil {
		fmt.Println(err)
		panic(err)
	}
	return orders
}

func (m *Model) AdminListOrder(businessName string) []Order {
	filter := bson.M{"businessname": businessName}
	cursor, err := m.colOrder.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer cursor.Close(context.TODO())
	var order []Order
	if err = cursor.All(context.TODO(), &order); err != nil {
		panic(err)
	}
	return order
}
