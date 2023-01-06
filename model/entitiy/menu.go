package entitiy

import "go.mongodb.org/mongo-driver/bson/primitive"

// menu state
const (
	Ready    = 1
	NotReady = 2
)

type Menu struct {
	ID         primitive.ObjectID `bson:"_id"`
	Name       string             `bson:"name,omitempty"`
	State      int                `bson:"state,omitempty"`
	Price      int                `bson:"price,omitempty"`
	Origin     string             `bson:"origin,omitempty"`
	Score      float32            `bson:"score,omitempty"`
	IsDeleted  bool               `bson:"is_deleted,omitempty"`
	Category   string             `bson:"category,omitempty"`
	BusinessID primitive.M        `bson:"business_id,omitempty"`
}
