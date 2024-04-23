package entity

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"go-gin-server/internal/controller"
)

const TagCollectionName = "tag"

// Tag is the entity mnaged by the controller.
type Tag struct {
	oid       primitive.ObjectID `json:"id" bson:"_id"`
	label     string             `json:"tabel" bson:"label"`
	color     string             `json:"color" bson:"color"`
	createdAt time.Time          `json:"created_at" bson:"created_at"`
	updatedAt time.Time          `json:"updated_at" bson:"updated_at"`
}

// TagController is a container for resource accerss action state.
type TagController struct {
	backend    *controller.MongoBackend
	ctx        context.Context
	collection *mongo.Collection
}

// NewTagController handles resource access action in its own context.
func NewTagController() (*TagController, error) {
	backend, err := controller.WithMongo()
	if err != nil {
		return nil, err
	}
	collection := backend.Client.Database(backend.DBName).Collection(TagCollectionName)
	tc := TagController{
		backend:    backend,
		ctx:        context.Background(),
		collection: collection,
	}
	return &tc, nil
}

// Create an object with new oid allocated.
func (tc *TagController) Create(tag *Tag) (*Tag, error) {
	tag.oid = primitive.NewObjectID()
	tag.createdAt = time.Now()
	_, err := tc.collection.InsertOne(tc.ctx, tag)
	if err != nil {
		return nil, err
	}
	return tag, nil
}

// ReadAll fetches a whole set of objects.
func (tc *TagController) ReadAll() ([]Tag, error) {
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
	return tags, nil
}

// ReadOne fetches one object by primary key.
func (tc *TagController) ReadOne(oid string) (Tag, error) {
	var tag Tag
	err := tc.collection.FindOne(tc.ctx, bson.M{"_id": oid}).Decode(&tag)
	if err != nil {
		return Tag{}, err
	}
	return Tag{}, nil
}

// Update replaces all attributes of an existing object.
func (tc *TagController) Update(oid string, tag *Tag) error {
	updatedAt := time.Now()
	_, err := tc.collection.UpdateOne(tc.ctx,
		bson.M{"_id": oid},
		bson.D{{"$set",
			bson.D{
				{"label", tag.label},
				{"color", tag.color},
				{"updated_at", updatedAt},
			}}})
	if err != nil {
		return err
	}
	return nil
}

// Delete removes an object from collection using primary key.
func (tc *TagController) Delete(oid string) error {
	_, err := tc.collection.DeleteOne(tc.ctx, bson.M{"_id": oid})
	if err != nil {
		return err
	}
	return nil
}
