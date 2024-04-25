package entity

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"go-gin-server/internal/logging"
	"go-gin-server/internal/repository"
)

var (
	testDatabase   *repository.TestDatabase
	testRepository *TagRepository
)

func setup() {
	err := logging.Init("testing", "INFO", "text")
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
		label: "label",
		color: "color",
	}
	rv, err := testRepository.Create(&tag)
	if assert.Nil(t, err) {
		assert.Equal(t, rv.label, tag.label)
		assert.Equal(t, rv.color, tag.color)
	}
}

func BenchmarkTagCreate(b *testing.B) {
	cleanAll()
	tag := Tag{
		label: "label",
		color: "color",
	}
	for i := 0; i < b.N; i++ {
		_, _ = testRepository.Create(&tag)
	}
}

func TestTagReadAll(t *testing.T) {
	cleanAll()
	tag := Tag{
		label: "label1",
		color: "color1",
	}
	rv1, err1 := testRepository.Create(&tag)
	require.Nil(t, err1)
	rv2, err2 := testRepository.ReadAll()
	if assert.Nil(t, err2) {
		assert.Equal(t, 1, len(rv2))
		assert.Equal(t, rv1.oid.Hex(), rv2[0].oid.Hex())
		assert.Equal(t, rv1.label, rv2[0].label)
		assert.Equal(t, rv1.color, rv2[0].color)
	}
}

func TestTagRead(t *testing.T) {
	cleanAll()
	tag := Tag{
		label: "label2",
		color: "color2",
	}
	rv1, err1 := testRepository.Create(&tag)
	require.Nil(t, err1)
	rv2, err2 := testRepository.ReadOne(rv1.oid.Hex())
	if assert.Nil(t, err2) {
		assert.Equal(t, rv1.oid.Hex(), rv2.oid.Hex())
		assert.Equal(t, rv1.label, rv2.label)
		assert.Equal(t, rv1.color, rv2.color)
	}
}
