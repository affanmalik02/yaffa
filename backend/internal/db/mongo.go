package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect(ctx context.Context, uri string) (*mongo.Client, error) {
	clientOpts := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func GetAPIKey(ctx context.Context, client *mongo.Client, key string) (map[string]interface{}, error) {
	// TODO: query keys collection and return document
	return nil, nil
}

func StoreMetadata(ctx context.Context, client *mongo.Client, doc map[string]interface{}) error {
	// TODO: insert or upsert document into metadata collection
	return nil
}

func QueryMetadata(ctx context.Context, client *mongo.Client, filters map[string]interface{}) ([]map[string]interface{}, error) {
	// TODO: query metadata collection
	return nil, nil
}