package model

import (
	"context"
	"fmt"
	"lecture/WBABEProject-23/logger"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"go.mongodb.org/mongo-driver/bson"
)

func (m *Model) CreateNewMenu(newMenu Menu, business string) int {
	objId, err := primitive.ObjectIDFromHex(business)
	if err != nil {
		panic(err)
	}
	filter := bson.M{"_id": objId}
	update := bson.M{"$push": bson.M{"menu": newMenu}}
	result, err := m.colBusiness.UpdateOne(context.TODO(), filter, update)
	if err == mongo.ErrNoDocuments {
		logger.Warn("No document was found with the business id")
		return 1
	} else if err != nil {
		logger.Error(err)
		return 2
	}
	fmt.Println(result)
	return 0
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
	if menu.State != 0 {
		update["$set"].(bson.M)["menu.$.state"] = menu.State
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
		update["$set"].(bson.M)["menu.$.is_deleted"] = menu.IsDeleted
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
