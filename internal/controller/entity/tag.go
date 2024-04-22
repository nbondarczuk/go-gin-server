package entity

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"go-gin-server/internal/controller"
)

const TagCollectionName = "tag"

// Tag is the entity mnaged by the controller.
type Tag struct {
	oid   primitive.ObjectID `json:"id" bson:"_id"`
	label string             `json:"tabel" bson:"label"`
	color string             `json:"color" bson:"color"`
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
	collecction := backend.client.Database(backend.DBName).Collection(TagCollectionName)
	tc := TagController{
		backend: backend,
		ctx:     context.Background(),
	}
	return &tc, nil
}

// Create an object with new oid allocated.
func (tc *TagController) Create(tag *Tag) (*Tag, error) {
	tag.oid = primitive.NewObjectID()
	tag, err := tc.collection.InsertOne(tc.ctx, *tag)
	if err != nil {
		return nil, err
	}
	return tag, nil
}

// Read an object by primary key.
func (tc *TagController) Read(oid string) (*Tag, error) {
	return nil, nil
}

// ReadAll does the same on the whole collection.
func (tc *TagController) ReadAll() ([]Tag, error) {
	return nil, nil
}

// Update replaces all attributes of an existing object.
func (tc *TagController) Update(tag *Tag) (*Tag, error) {
	return nil, nil
}

// Patch replaces non null attributes of an existing object.
func (tc *TagController) Patch(tag *Tag) error {
	return nil
}

// Delete removes an object from collection using primary key.
func (tc *TagController) Delete(oid string) error {
	return nil
}
