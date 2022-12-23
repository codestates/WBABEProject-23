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
	Name      string  `bson:"name"`
	Status    int     `bson:"status"`
	Price     int     `bson:"price"`
	Origin    string  `bson:"origin"`
	Score     float32 `bson:"score"`
	IsDeleted bool    `bson:"is-deleted"`
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
	objId, err := primitive.ObjectIDFromHex(business)
	if err != nil {
		panic(err)
	}
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

func (m *Model) ModifyMenu(toUpdate string, business string, menu Menu) {
	objId, err := primitive.ObjectIDFromHex(business)
	if err != nil {
		panic(err)
	}
	filter := bson.M{"_id": objId, "menu.name": toUpdate}
	update := bson.M{"$set": bson.M{}}
	if menu.Name != "" {
		update["$set"].(bson.M)["menu.$.name"] = menu.Name
	}
	if menu.Status != 0 {
		update["$set"].(bson.M)["menu.$.status"] = menu.Status
	}
	if menu.Price != 0 {
		update["$set"].(bson.M)["menu.$.price"] = menu.Price
	}
	if menu.Origin != "" {
		update["$set"].(bson.M)["menu.$.origin"] = menu.Origin
	}
	if menu.Score != 0 {
		update["$set"].(bson.M)["menu.$.score"] = menu.Score
	}
	if menu.IsDeleted {
		update["$set"].(bson.M)["menu.$.is-deleted"] = menu.IsDeleted
	}
	result, err := m.colBusiness.UpdateOne(context.TODO(), filter, update)
	if err == mongo.ErrNoDocuments {
		fmt.Println("No document was found with the business id or menu name")
		return
	} else if err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Println(result.MatchedCount, result.ModifiedCount)
}
