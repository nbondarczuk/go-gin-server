package entity

import (
	"os"
	"testing"

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
