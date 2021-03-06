// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/runtime/security"
	"github.com/go-openapi/spec"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/object.application/models"
	"github.com/object.application/restapi/operations/object_management"
	"github.com/object.application/restapi/operations/secure_object_management"
	securityops "github.com/object.application/restapi/operations/security"
)

// NewObjectApplicationAPI creates a new ObjectApplication instance
func NewObjectApplicationAPI(spec *loads.Document) *ObjectApplicationAPI {
	return &ObjectApplicationAPI{
		handlers:            make(map[string]map[string]http.Handler),
		formats:             strfmt.Default,
		defaultConsumes:     "application/json",
		defaultProduces:     "application/json",
		customConsumers:     make(map[string]runtime.Consumer),
		customProducers:     make(map[string]runtime.Producer),
		PreServerShutdown:   func() {},
		ServerShutdown:      func() {},
		spec:                spec,
		useSwaggerUI:        false,
		ServeError:          errors.ServeError,
		BasicAuthenticator:  security.BasicAuth,
		APIKeyAuthenticator: security.APIKeyAuth,
		BearerAuthenticator: security.BearerAuth,

		TxtConsumer: runtime.TextConsumer(),

		JSONProducer: runtime.JSONProducer(),

		SecurityGetAuthCallbackHandler: securityops.GetAuthCallbackHandlerFunc(func(params securityops.GetAuthCallbackParams) middleware.Responder {
			return middleware.NotImplemented("operation security.GetAuthCallback has not yet been implemented")
		}),
		SecurityGetLoginHandler: securityops.GetLoginHandlerFunc(func(params securityops.GetLoginParams) middleware.Responder {
			return middleware.NotImplemented("operation security.GetLogin has not yet been implemented")
		}),
		ObjectManagementDeleteObjectHandler: object_management.DeleteObjectHandlerFunc(func(params object_management.DeleteObjectParams) middleware.Responder {
			return middleware.NotImplemented("operation object_management.DeleteObject has not yet been implemented")
		}),
		SecureObjectManagementDeleteSecureObjectHandler: secure_object_management.DeleteSecureObjectHandlerFunc(func(params secure_object_management.DeleteSecureObjectParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation secure_object_management.DeleteSecureObject has not yet been implemented")
		}),
		ObjectManagementGetObjectHandler: object_management.GetObjectHandlerFunc(func(params object_management.GetObjectParams) middleware.Responder {
			return middleware.NotImplemented("operation object_management.GetObject has not yet been implemented")
		}),
		SecureObjectManagementGetSecureObjectHandler: secure_object_management.GetSecureObjectHandlerFunc(func(params secure_object_management.GetSecureObjectParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation secure_object_management.GetSecureObject has not yet been implemented")
		}),
		ObjectManagementPutObjectHandler: object_management.PutObjectHandlerFunc(func(params object_management.PutObjectParams) middleware.Responder {
			return middleware.NotImplemented("operation object_management.PutObject has not yet been implemented")
		}),
		SecureObjectManagementPutSecureObjectHandler: secure_object_management.PutSecureObjectHandlerFunc(func(params secure_object_management.PutSecureObjectParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation secure_object_management.PutSecureObject has not yet been implemented")
		}),

		OauthSecurityAuth: func(token string, scopes []string) (*models.Principal, error) {
			return nil, errors.NotImplemented("oauth2 bearer auth (OauthSecurity) has not yet been implemented")
		},
		// default authorizer is authorized meaning no requests are blocked
		APIAuthorizer: security.Authorized(),
	}
}

