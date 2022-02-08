package service

import (
	"context"

	"github.com/object.application/internal/entity"
)

type ObjectService interface {
	PutObject(ctx context.Context, bucket string, objectContent string) (int64, error)
	GetObject(ctx context.Context, bucket string, objectID int64) (*entity.Object, error)
	DeleteObject(ctx context.Context, bucket string, objectID int64) error
}
