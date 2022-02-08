package sqlite

import (
	"context"
	"errors"

	"github.com/object.application/internal/common"
	"github.com/object.application/internal/entity"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type SqliteRepository struct {
	Db *gorm.DB
}

func New(db_path string) (*SqliteRepository, error) {
	config := gorm.Config{}
	db, err := gorm.Open(sqlite.Open(db_path), &config)
	if err != nil {
		return nil, common.NewErrInternalError(err, "failed to create sqlLite repository")
	}

	return &SqliteRepository{db}, nil
}

func (r *SqliteRepository) GetObject(ctx context.Context, bucket string, objectID int64) (*entity.Object, error) {

	res := entity.Object{}

	if err := r.Db.WithContext(ctx).Table("Objects").Where("ObjectId = ? and Bucket = ?", objectID, bucket).First(&res).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &res, common.NewErrNoObjectFound(objectID, bucket)
		}
		return &res, common.NewErrInternalError(err, "failed to get object")
	}

	return &res, nil
}

func (r *SqliteRepository) GetObjectByBucketAndContent(ctx context.Context, bucket string, content string) (*int64, error) {

	res := entity.Object{}

	if err := r.Db.WithContext(ctx).Table("Objects").Where("Payload = ? and Bucket = ?", content, bucket).First(&res).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, common.NewErrInternalError(err, "failed to get object by bucket and content")
	}

	return &res.ObjectID, nil
}

func (r *SqliteRepository) DeleteObject(ctx context.Context, bucket string, objectID int64) error {

	res := r.Db.WithContext(ctx).Where("ObjectID = ? AND Bucket = ?", objectID, bucket).Delete(&entity.Object{})
	if err := res.Error; err != nil {

		return common.NewErrInternalError(err, "failed to delete object")
	}

	if res.RowsAffected == 0 {
		return common.NewErrNoObjectFound(objectID, bucket)
	}

	return nil
}

func (r *SqliteRepository) PutObject(ctx context.Context, bucket string, payload string) (int64, error) {

	object := entity.Object{
		Bucket:  bucket,
		Payload: payload,
	}
	res := r.Db.WithContext(ctx).Create(&object)
	if err := res.Error; err != nil {

		return 0, common.NewErrInternalError(err, "failed to create object")
	}

	return object.ObjectID, nil
}
