// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"
	"os"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"
	log "github.com/sirupsen/logrus"

	"github.com/object.application/internal/conf"
	"github.com/object.application/internal/handler"
	"github.com/object.application/internal/repo/sqlite"
	"github.com/object.application/internal/service"
	"github.com/object.application/restapi/operations"
	"github.com/object.application/restapi/operations/object_management"
)

//go:generate swagger generate server --target ../../object.application --name ObjectApplication --spec ../target/swagger.yaml --principal interface{}
type configurationFlags struct {
	ConfFile string `short:"c" long:"conf" description:"Path to configuration file" value-name:"FILE"`
}

var confFlags configurationFlags

func configureFlags(api *operations.ObjectApplicationAPI) {
	api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{
		{
			ShortDescription: "Configuration Options",
			Options:          &confFlags,
		},
	}
}

func configureAPI(api *operations.ObjectApplicationAPI) http.Handler {
	configureGlobal()
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONProducer = runtime.JSONProducer()

	api.TxtConsumer = runtime.TextConsumer()

	repo, err := sqlite.New(conf.GetSqliteFilePath())
	if err != nil {
		log.WithError(err).Error("failed to connect to the sqlite db")
		panic(err)
	}

	objectService := service.New(repo)

	api.ObjectManagementDeleteObjectHandler = object_management.DeleteObjectHandlerFunc(handler.DeleteObject(objectService))

	api.ObjectManagementGetObjectHandler = object_management.GetObjectHandlerFunc(handler.GetObject(objectService))

	api.ObjectManagementPutObjectHandler = object_management.PutObjectHandlerFunc(handler.PutObject(objectService))

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(s *http.Server, scheme, addr string) {

}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
func configureGlobal() {

	// Parse configuration
	conf.ParseConfiguration(confFlags.ConfFile)

	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(conf.GetLogsLevel())
}
