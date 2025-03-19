package server // import "go.khulnasoft.com/api/server"

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"go.khulnasoft.com/api"
	"go.khulnasoft.com/api/server/httputils"
	"go.khulnasoft.com/api/server/middleware"
)

func TestMiddlewares(t *testing.T) {
	srv := &Server{}

	m, err := middleware.NewVersionMiddleware("0.1omega2", api.DefaultVersion, api.MinSupportedAPIVersion)
	if err != nil {
		t.Fatal(err)
	}
	srv.UseMiddleware(*m)

	req, _ := http.NewRequest(http.MethodGet, "/containers/json", nil)
	resp := httptest.NewRecorder()
	ctx := context.Background()

	localHandler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, vars map[string]string) error {
		if httputils.VersionFromContext(ctx) == "" {
			t.Fatal("Expected version, got empty string")
		}

		if sv := w.Header().Get("Server"); !strings.Contains(sv, "Docker/0.1omega2") {
			t.Fatalf("Expected server version in the header `Docker/0.1omega2`, got %s", sv)
		}

		return nil
	}

	handlerFunc := srv.handlerWithGlobalMiddlewares(localHandler)
	if err := handlerFunc(ctx, resp, req, map[string]string{}); err != nil {
		t.Fatal(err)
	}
}
