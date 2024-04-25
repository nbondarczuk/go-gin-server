package entity

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"go-gin-server/internal/logging"
	"go-gin-server/internal/repository"
)

const TagCollectionName = "tag"

// Tag is the entity mnaged by the repository.
type Tag struct {
	ID      primitive.ObjectID `json:"id" bson:"_id"`
	Label   string             `json:"label" bson:"label"`
	Color   string             `json:"color" bson:"color"`
	Created time.Time          `json:"created" bson:"created"`
	Updated time.Time          `json:"updated" bson:"updated"`
}

// TagRepository is a container for resource accerss action state.
type TagRepository struct {
	repository *repository.MongoRepository
	ctx        context.Context
	collection *mongo.Collection
}

// NewTagRepository handles resource access action in its own context.
func NewTagRepository() (*TagRepository, error) {
	repository, err := repository.WithMongo()
	if err != nil {
		return nil, err
	}
	collection := repository.Client.Database(repository.DBName).Collection(TagCollectionName)
	tc := TagRepository{
		repository: repository,
		ctx:        context.Background(),
		collection: collection,
	}
	return &tc, nil
}

// Create an object with new oid allocated.
func (tc *TagRepository) Create(tag *Tag) (*Tag, error) {
	if tag.ID.IsZero() {
		tag.ID = primitive.NewObjectID()
	}
	tag.Created = time.Now()
	result, err := tc.collection.InsertOne(tc.ctx, tag)
	if err != nil {
		return nil, err
	}
	logging.Logger.Debug("Created tag entity", slog.String("ID", fmt.Sprintf("%v", result.InsertedID)))
	return tag, nil
}

// Read fetches a whole set of objects.
func (tc *TagRepository) Read() ([]Tag, error) {
	cursor, err := tc.collection.Find(tc.ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(tc.ctx)

	tags := make([]Tag, 0)
	for cursor.Next(tc.ctx) {
		var tag Tag
		cursor.Decode(&tag)
		tags = append(tags, tag)
	}
	logging.Logger.Debug("Read tag entities", slog.Int("count", len(tags)))
	return tags, nil
}

// ReadOne fetches one object by primary key.
func (tc *TagRepository) ReadOne(id string) (Tag, error) {
	var tag Tag
	ID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return Tag{}, err
	}
	err = tc.collection.FindOne(tc.ctx, bson.M{"_id": ID}).Decode(&tag)
	if err != nil {
		return Tag{}, err
	}
	logging.Logger.Debug("Read tag entity", slog.String("ID", fmt.Sprintf("%v", tag.ID)))
	return tag, nil
}

// UpdateOne replaces all attributes of an existing object.
func (tc *TagRepository) UpdateOne(id string, tag *Tag) error {
	updated := time.Now()
	ID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = tc.collection.UpdateOne(tc.ctx,
		bson.M{"_id": ID},
		bson.D{{"$set",
			bson.D{
				{"Label", tag.Label},
				{"Color", tag.Color},
				{"Updated", updated},
			}}})
	if err != nil {
		return err
	}
	logging.Logger.Debug("Updated tag entity", slog.String("ID", fmt.Sprintf("%v", ID)))
	return nil
}

// DeleteOne removes an object from collection using primary key.
func (tc *TagRepository) DeleteOne(id string) error {
	ID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = tc.collection.DeleteOne(tc.ctx, bson.M{"_id": ID})
	if err != nil {
		return err
	}
	logging.Logger.Debug("Updated tag entity", slog.String("ID", fmt.Sprintf("%v", ID)))
	return nil
}

// Delete removes all objects from collection.
func (tc *TagRepository) Drop() error {
	err := tc.collection.Drop(tc.ctx)
	if err != nil {
		return err
	}
	return nil
}
