package model

import (
	"context"
	"fmt"
	"lecture/WBABEProject-23/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Model struct {
	client      *mongo.Client
	colBusiness *mongo.Collection
	colUser     *mongo.Collection
	colOrder    *mongo.Collection
	colReview   *mongo.Collection
}

func NewModel(config *config.Config) (*Model, error) {
	r := &Model{}

	var err error
	mgUrl := fmt.Sprintf("%v", config.DB["admin"]["host"])
	credential := options.Credential{
		AuthSource: "admin",
		Username:   fmt.Sprintf("%v", config.DB["admin"]["user"]),
		Password:   fmt.Sprintf("%v", config.DB["admin"]["pass"]),
	}
	if r.client, err = mongo.Connect(context.Background(), options.Client().ApplyURI(mgUrl).SetAuth(credential)); err != nil {
		return nil, err
	} else if err := r.client.Ping(context.Background(), nil); err != nil {
		return nil, err
	} else {
		db := r.client.Database(fmt.Sprintf("%v", config.DB["admin"]["name"]))
		r.colBusiness = db.Collection("business")
		r.colUser = db.Collection("user")
		r.colOrder = db.Collection("order")
		r.colReview = db.Collection("review")
	}
	return r, nil
}
