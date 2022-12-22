package model

import (
	"context"
	"fmt"
	"lecture/WBABEProject-23/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Model struct {
	client *mongo.Client
}

func NewModel(config *config.Config) (*Model, error) {
	r := &Model{}

	var err error
	mgUrl := fmt.Sprintf("%v", config.DB["admin"]["host"])
	if r.client, err = mongo.Connect(context.Background(), options.Client().ApplyURI(mgUrl)); err != nil {
		return nil, err
	} else if err := r.client.Ping(context.Background(), nil); err != nil {
		return nil, err
	} else {
		db := r.client.Database("WBA_online_ordering")
		fmt.Println(db)
	}
	return r, nil
}
