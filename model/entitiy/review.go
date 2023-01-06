package entitiy

import "go.mongodb.org/mongo-driver/bson/primitive"

type Review struct {
	OrderID primitive.ObjectID `bson:"order_id" json:"order_id"`
	MenuID  primitive.ObjectID `bson:"menu_id" json:"menu_id"`
	Orderer string             `bson:"orderer"`
	Content string             `bson:"content"`
	Score   float32            `bson:"score"`
}
