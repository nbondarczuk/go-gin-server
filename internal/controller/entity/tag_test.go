package entity

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"go-gin-server/internal/controller"
)

type TagControllerSuite struct {
	suite.Suite
	testDatabase *controller.TestDatabase
	backend      *controller.MongoBackend
}

func (s *TagControllerSuite) SetupSuite() {
	s.testDatabase = controller.SetupTestDatabase()
	controller.InitWithMongo("testdb", s.testDatabase.DbAddress)
	var err error
	s.backend, err = controller.WithMongo()
	if err != nil {
		panic(err)
	}
}

func (s *TagControllerSuite) TearDownSuite() {
	s.testDatabase.TearDown()
}

func (s *TagControllerSuite) TestTagCreate() {
	s.Run("new simpe object creation", func() {
		tc, err := NewTagController()
		if err != nil {
			panic(err)
		}
		tag := Tag{
			label: "label",
			color: "color",
		}
		rv, err := tc.Create(&tag)
		s.Nil(err)
		s.Equal(rv.label, tag.label)
		s.Equal(rv.color, tag.color)
	})
}

func TestAll(t *testing.T) {
	suite.Run(t, new(TagControllerSuite))
}
