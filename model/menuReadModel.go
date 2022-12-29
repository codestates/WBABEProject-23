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
			/*
			데이터베이스가 아니라, 결과를 가져온 후 golang에서 Sort 하는 이유는 무엇인가요? 
			데이터의 양이 많아질수록 언어 내에서 정렬을 하는데는 오랜 시간이 걸립니다.
			MongoDB에서 지원하는 Sort 기능을 활용하시는 것을 추천 드립니다.
			
			참고 링크입니다.
			https://www.mongodb.com/docs/drivers/go/current/fundamentals/crud/read-operations/sort/
			*/
			sort.Slice(menu, func(i, j int) bool {
				value1 := reflect.ValueOf(menu[i]).FieldByName(strings.Title(sortBy)).Interface()
				value2 := reflect.ValueOf(menu[j]).FieldByName(strings.Title(sortBy)).Interface()
				/*
				형 별로 나누어서 처리하는 이유는 무엇인가요?
				*/
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

/*
Menu 라는 네이밍이 들어가지 않아도 괜찮을 것 같습니다.
Review라는 네이밍만으로도 충분해 보입니다. Review라는 struct 안에도 Menu에 대한 정보가 들어있으니까요.
*/
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
	reviewOption := options.Find().SetProjection(bson.M{"orderer": 1, "menuname": 1, "content": 1, "score": 1})
	cursor, err := m.colReview.Find(context.TODO(), filter, reviewOption)
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
