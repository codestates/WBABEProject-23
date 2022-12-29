package model

import "go.mongodb.org/mongo-driver/bson/primitive"

// user type
const (
	Orderer  = 1
	Recipent = 2
)

type User struct {
	Id       primitive.ObjectID `bson:"_id"`
	Name     string             `bson:"name"`
	/*
	이 필드는 무슨 역할을 하게 되나요? 단어만으로는 어떤 필드인지 유추하기가 힘들어 보입니다.
	필드같은 경우에는 직관적인 단어를 사용하시는 것이 좋습니다.
	*/
	Pnum     string             `bson:"pnum"`
	UserType int                `bson:"userType"`
}
