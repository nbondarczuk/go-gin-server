package entity

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/mongo"

	"go-gin-server/internal/logging"
	"go-gin-server/internal/repository"
)

var (
	testDatabase   *repository.TestDatabase
	testRepository *TagRepository
)

func setup() {
	err := logging.Init("testing", "DEBUG", "text")
	if err != nil {
		panic(err)
	}
	testDatabase = repository.SetupTestDatabase()
	repository.InitWithMongo("testdb", testDatabase.DbAddress)
	_, err = repository.WithMongo()
	if err != nil {
		panic(err)
	}
	testRepository, err = NewTagRepository()
	if err != nil {
		panic(err)
	}
	cleanAll()
}

func tearDown() {
	cleanAll()
	testDatabase.TearDown()
}

func cleanAll() {
	err := testRepository.Drop()
	if err != nil {
		panic(err)
	}
}

// TestMain is the main entry point for testing and benchmarking.
func TestMain(m *testing.M) {
	setup()
	rc := m.Run()
	tearDown()
	os.Exit(rc)
}

func TestTagCreate(t *testing.T) {
	cleanAll()
	tag := Tag{
		Label: "label1",
		Color: "color1",
	}
	rv, err := testRepository.Create(&tag)
	if assert.Nil(t, err) {
		assert.False(t, rv.ID.IsZero())
		assert.Equal(t, tag.Label, rv.Label)
		assert.Equal(t, tag.Color, rv.Color)
	}
}

func BenchmarkTagCreate(b *testing.B) {
	cleanAll()
	tag := Tag{
		Label: "label",
		Color: "color",
	}
	for i := 0; i < b.N; i++ {
		_, _ = testRepository.Create(&tag)
	}
}

func TestTagRead(t *testing.T) {
	cleanAll()
	tag := Tag{
		Label: "label2",
		Color: "color2",
	}
	rv1, err1 := testRepository.Create(&tag)
	require.Nil(t, err1)
	rv2, err2 := testRepository.Read()
	if assert.Nil(t, err2) {
		assert.Equal(t, 1, len(rv2))
		assert.Equal(t, rv1.ID.Hex(), rv2[0].ID.Hex())
		assert.Equal(t, rv1.Label, rv2[0].Label)
		assert.Equal(t, rv1.Color, rv2[0].Color)
	}
}

func TestTagReadOne(t *testing.T) {
	cleanAll()
	tag := Tag{
		Label: "label3",
		Color: "color3",
	}
	rv1, err1 := testRepository.Create(&tag)
	require.Nil(t, err1)
	require.False(t, rv1.ID.IsZero())
	rv2, err2 := testRepository.ReadOne(rv1.ID.Hex())
	if assert.Nil(t, err2) {
		assert.False(t, rv2.ID.IsZero())
		assert.Equal(t, rv1.ID.Hex(), rv2.ID.Hex())
		assert.Equal(t, rv1.Label, rv2.Label)
		assert.Equal(t, rv1.Color, rv2.Color)
	}
}

func TestTagUpdateOne(t *testing.T) {
	cleanAll()
	tag := Tag{
		Label: "label4",
		Color: "color4",
	}
	rv1, err1 := testRepository.Create(&tag)
	require.Nil(t, err1)
	require.False(t, rv1.ID.IsZero())
	tag2 := Tag{
		Label: "label4.2",
		Color: "color4.2",
	}
	err2 := testRepository.UpdateOne(rv1.ID.Hex(), &tag2)
	require.Nil(t, err2)
	rv3, err3 := testRepository.ReadOne(rv1.ID.Hex())
	if assert.Nil(t, err3) {
		assert.False(t, rv3.ID.IsZero())
		assert.Equal(t, rv1.ID.Hex(), rv3.ID.Hex())
		assert.Equal(t, tag2.Label, rv3.Label)
		assert.Equal(t, tag2.Color, rv3.Color)
	}
}

func TestTagDeleteOne(t *testing.T) {
	cleanAll()
	tag := Tag{
		Label: "label5",
		Color: "color5",
	}
	rv1, err1 := testRepository.Create(&tag)
	require.Nil(t, err1)
	require.False(t, rv1.ID.IsZero())
	err2 := testRepository.DeleteOne(rv1.ID.Hex())
	require.Nil(t, err2)
	_, err3 := testRepository.ReadOne(rv1.ID.Hex())
	if assert.NotNil(t, err3) {
		assert.Error(t, mongo.ErrNoDocuments, err3)
	}
}

func TestDrop(t *testing.T) {
	cleanAll()
	tag := Tag{
		Label: "label5",
		Color: "color5",
	}
	rv1, err1 := testRepository.Create(&tag)
	require.Nil(t, err1)
	require.False(t, rv1.ID.IsZero())
	err2 := testRepository.Drop()
	require.Nil(t, err2)
	_, err3 := testRepository.ReadOne(rv1.ID.Hex())
	if assert.NotNil(t, err3) {
		assert.Error(t, mongo.ErrNoDocuments, err3)
	}
}
