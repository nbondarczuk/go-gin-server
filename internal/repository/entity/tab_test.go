package entity

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"go-gin-server/internal/repository"
)

var (
	testDatabase   *repository.TestDatabase
	testRepository *TagRepository
)

func setup() {
	testDatabase = repository.SetupTestDatabase()
	repository.InitWithMongo("testdb", testDatabase.DbAddress)
	var err error
	_, err = repository.WithMongo()
	if err != nil {
		panic(err)
	}
	testRepository, err = NewTagRepository()
	if err != nil {
		panic(err)
	}
}

func tearDown() {
	testDatabase.TearDown()
}

// TestMain is the main entry point for testing and benchmarking.
func TestMain(m *testing.M) {
	setup()
	rc := m.Run()
	tearDown()
	os.Exit(rc)
}

func TestTagCreate(t *testing.T) {
	tag := Tag{
		label: "label",
		color: "color",
	}
	rv, err := testRepository.Create(&tag)
	assert.Nil(t, err)
	assert.Equal(t, rv.label, tag.label)
	assert.Equal(t, rv.color, tag.color)
}

func BenchmarkTagCreate(b *testing.B) {
	tag := Tag{
		label: "label",
		color: "color",
	}
	for i := 0; i < b.N; i++ {
		_, _ = testRepository.Create(&tag)
	}
}