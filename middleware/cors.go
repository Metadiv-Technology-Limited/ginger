package middleware

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type CorsConfig struct {
	AllowAllOrigins bool

	// AllowOrigins is a list of origins a cross-domain request can be executed from.
	// If the special "*" value is present in the list, all origins will be allowed.
	// Default value is []
	AllowOrigins []string

	// AllowOriginFunc is a custom function to validate the origin. It takes the origin
	// as an argument and returns true if allowed or false otherwise. If this option is
	// set, the content of AllowOrigins is ignored.
	AllowOriginFunc func(origin string) bool

	// AllowMethods is a list of methods the client is allowed to use with
	// cross-domain requests. Default value is simple methods (GET, POST, PUT, PATCH, DELETE, HEAD, and OPTIONS)
	AllowMethods []string

	// AllowPrivateNetwork indicates whether the response should include allow private network header
	AllowPrivateNetwork bool

	// AllowHeaders is list of non simple headers the client is allowed to use with
	// cross-domain requests.
	AllowHeaders []string

	// AllowCredentials indicates whether the request can include user credentials like
	// cookies, HTTP authentication or client side SSL certificates.
	AllowCredentials bool

	// ExposeHeaders indicates which headers are safe to expose to the API of a CORS
	// API specification
	ExposeHeaders []string

	// MaxAge indicates how long (with second-precision) the results of a preflight request
	// can be cached
	MaxAge time.Duration

	// Allows to add origins like http://some-domain/*, https://api.* or http://some.*.subdomain.com
	AllowWildcard bool

	// Allows usage of popular browser extensions schemas
	AllowBrowserExtensions bool

	// Allows usage of WebSocket protocol
	AllowWebSockets bool

	// Allows usage of file:// schema (dangerous!) use it only when you 100% sure it's needed
	AllowFiles bool

	// Allows to pass custom OPTIONS response status code for old browsers / clients
	OptionsResponseStatusCode int
}

/*
CORS setup cors middleware
*/
func CORS(config CorsConfig) gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowAllOrigins:           config.AllowAllOrigins,
		AllowOrigins:              config.AllowOrigins,
		AllowOriginFunc:           config.AllowOriginFunc,
		AllowMethods:              config.AllowMethods,
		AllowHeaders:              config.AllowHeaders,
		AllowCredentials:          config.AllowCredentials,
		ExposeHeaders:             config.ExposeHeaders,
		MaxAge:                    config.MaxAge,
		AllowWildcard:             config.AllowWildcard,
		AllowBrowserExtensions:    config.AllowBrowserExtensions,
		AllowWebSockets:           config.AllowWebSockets,
		AllowFiles:                config.AllowFiles,
		OptionsResponseStatusCode: config.OptionsResponseStatusCode,
	})
}
