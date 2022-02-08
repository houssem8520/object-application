package common

import "github.com/pkg/errors"

// Error custom is returned.
type (
	ErrInternalError struct{ error }
	ErrNoObjectFound struct{ error }
)

func NewErrInternalError(err error, msg string, args ...interface{}) ErrInternalError {
	return ErrInternalError{errors.Wrapf(err, msg, args...)}
}

func NewErrNoObjectFound(id int64, bucket string) ErrNoObjectFound {
	return ErrNoObjectFound{errors.Errorf("failed to find object with id: %v in the bucket: %v ", id, bucket)}
}
