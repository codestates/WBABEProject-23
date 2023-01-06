package entitiy

import "go.mongodb.org/mongo-driver/bson/primitive"

// user type
const (
	Orderer  = 1
	Recipent = 2
)

type User struct {
	Id       primitive.ObjectID `bson:"_id"`
	Name     string             `bson:"name"`
	Pnum     string             `bson:"pnum"`
	UserType int                `bson:"userType"`
}
