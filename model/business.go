package model

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"go.mongodb.org/mongo-driver/bson"
)

type Business struct {
	Id     primitive.ObjectID `bson:"_id"`
	Name   string             `bson:"name"`
	Admin  primitive.ObjectID `bson:"admin"`
	Menu   []Menu             `bson:"menu"`
	Review []Review           `bson:"review"`
}

// menu status
const (
	Ready    = 1
	NotReady = 2
)

type Menu struct {
	Name   string  `bson:"name"`
	Status int     `bson:"status"`
	Price  int     `bson:"price"`
	Origin string  `bson:"origin"`
	Score  float32 `bson:"score"`
}

type Review struct {
	Orderer  primitive.ObjectID `bson:"orderer"`
	MenuName string             `bson:"menu-name"`
	Content  string             `bson:"content"`
	Score    int                `bson:"score"`
}

func (m *Model) CreateNewMenu(newMenu Menu, business string) {
	// data, err := bson.Marshal(newMenu)
	// if err != nil {
	// 	fmt.Println(err)
	// 	panic(err)
	// }
	objId, _ := primitive.ObjectIDFromHex(business)
	filter := bson.M{"_id": objId}
	update := bson.M{"$push": bson.M{"menu": newMenu}}
	result, err := m.colBusiness.UpdateOne(context.TODO(), filter, update)
	if err == mongo.ErrNoDocuments {
		fmt.Println("No document was found with the business id")
		return
	} else if err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Println(result)

	// cursor, err := m.colBusiness.Find(context.TODO(), filter)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer cursor.Close(context.TODO())

	// // Iterate through the cursor and print the documents
	// var doc bson.M
	// for cursor.Next(context.TODO()) {
	// 	err := cursor.Decode(&doc)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	fmt.Println(doc)
	// }
}
