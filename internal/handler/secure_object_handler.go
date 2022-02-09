package handler

import (
	"context"
	"time"

	"github.com/go-openapi/runtime/middleware"
	"github.com/object.application/internal/common"
	"github.com/object.application/internal/service"
	"github.com/object.application/models"
	"github.com/object.application/restapi/operations/object_management"
	"github.com/object.application/restapi/operations/secure_object_management"
)

func DeleteSecureObject(s service.ObjectService) func(params secure_object_management.DeleteSecureObjectParams,principal *models.Principal) middleware.Responder {
	return func(params secure_object_management.DeleteSecureObjectParams,principal *models.Principal) middleware.Responder {
		// Propagate HTTP req context with a timeout.
		ctx, cancel := context.WithTimeout(params.HTTPRequest.Context(), time.Minute)
		defer cancel()

		if err := s.DeleteObject(ctx, params.Bucket, params.ObjectID); err != nil {
			switch err.(type) {
			case common.ErrNoObjectFound:
				return secure_object_management.NewDeleteSecureObjectNotFound()
			default:
				return secure_object_management.NewDeleteSecureObjectInternalServerError()
			}
		}
		return object_management.NewDeleteObjectOK()
	}
}

func PutSecureObject(s service.ObjectService) func(params secure_object_management.PutSecureObjectParams,principal *models.Principal) middleware.Responder {
	return func(params secure_object_management.PutSecureObjectParams,principal *models.Principal) middleware.Responder {
		// Propagate HTTP req context with a timeout.
		ctx, cancel := context.WithTimeout(params.HTTPRequest.Context(), time.Minute)
		defer cancel()

		objectID, err := s.PutObject(ctx, params.Bucket, params.HTTPText)

		payload := object_management.PutObjectCreatedBody{
			ID: objectID,
		}
		if err != nil {

			return secure_object_management.NewPutSecureObjectInternalServerError()
		}

		return object_management.NewPutObjectCreated().WithPayload(&payload)
	}
}

func GetSecureObject(s service.ObjectService) func(params secure_object_management.GetSecureObjectParams,principal *models.Principal) middleware.Responder {
	return func(params secure_object_management.GetSecureObjectParams,principal *models.Principal) middleware.Responder {
		// Propagate HTTP req context with a timeout.
		ctx, cancel := context.WithTimeout(params.HTTPRequest.Context(), time.Minute)
		defer cancel()

		object, err := s.GetObject(ctx, params.Bucket, params.ObjectID)

		if err != nil {
			switch err.(type) {
			case common.ErrNoObjectFound:
				return secure_object_management.NewGetSecureObjectNotFound()
			default:
				return secure_object_management.NewGetSecureObjectInternalServerError()
			}
		}

		payload := secure_object_management.GetSecureObjectOKBody{
			Object: &models.Object{
				Data: &object.Payload,
				ID:   &object.ObjectID,
			},
		}
		return secure_object_management.NewGetSecureObjectOK().WithPayload(&payload)
	}
}

