package db

import (
	"context"
	"fmt"

	"github.com/shkurenko-sa/hcount_go/internal/users"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type db struct {
	collection *mongo.Collection
}

func (d *db) Create(ctx context.Context, user users.User) (string, error) {
	result, err := d.collection.InsertOne(ctx, user)
	if err != nil {
		return "", err
	}
	oid, ok := result.InsertedID.(primitive.ObjectID)
	if ok {
		return oid.Hex(), nil
	}
	return "", fmt.Errorf("Failed to convert obj_id to string")
}

func (d *db) Retrive(ctx context.Context, id string) (users.User, error) {
	panic("pass")
}

func (d *db) Update(ctx context.Context, user users.User) error {
	return nil
}

func (d *db) Delete(ctx context.Context, user users.User) error {
	return nil
}

func NewStorage(database *mongo.Database, collection string) users.Storage {
	return &db{
		collection: database.Collection(collection),
	}
}
