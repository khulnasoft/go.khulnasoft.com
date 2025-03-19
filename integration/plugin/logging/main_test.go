package logging // import "go.khulnasoft.com/integration/plugin/logging"

import (
	"context"
	"os"
	"testing"

	"go.khulnasoft.com/testutil"
	"go.khulnasoft.com/testutil/environment"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"
)

var (
	testEnv     *environment.Execution
	baseContext context.Context
)

func TestMain(m *testing.M) {
	shutdown := testutil.ConfigureTracing()
	ctx, span := otel.Tracer("").Start(context.Background(), "integration/plugin/logging.TestMain")
	baseContext = ctx

	var err error
	testEnv, err = environment.New(ctx)
	if err != nil {
		span.SetStatus(codes.Error, err.Error())
		span.End()
		shutdown(ctx)
		panic(err)
	}
	err = environment.EnsureFrozenImagesLinux(ctx, testEnv)
	if err != nil {
		span.SetStatus(codes.Error, err.Error())
		span.End()
		shutdown(ctx)
		panic(err)
	}

	testEnv.Print()
	code := m.Run()
	span.End()
	if code != 0 {
		span.SetStatus(codes.Error, "m.Run() exited with non-zero code")
	}
	shutdown(ctx)
	os.Exit(code)
}
