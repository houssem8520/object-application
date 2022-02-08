package service

import (
	"context"

	"github.com/object.application/internal/entity"
	"github.com/object.application/internal/repo"
)

type ObjectsService struct {
	objectRepository repo.ObjectRepository
}

func New(objectRepository repo.ObjectRepository) *ObjectsService {
	return &ObjectsService{
		objectRepository: objectRepository,
	}
}

func (s *ObjectsService) GetObject(ctx context.Context, bucket string, objectID int64) (*entity.Object, error) {

	object, err := s.objectRepository.GetObject(ctx, bucket, objectID)

	if err != nil {
		return nil, err
	}

	return object, nil
}

func (s *ObjectsService) DeleteObject(ctx context.Context, bucket string, objectID int64) error {

	err := s.objectRepository.DeleteObject(ctx, bucket, objectID)

	return err
}

func (s *ObjectsService) PutObject(ctx context.Context, bucket string, payload string) (int64, error) {

	objectId, err := s.objectRepository.GetObjectByBucketAndContent(ctx, bucket, payload)
	if err != nil {
		return 0, err
	}

	if objectId != nil {
		return *objectId, nil
	}

	return s.objectRepository.PutObject(ctx, bucket, payload)
}
