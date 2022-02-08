package service_test

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/object.application/internal/entity"
	mock_repo "github.com/object.application/internal/repo/mock"
	"github.com/object.application/internal/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ObjectServiceTestSuite struct {
	suite.Suite
	mockObjectRepository *mock_repo.MockObjectRepository
	objectService        *service.ObjectsService
	ctrl                 *gomock.Controller
	assert               *assert.Assertions
}

func (s *ObjectServiceTestSuite) SetupSuite() {
	s.assert = assert.New(s.T())
	s.ctrl = gomock.NewController(s.T())
}

func (s *ObjectServiceTestSuite) SetupTest() {
	s.mockObjectRepository = mock_repo.NewMockObjectRepository(s.ctrl)
	s.objectService = service.New(s.mockObjectRepository)
}

func TestObjectServiceTestSuite(t *testing.T) {
	suite.Run(t, new(ObjectServiceTestSuite))
}

func (s *ObjectServiceTestSuite) TestGetObject() {

	expectedObject := entity.Object{
		ObjectID: 1,
		Bucket:   "MyFirstBucket",
		Payload: `GET / HTTP/1.1
User-Agent: PostmanRuntime/7.29.0
Accept: */*
Accept-Encoding: gzip, deflate, br
Connection: keep-alive
Referer: http://google.com/
Host: www.google.com`,
	}

	s.mockObjectRepository.EXPECT().GetObject(gomock.Any(), expectedObject.Bucket, expectedObject.ObjectID).Return(&expectedObject, nil)

	object, err := s.objectService.GetObject(context.Background(), "MyFirstBucket", 1)

	s.assert.Nil(err)
	s.assert.Equal(expectedObject, *object)

}
