package model

import (
	"context"
	"fmt"
	"reflect"
	"sort"
	"strings"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"go.mongodb.org/mongo-driver/bson"
)

type Business struct {
	Id    primitive.ObjectID `bson:"_id"`
	Name  string             `bson:"name"`
	Admin primitive.ObjectID `bson:"admin"`
	Menu  []Menu             `bson:"menu"`
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
	IsDeleted bool    `bson:"is_deleted"`
	Category  string  `bson:"category"`
}

func (m *Model) CreateNewMenu(newMenu Menu, business string) {
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

func (m *Model) ListMenu(toList string, sortBy string, sortOrder int) []Menu {
	var result Business
	filter := bson.M{"name": toList}
	option := options.FindOne().SetProjection(bson.M{"menu": 1})

	if err := m.colBusiness.FindOne(context.TODO(), filter, option).Decode(&result); err != nil {
		fmt.Println(err)
		panic(err)
	}
	menu := result.Menu
	if sortBy != "" {
		if sortOrder == 1 {
			sort.Slice(menu, func(i, j int) bool {
				value1 := reflect.ValueOf(menu[i]).FieldByName(strings.Title(sortBy)).Interface()
				value2 := reflect.ValueOf(menu[j]).FieldByName(strings.Title(sortBy)).Interface()
				switch value1.(type) {
				case int:
					return value1.(int) < value2.(int)
				case float64:
					return value1.(float64) < value2.(float64)
				case string:
					return value1.(string) < value2.(string)
				default:
					// handle other types as needed
					return false
				}
			})
		} else {
			sort.Slice(menu, func(i, j int) bool {
				value1 := reflect.ValueOf(menu[i]).FieldByName(strings.Title(sortBy)).Interface()
				value2 := reflect.ValueOf(menu[j]).FieldByName(strings.Title(sortBy)).Interface()
				switch value1.(type) {
				case int:
					return value1.(int) > value2.(int)
				case float64:
					return value1.(float64) > value2.(float64)
				case string:
					return value1.(string) > value2.(string)
				default:
					// handle other types as needed
					return false
				}
			})
		}
	}
	return menu
}

func (m *Model) ReadMenuReview(toRead primitive.ObjectID, menuName string) map[string]interface{} {
	var result Business
	option := options.FindOne().SetProjection(bson.M{"menu": 1})
	if err := m.colBusiness.FindOne(context.TODO(), bson.M{"_id": toRead}, option).Decode(&result); err == mongo.ErrNoDocuments {
		fmt.Println("no document")
		return map[string]interface{}{}
	} else if err != nil {
		fmt.Println(err)
		panic(err)
	}
	var score float32
	for _, res := range result.Menu {
		if res.Name == menuName {
			score = res.Score
			break
		}
	}

	filter := bson.M{"businessid": toRead, "menuname": menuName}
	cursor, err := m.colReview.Find(context.TODO(), filter)
	if err != nil {
		panic(err)
	}
	defer cursor.Close(context.TODO())
	var review []Review
	if err = cursor.All(context.TODO(), &review); err != nil {
		panic(err)
	}
	val := map[string]interface{}{
		"score":  score,
		"review": review,
	}
	return val
}
