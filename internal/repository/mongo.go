package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	DBName string
	URL    string
)

// Init opens connection to Mongo DB server. It impelments logic specific
// for this kind of data store.  The parameter like db name and url are
// saved in the package state. They are in fact immutable after init.
// Thet will be used to produce concrete connection to a database with an URL.
func InitWithMongo(name, url string) error {
	DBName = name
	URL = url
	return nil
}

// MongoRepository keeps the DB connecton state to perform DB operations.
type MongoRepository struct {
	Client *mongo.Client
	DBName string
	URL    string
}

// WithMongo produces connection handle to be used in the DB operations.
func WithMongo() (*MongoRepository, error) {
	opts := options.Client().ApplyURI(URL)
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		return nil, err
	}
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		return nil, err
	}
	r := &MongoRepository{
		Client: client,
		DBName: DBName,
	}
	return r, nil
}
