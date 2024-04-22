package controller

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	MongoDBName string
	MongoURL    string
)

// Init opens connection to Mongo DB server. It impelments logic specific
// for this kind of data store.  The parameter like db name and url are
// saved in the package state. They are in fact immutable after init.
// Thet will be used to produce concrete connection to a database with an URL.
func InitWithMongo(name, url string) error {
	MongoDBName = name
	MongURL = url
	return nil
}

// MongoBackend keeps the DB connecton state to perform DB operations.
type MongoBackend struct {
	client *mongo.Client
	DBName string
	URL    string
}

// WithMongo produces connection handle to be used in the trasactions.
func WithMongo(name string) (*MongoBackend, error) {
	opts := options.Client().ApplyURI(MongoURL)
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		return nil, ErrClientConnect
	}
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		return nil, ErrClientPing
	}
	backend := &MongoBackend{
		client: client,
		DBName: MongoDBName,
	}
	return backend, nil
}
