package server // import "go.khulnasoft.com/api/server"

import (
	"github.com/containerd/log"
	"go.khulnasoft.com/api/server/httputils"
	"go.khulnasoft.com/api/server/middleware"
)

// handlerWithGlobalMiddlewares wraps the handler function for a request with
// the server's global middlewares. The order of the middlewares is backwards,
// meaning that the first in the list will be evaluated last.
func (s *Server) handlerWithGlobalMiddlewares(handler httputils.APIFunc) httputils.APIFunc {
	next := handler

	for _, m := range s.middlewares {
		next = m.WrapHandler(next)
	}

	if log.GetLevel() == log.DebugLevel {
		next = middleware.DebugRequestMiddleware(next)
	}

	return next
}
