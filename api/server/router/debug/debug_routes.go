package debug // import "go.khulnasoft.com/api/server/router/debug"

import (
	"context"
	"net/http"
	"net/http/pprof"
)

func handlePprof(ctx context.Context, w http.ResponseWriter, r *http.Request, vars map[string]string) error {
	pprof.Handler(vars["name"]).ServeHTTP(w, r)
	return nil
}
