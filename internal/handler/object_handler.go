package handler

import (
	"context"
	"time"

	"github.com/go-openapi/runtime/middleware"
	"github.com/object.application/internal/common"
	"github.com/object.application/internal/service"
	"github.com/object.application/models"
	"github.com/object.application/restapi/operations/object_management"
)

func DeleteObject(s service.ObjectService) func(params object_management.DeleteObjectParams) middleware.Responder {
	return func(params object_management.DeleteObjectParams) middleware.Responder {
		// Propagate HTTP req context with a timeout.
		ctx, cancel := context.WithTimeout(params.HTTPRequest.Context(), time.Minute)
		defer cancel()

		if err := s.DeleteObject(ctx, params.Bucket, params.ObjectID); err != nil {
			switch err.(type) {
			case common.ErrNoObjectFound:
				return object_management.NewDeleteObjectNotFound()
			default:
				return object_management.NewDeleteObjectInternalServerError()
			}
		}
		return object_management.NewDeleteObjectOK()
	}
}

func PutObject(s service.ObjectService) func(params object_management.PutObjectParams) middleware.Responder {
	return func(params object_management.PutObjectParams ) middleware.Responder {
		// Propagate HTTP req context with a timeout.
		ctx, cancel := context.WithTimeout(params.HTTPRequest.Context(), time.Minute)
		defer cancel()

		objectID, err := s.PutObject(ctx, params.Bucket, params.HTTPText)

		payload := object_management.PutObjectCreatedBody{
			ID: objectID,
		}
		if err != nil {

			return object_management.NewPutObjectInternalServerError()
		}

		return object_management.NewPutObjectCreated().WithPayload(&payload)
	}
}

func GetObject(s service.ObjectService) func(params object_management.GetObjectParams) middleware.Responder {
	return func(params object_management.GetObjectParams) middleware.Responder {
		// Propagate HTTP req context with a timeout.
		ctx, cancel := context.WithTimeout(params.HTTPRequest.Context(), time.Minute)
		defer cancel()

		object, err := s.GetObject(ctx, params.Bucket, params.ObjectID)

		if err != nil {
			switch err.(type) {
			case common.ErrNoObjectFound:
				return object_management.NewGetObjectNotFound()
			default:
				return object_management.NewGetObjectInternalServerError()
			}
		}

		payload := object_management.GetObjectOKBody{
			Object: &models.Object{
				Data: &object.Payload,
				ID:   &object.ObjectID,
			},
		}
		return object_management.NewGetObjectOK().WithPayload(&payload)
	}
}

