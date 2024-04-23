package controller

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	BackendDBName string
	BackendURL    string
)

// Init opens connection to Mongo DB server. It impelments logic specific
// for this kind of data store.  The parameter like db name and url are
// saved in the package state. They are in fact immutable after init.
// Thet will be used to produce concrete connection to a database with an URL.
func InitWithMongo(name, url string) error {
	BackendDBName = name
	BackendURL = url
	return nil
}

// MongoBackend keeps the DB connecton state to perform DB operations.
type MongoBackend struct {
	Client *mongo.Client
	DBName string
	URL    string
}

// WithMongo produces connection handle to be used in the trasactions.
func WithMongo() (*MongoBackend, error) {
	opts := options.Client().ApplyURI(BackendURL)
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		return nil, ErrBackendClientConnect
	}
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		return nil, ErrBackendClientPing
	}
	backend := &MongoBackend{
		Client: client,
		DBName: BackendDBName,
	}
	return backend, nil
}
