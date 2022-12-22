package model

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Business struct {
	Id     primitive.ObjectID `bson:"_id"`
	Name   string             `bson:"name"`
	Admin  User               `bson:"admin"`
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
	Orderer  User   `bson:"orderer"`
	MenuName string `bson:"menu-name"`
	Content  string `bson:"content"`
	Score    int    `bson:"score"`
}

func (m *Model) CreateNewMenu(newMenu Menu, business string) {
	data, _ := bson.Marshal(newMenu)
	filter := bson.M{"_id": business}
	update := bson.M{"$push": data}
	result, err := m.colBusiness.UpdateOne(context.TODO(), filter, update)
	if err == mongo.ErrNoDocuments {
		fmt.Println("No document was found with the pnum")
		return
	} else if err != nil {
		panic(err)
	}
	fmt.Println(result)
}
