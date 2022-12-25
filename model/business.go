package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Business struct {
	Id    primitive.ObjectID `bson:"_id"`
	Name  string             `bson:"name"`
	Admin primitive.ObjectID `bson:"admin"`
	Menu  []Menu             `bson:"menu"`
}

// menu state
const (
	Ready    = 1
	NotReady = 2
)

type Menu struct {
	Name      string  `bson:"name"`
	State     int     `bson:"state"`
	Price     int     `bson:"price"`
	Origin    string  `bson:"origin"`
	Score     float32 `bson:"score"`
	IsDeleted bool    `bson:"is_deleted"`
	Category  string  `bson:"category"`
}
