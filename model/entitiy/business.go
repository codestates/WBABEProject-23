package entitiy

import "go.mongodb.org/mongo-driver/bson/primitive"

type Business struct {
	ID    primitive.ObjectID `bson:"_id,omitempty"`
	Name  string             `bson:"name,omitempty"`
	Admin primitive.ObjectID `bson:"admin,omitempty"`
	Menu  []Menu             `bson:"menu,omitempty"`
}
