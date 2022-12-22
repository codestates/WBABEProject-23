package model

// user type
const (
	Orderer  = 1
	Recipent = 2
)

type User struct {
	Name     string `bson:"name"`
	Pnum     string `bson:"pnum"`
	UserType int    `bson:"userType"`
}