/*ObjectApplicationAPI This API is dedicated to store objects to a bucket. */
type ObjectApplicationAPI struct {
	spec            *loads.Document
	context         *middleware.Context
	handlers        map[string]map[string]http.Handler
	formats         strfmt.Registry
	customConsumers map[string]runtime.Consumer
	customProducers map[string]runtime.Producer
	defaultConsumes string
	defaultProduces string
	Middleware      func(middleware.Builder) http.Handler
	useSwaggerUI    bool

	// BasicAuthenticator generates a runtime.Authenticator from the supplied basic auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BasicAuthenticator func(security.UserPassAuthentication) runtime.Authenticator

	// APIKeyAuthenticator generates a runtime.Authenticator from the supplied token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	APIKeyAuthenticator func(string, string, security.TokenAuthentication) runtime.Authenticator

	// BearerAuthenticator generates a runtime.Authenticator from the supplied bearer token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BearerAuthenticator func(string, security.ScopedTokenAuthentication) runtime.Authenticator

	// TxtConsumer registers a consumer for the following mime types:
	//   - text/plain
	TxtConsumer runtime.Consumer

	// JSONProducer registers a producer for the following mime types:
	//   - application/json
	JSONProducer runtime.Producer

	// OauthSecurityAuth registers a function that takes an access token and a collection of required scopes and returns a principal
	// it performs authentication based on an oauth2 bearer token provided in the request
	OauthSecurityAuth func(string, []string) (*models.Principal, error)

	// APIAuthorizer provides access control (ACL/RBAC/ABAC) by providing access to the request and authenticated principal
	APIAuthorizer runtime.Authorizer

	// SecurityGetAuthCallbackHandler sets the operation handler for the get auth callback operation
	SecurityGetAuthCallbackHandler securityops.GetAuthCallbackHandler
	// SecurityGetLoginHandler sets the operation handler for the get login operation
	SecurityGetLoginHandler securityops.GetLoginHandler
	// ObjectManagementDeleteObjectHandler sets the operation handler for the delete object operation
	ObjectManagementDeleteObjectHandler object_management.DeleteObjectHandler
	// SecureObjectManagementDeleteSecureObjectHandler sets the operation handler for the delete secure object operation
	SecureObjectManagementDeleteSecureObjectHandler secure_object_management.DeleteSecureObjectHandler
	// ObjectManagementGetObjectHandler sets the operation handler for the get object operation
	ObjectManagementGetObjectHandler object_management.GetObjectHandler
	// SecureObjectManagementGetSecureObjectHandler sets the operation handler for the get secure object operation
	SecureObjectManagementGetSecureObjectHandler secure_object_management.GetSecureObjectHandler
	// ObjectManagementPutObjectHandler sets the operation handler for the put object operation
	ObjectManagementPutObjectHandler object_management.PutObjectHandler
	// SecureObjectManagementPutSecureObjectHandler sets the operation handler for the put secure object operation
	SecureObjectManagementPutSecureObjectHandler secure_object_management.PutSecureObjectHandler

	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)

	// PreServerShutdown is called before the HTTP(S) server is shutdown
	// This allows for custom functions to get executed before the HTTP(S) server stops accepting traffic
	PreServerShutdown func()

	// ServerShutdown is called when the HTTP(S) server is shut down and done
	// handling all active connections and does not accept connections any more
	ServerShutdown func()

	// Custom command line argument groups with their descriptions
	CommandLineOptionsGroups []swag.CommandLineOptionsGroup

	// User defined logger function.
	Logger func(string, ...interface{})
}

// UseRedoc for documentation at /docs
func (o *ObjectApplicationAPI) UseRedoc() {
	o.useSwaggerUI = false
}

// UseSwaggerUI for documentation at /docs
func (o *ObjectApplicationAPI) UseSwaggerUI() {
	o.useSwaggerUI = true
}

