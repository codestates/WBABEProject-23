package model

import (
	"context"
	"fmt"
	"lecture/WBABEProject-23/logger"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"go.mongodb.org/mongo-driver/bson"
)

/*
New라는 네이밍은 들어가지 않아도 명확해 보입니다. CreateMenu로 변경하는건 어떨까요?
또한, 함수의 이름과는 달리 업데이트가 이루어 지고 있습니다.
*/
func (m *Model) CreateNewMenu(newMenu Menu, business string) int {
	objId, err := primitive.ObjectIDFromHex(business)
	if err != nil {
		panic(err)
	}
	filter := bson.M{"_id": objId}
	update := bson.M{"$push": bson.M{"menu": newMenu}}
	/*
	메뉴를 Create 하는 역할인데 Update를 하는 이유는 무엇인가요?
	생성후, 업데이트를 하는 것으로 보이는데 이렇게 구성하신 특별한 이유가 있을까요?
	*/
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

/*
업데이트를 하는 곳에서 UpdateOne이라는 함수를 이용하고 있으니, 통일성을 위해 UpdateMenu와 같은 네이밍은 어떠할까요?
*/
func (m *Model) ModifyMenu(toUpdate string, business string, menu Menu) {
	objId, err := primitive.ObjectIDFromHex(business)
	if err != nil {
		panic(err)
	}
	filter := bson.M{"_id": objId, "menu.name": toUpdate}
	update := bson.M{"$set": bson.M{}}
	/*
	1. Validator를 통해 Input 값에 대해 유효성 검살르 진행하시면 좋을 것 같습니다.
		들어오는 데이터에 대해서 검증하려면 Gin에서 제공하는 validtor 기능을 이용하시면 좋을 것 같습니다.
		required 필드같은 경우도 제어할 수 있습니다.
		아래의 링크를 참고해보시기 바랍니다.
		https://gin-gonic.com/docs/examples/custom-validators/
	*/
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
