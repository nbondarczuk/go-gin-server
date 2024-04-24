package entity

import (
	"go-gin-server/internal/controller"
	"os"
	"testing"
)

var testDatabase *controller.TestDatabase

func Setup() {
	testDatabase = controller.SetupTestDatabase()
	controller.InitWithMongo("testdb", testDatabase.DbAddress)
	var err error
	_, err = controller.WithMongo()
	if err != nil {
		panic(err)
	}
}

func TearDown() {
	testDatabase.TearDown()
}

func TestMain(m *testing.M) {
	Setup()

	rc := m.Run()

	TearDown()

	os.Exit(rc)
}

func BenchmarkTagControllerCreate(b *testing.B) {
	Setup()

	tc, err := NewTagController()
	if err != nil {
		panic(err)
	}
	tag := Tag{
		label: "label",
		color: "color",
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, _ = tc.Create(&tag)
	}

	b.StopTimer()

	TearDown()
}
