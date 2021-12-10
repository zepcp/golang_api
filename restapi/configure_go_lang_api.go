// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"golang_api/models"
	"golang_api/restapi/operations"
	"golang_api/restapi/operations/books"
	"golang_api/restapi/operations/users"
)

//go:generate swagger generate server --target ../../golang_api --name GoLangAPI --spec ../api.yaml --principal interface{}

func configureFlags(api *operations.GoLangAPIAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.GoLangAPIAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	// Applies when the "Authorization" header is set
	if api.BearerAuth == nil {
		api.BearerAuth = func(token string) (interface{}, error) {
			return "ok", nil
			accessToken := strings.ReplaceAll(token, "Bearer ", "")
			accessToken = strings.TrimSpace(accessToken)

			if accessToken == "" {
				return nil, errors.New(http.StatusForbidden, "Missing Authorization")
			}
			tk := &models.Token{}

			_, err := jwt.ParseWithClaims(accessToken, tk, func(token *jwt.Token) (interface{}, error) {
				return []byte("secret"), nil
			})

			if err != nil {
				return nil, errors.New(http.StatusForbidden, "Invalid Authorization")
			}

			return tk.UserID, nil
		}
	}

	// Set your custom authorizer if needed. Default one is security.Authorized()
	// Expected interface runtime.Authorizer
	//
	// Example:
	// api.APIAuthorizer = security.Authorized()
	if api.BooksDeleteBooksIDHandler == nil {
		api.BooksDeleteBooksIDHandler = books.DeleteBooksIDHandlerFunc(func(params books.DeleteBooksIDParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation books.DeleteBooksID has not yet been implemented")
		})
	}
	if api.UsersDeleteUsersIDHandler == nil {
		api.UsersDeleteUsersIDHandler = users.DeleteUsersIDHandlerFunc(func(params users.DeleteUsersIDParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation users.DeleteUsersID has not yet been implemented")
		})
	}
	if api.BooksGetBooksHandler == nil {
		api.BooksGetBooksHandler = books.GetBooksHandlerFunc(func(params books.GetBooksParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation books.GetBooks has not yet been implemented")
		})
	}
	if api.BooksGetBooksIDHandler == nil {
		api.BooksGetBooksIDHandler = books.GetBooksIDHandlerFunc(func(params books.GetBooksIDParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation books.GetBooksID has not yet been implemented")
		})
	}
	if api.UsersGetUsersHandler == nil {
		api.UsersGetUsersHandler = users.GetUsersHandlerFunc(func(params users.GetUsersParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation users.GetUsers has not yet been implemented")
		})
	}
	if api.UsersGetUsersIDHandler == nil {
		api.UsersGetUsersIDHandler = users.GetUsersIDHandlerFunc(func(params users.GetUsersIDParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation users.GetUsersID has not yet been implemented")
		})
	}
	if api.BooksPatchBooksIDHandler == nil {
		api.BooksPatchBooksIDHandler = books.PatchBooksIDHandlerFunc(func(params books.PatchBooksIDParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation books.PatchBooksID has not yet been implemented")
		})
	}
	if api.UsersPatchUsersIDHandler == nil {
		api.UsersPatchUsersIDHandler = users.PatchUsersIDHandlerFunc(func(params users.PatchUsersIDParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation users.PatchUsersID has not yet been implemented")
		})
	}
	if api.BooksPostBooksHandler == nil {
		api.BooksPostBooksHandler = books.PostBooksHandlerFunc(func(params books.PostBooksParams) middleware.Responder {
			return middleware.NotImplemented("operation books.PostBooks has not yet been implemented")
		})
	}
	if api.UsersPostUsersHandler == nil {
		api.UsersPostUsersHandler = users.PostUsersHandlerFunc(func(params users.PostUsersParams) middleware.Responder {
			return middleware.NotImplemented("operation users.PostUsers has not yet been implemented")
		})
	}
	if api.UsersPostUsersLoginHandler == nil {
		api.UsersPostUsersLoginHandler = users.PostUsersLoginHandlerFunc(func(params users.PostUsersLoginParams) middleware.Responder {
			return middleware.NotImplemented("operation users.PostUsersLogin has not yet been implemented")
		})
	}

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
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
