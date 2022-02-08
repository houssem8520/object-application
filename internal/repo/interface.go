package repo

import (
	"context"

	"github.com/object.application/internal/entity"
)

type ObjectRepository interface {
	PutObject(ctx context.Context, bucket string, objectContent string) (int64, error)
	GetObject(ctx context.Context, bucket string, objectID int64) (*entity.Object, error)
	DeleteObject(ctx context.Context, bucket string, objectID int64) error
	GetObjectByBucketAndContent(ctx context.Context, bucket string, content string) (*int64, error)
}
