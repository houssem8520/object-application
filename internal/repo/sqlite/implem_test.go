package sqlite_test

import (
	"context"
	"io/ioutil"
	"testing"

	"github.com/object.application/internal/common"
	"github.com/object.application/internal/repo/sqlite"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type SqliteRepoTestSuite struct {
	suite.Suite
	db         *gorm.DB
	repository *sqlite.SqliteRepository
	assert     *assert.Assertions
}

func (s *SqliteRepoTestSuite) SetupSuite() {

	repo, err := sqlite.New("file::memory:?cache=shared")

	if err != nil {
		panic(err)
	}

	s.repository = repo

	t := s.T()
	s.assert = assert.New(t)

	s.db = repo.Db

	data, err := ioutil.ReadFile("../../../resources/init.sql")
	if err != nil {
		s.T().Fatalf("Prepare Test failed: %v", err)
	}

	tx := s.db.Exec(string(data))

	if tx.Error != nil {
		panic(err)
	}
}

func TestSqliteRepoTestSuite(t *testing.T) {
	suite.Run(t, new(SqliteRepoTestSuite))
}

func (s *SqliteRepoTestSuite) TestPutObject() {

	objectId, err := s.repository.PutObject(context.Background(), "bucket", "first_object")

	s.assert.Nil(err)
	s.assert.NotEmpty(objectId)
}

func (s *SqliteRepoTestSuite) TestDeleteObject() {

	s.db.Exec("INSERT INTO OBJECTS (ObjectID,Bucket,Payload) VALUES (1,'bucket','PAYLOAD')")
	err := s.repository.DeleteObject(context.Background(), "bucket", 1)

	s.assert.Nil(err)
}

func (s *SqliteRepoTestSuite) TestDeleteObject_NotFound() {

	err := s.repository.DeleteObject(context.Background(), "bucket", 1)

	s.assert.NotNil(err)
	s.assert.ErrorAs(err, &common.ErrNoObjectFound{})
}

func (s *SqliteRepoTestSuite) TestGetObject() {

	s.db.Exec("INSERT INTO OBJECTS (ObjectID,Bucket,Payload) VALUES (1,'bucket','PAYLOAD')")
	object, err := s.repository.GetObject(context.Background(), "bucket", 1)

	s.Assert().Nil(err)

	s.Assert().Equal("PAYLOAD", object.Payload)
}

func (s *SqliteRepoTestSuite) TearDownTest() {
	s.db.Exec("DELETE FROM OBJECTS")
}