// SetDefaultProduces sets the default produces media type
func (o *ObjectApplicationAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *ObjectApplicationAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *ObjectApplicationAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *ObjectApplicationAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *ObjectApplicationAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *ObjectApplicationAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *ObjectApplicationAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the ObjectApplicationAPI
func (o *ObjectApplicationAPI) Validate() error {
	var unregistered []string

	if o.TxtConsumer == nil {
		unregistered = append(unregistered, "TxtConsumer")
	}

	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.OauthSecurityAuth == nil {
		unregistered = append(unregistered, "OauthSecurityAuth")
	}

	if o.SecurityGetAuthCallbackHandler == nil {
		unregistered = append(unregistered, "security.GetAuthCallbackHandler")
	}
	if o.SecurityGetLoginHandler == nil {
		unregistered = append(unregistered, "security.GetLoginHandler")
	}
	if o.ObjectManagementDeleteObjectHandler == nil {
		unregistered = append(unregistered, "object_management.DeleteObjectHandler")
	}
	if o.SecureObjectManagementDeleteSecureObjectHandler == nil {
		unregistered = append(unregistered, "secure_object_management.DeleteSecureObjectHandler")
	}
	if o.ObjectManagementGetObjectHandler == nil {
		unregistered = append(unregistered, "object_management.GetObjectHandler")
	}
	if o.SecureObjectManagementGetSecureObjectHandler == nil {
		unregistered = append(unregistered, "secure_object_management.GetSecureObjectHandler")
	}
	if o.ObjectManagementPutObjectHandler == nil {
		unregistered = append(unregistered, "object_management.PutObjectHandler")
	}
	if o.SecureObjectManagementPutSecureObjectHandler == nil {
		unregistered = append(unregistered, "secure_object_management.PutSecureObjectHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *ObjectApplicationAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *ObjectApplicationAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {
	result := make(map[string]runtime.Authenticator)
	for name := range schemes {
		switch name {
		case "OauthSecurity":
			result[name] = o.BearerAuthenticator(name, func(token string, scopes []string) (interface{}, error) {
				return o.OauthSecurityAuth(token, scopes)
			})

		}
	}
	return result
}

// Authorizer returns the registered authorizer
func (o *ObjectApplicationAPI) Authorizer() runtime.Authorizer {
	return o.APIAuthorizer
}

// ConsumersFor gets the consumers for the specified media types.
// MIME type parameters are ignored here.
func (o *ObjectApplicationAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {
	result := make(map[string]runtime.Consumer, len(mediaTypes))
	for _, mt := range mediaTypes {
		switch mt {
		case "text/plain":
			result["text/plain"] = o.TxtConsumer
		}

		if c, ok := o.customConsumers[mt]; ok {
			result[mt] = c
		}
	}
	return result
}

// ProducersFor gets the producers for the specified media types.
// MIME type parameters are ignored here.
func (o *ObjectApplicationAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {
	result := make(map[string]runtime.Producer, len(mediaTypes))
	for _, mt := range mediaTypes {
		switch mt {
		case "application/json":
			result["application/json"] = o.JSONProducer
		}

		if p, ok := o.customProducers[mt]; ok {
			result[mt] = p
		}
	}
	return result
}

// HandlerFor gets a http.Handler for the provided operation method and path
func (o *ObjectApplicationAPI) HandlerFor(method, path string) (http.Handler, bool) {
	if o.handlers == nil {
		return nil, false
	}
	um := strings.ToUpper(method)
	if _, ok := o.handlers[um]; !ok {
		return nil, false
	}
	if path == "/" {
		path = ""
	}
	h, ok := o.handlers[um][path]
	return h, ok
}

// Context returns the middleware context for the object application API
func (o *ObjectApplicationAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *ObjectApplicationAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened
	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/auth/callback"] = securityops.NewGetAuthCallback(o.context, o.SecurityGetAuthCallbackHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/login"] = securityops.NewGetLogin(o.context, o.SecurityGetLoginHandler)
	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/objects/{bucket}/{objectID}"] = object_management.NewDeleteObject(o.context, o.ObjectManagementDeleteObjectHandler)
	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/secure/objects/{bucket}/{objectID}"] = secure_object_management.NewDeleteSecureObject(o.context, o.SecureObjectManagementDeleteSecureObjectHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/objects/{bucket}/{objectID}"] = object_management.NewGetObject(o.context, o.ObjectManagementGetObjectHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/secure/objects/{bucket}/{objectID}"] = secure_object_management.NewGetSecureObject(o.context, o.SecureObjectManagementGetSecureObjectHandler)
	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/objects/{bucket}"] = object_management.NewPutObject(o.context, o.ObjectManagementPutObjectHandler)
	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/secure/objects/{bucket}"] = secure_object_management.NewPutSecureObject(o.context, o.SecureObjectManagementPutSecureObjectHandler)
}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *ObjectApplicationAPI) Serve(builder middleware.Builder) http.Handler {
	o.Init()

	if o.Middleware != nil {
		return o.Middleware(builder)
	}
	if o.useSwaggerUI {
		return o.context.APIHandlerSwaggerUI(builder)
	}
	return o.context.APIHandler(builder)
}

// Init allows you to just initialize the handler cache, you can then recompose the middleware as you see fit
func (o *ObjectApplicationAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}

// RegisterConsumer allows you to add (or override) a consumer for a media type.
func (o *ObjectApplicationAPI) RegisterConsumer(mediaType string, consumer runtime.Consumer) {
	o.customConsumers[mediaType] = consumer
}

// RegisterProducer allows you to add (or override) a producer for a media type.
func (o *ObjectApplicationAPI) RegisterProducer(mediaType string, producer runtime.Producer) {
	o.customProducers[mediaType] = producer
}

// AddMiddlewareFor adds a http middleware to existing handler
func (o *ObjectApplicationAPI) AddMiddlewareFor(method, path string, builder middleware.Builder) {
	um := strings.ToUpper(method)
	if path == "/" {
		path = ""
	}
	o.Init()
	if h, ok := o.handlers[um][path]; ok {
		o.handlers[method][path] = builder(h)
	}
}
