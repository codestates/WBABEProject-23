package model

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	Receipting = iota + 1
	ReceiptCancled
	ReceiptComplete
	AdditionalReceipting
	AdditionalReceiptingComplete
	AdditionalReceiptCancled
	AdditionalReceiptCooking
	ReceiptCooking
	Delivering
	DeliverComplete
)

type Order struct {
	Id           primitive.ObjectID `bson:"_id"`
	OrderID      int64              `bson:"orderid`
	Orderer      string             `bson:"orderer"`
	State        int                `bson:"state"`
	BusinessName string             `bson:"businessname`
	Menu         []MenuNum          `bson:"menu`
	/* 
	CreatedAt 이외에 UpdatedAt도 같이 넣어주신다면 추후 히스토리 추적에 용이합니다. 값이 언제 업데이트 되었는지 확인할 수 있으니까요.
	*/
	CreatedAt    time.Time          `bson:"createdat`
}

type MenuNum struct {
	MenuName   string `bson:"menuname"`
	Number     int    `bson:"number"`
	IsReviewed bool   `bson:"isreviewed"`
}

func (m *Model) MakeOrder(order Order) {
	/*
	오더를 만드는 것과 큰 연관이 없는 아래와 같은 로직들(현재 시간 가져오기)은 유틸성 함수에 속합니다.
	따라서 유틸 함수를 제작해서 여러 곳에서 호출하며 재사용 가능하도록 구성하면 좋습니다.
	*/
	now := time.Now().UTC()

	// Set the start and end of the range to the start and end of the desired day
	start := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)
	end := start.AddDate(0, 0, 1)
	filter := bson.M{"createdat": bson.M{
		"$gte": start,
		"$lt":  end,
	}}
	count, err := m.colOrder.CountDocuments(context.TODO(), filter)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	order.OrderID = count + 1
	fmt.Println(order)
	result, err := m.colOrder.InsertOne(context.TODO(), order)
	if err != nil {
		panic(err)
	}
	fmt.Println(result.InsertedID)
}
