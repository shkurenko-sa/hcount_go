package server

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func NewClient(ctx context.Context, s *Server) (*mongo.Database, error) {
	mongo_url := "mongodb://" + s.config.MongoUser + ":" + s.config.MongoPassword + "@mongo:27017/"
	client_options := options.Client().ApplyURI(mongo_url)

	client, err := mongo.NewClient(client_options)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	return client.Database(s.config.MongoDB), nil

}
