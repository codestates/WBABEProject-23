package model

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Review struct {
	OrderID    primitive.ObjectID `bson:"orderid`
	BusinessID primitive.ObjectID `bson:"businessid`
	Orderer    string             `bson:"orderer"`
	MenuName   string             `bson:"menuname"`
	Content    string             `bson:"content"`
	Score      float32            `bson:"score"`
}

/*
1. WriteReview 보다는 CreateReview라는 네이밍이 적절해 보입니다. 특별한 케이스가 아닌이상 CRUD 기능에 대한 네이밍은 통일을 해주시는 것이 가독성이 좋습니다.

2. 함수 내에서 많은 작업들을 처리하고 있습니다. 리뷰를 생성(작성)하는 것과 연관이 없는 로직들은 따로 함수로 뺴내어 분리해보는 것은 어떨까요? 이렇게 로직을 분리하다보면 공통적으로 사용하는 유틸 함수를 만들 수 있는 포인트가 되기도 하고, 재사용 또한 가능해집니다.

*/
func (m *Model) WriteReview(review Review) int {
	orderFilter := bson.M{"_id": review.OrderID, "orderer": review.Orderer, "state": DeliverComplete}
	orderProjection := bson.M{"menu": bson.M{"$elemMatch": bson.M{"menuname": review.MenuName, "isreviewed": false}}}
	orderFindOption := options.FindOne().SetProjection(orderProjection)
	isIn := m.colOrder.FindOne(context.TODO(), orderFilter, orderFindOption)
	if isIn.Err() == mongo.ErrNoDocuments {
		return 0
	} else if isIn.Err() != nil {
		panic(isIn.Err())
	}
	result, err := m.colReview.InsertOne(context.TODO(), review)
	if err != nil {
		panic(err)
	}
	orderUpdate := bson.M{"$set": bson.M{
		"menu.$[i].isreviewed": true,
		"menu.$[i].review":     bson.M{"$ref": "review", "$id": result.InsertedID},
	}}
	orderArrayFilters := options.ArrayFilters{
		Filters: []interface{}{bson.M{"i.menuname": review.MenuName}},
	}
	orderUpdateOption := options.Update().SetArrayFilters(orderArrayFilters)
	orderUpdateResult, err := m.colOrder.UpdateOne(context.TODO(), orderFilter, orderUpdate, orderUpdateOption)
	if err != nil {
		panic(err)
	}
	fmt.Println(orderUpdateResult)
	/*
	의미 없는 주석은 지워주시는 편이 좋습니다.
	*/
	//////////////////////////////////////////////////////////////////////////////////////
	pipeline := []bson.M{
		{
			"$match": bson.M{
				"businessid": review.BusinessID,
				"menuname":   review.MenuName,
			},
		},
		{
			"$group": bson.M{
				"_id": bson.M{
					"businessid": "$businessid",
					"menuname":   "$menuname",
				},
				"totalScore": bson.M{
					"$sum": "$score",
				},
				"count": bson.M{
					"$sum": 1,
				},
			},
		},
	}
	cursor, err := m.colReview.Aggregate(context.TODO(), pipeline)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	var sum struct {
		TotalScore float32 `bson:"totalScore"`
		Count      int     `bson:"count`
	}
	if cursor.Next(context.TODO()) {
		if err := cursor.Decode(&sum); err != nil {
			fmt.Println(err)
			panic(err)
		}
	} else {
		fmt.Println("No documents found")
	}
	avg := sum.TotalScore / float32(sum.Count)
	//////////////////////////////////////////////////////////////////////////
	businessFilter := bson.M{"_id": review.BusinessID}
	businessArrayFilters := options.ArrayFilters{
		Filters: []interface{}{bson.M{"i.name": review.MenuName}},
	}
	businessOption := options.Update().SetArrayFilters(businessArrayFilters)
	businessUpdate := bson.M{"$set": bson.M{
		"menu.$[i].score": avg,
	}}
	businessUpdateResult, err := m.colBusiness.UpdateOne(context.TODO(), businessFilter, businessUpdate, businessOption)
	if err != nil {
		panic(err)
	}
	fmt.Println(businessUpdateResult)
	return 1
}
