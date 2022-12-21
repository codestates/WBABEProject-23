package model

import "go.mongodb.org/mongo-driver/mongo"

type Model struct {
	client *mongo.Client
}
