package daemon // import "go.khulnasoft.com/integration/daemon"

import (
	"context"
	"os"
	"testing"

	"go.khulnasoft.com/testutil/environment"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"
)

var (
	testEnv     *environment.Execution
	baseContext context.Context
)

func TestMain(m *testing.M) {
	var err error

	ctx, span := otel.Tracer("").Start(context.Background(), "integration/daemon/TestMain")
	baseContext = ctx

	testEnv, err = environment.New(ctx)
	if err != nil {
		span.SetStatus(codes.Error, err.Error())
		span.End()
		panic(err)
	}
	err = environment.EnsureFrozenImagesLinux(ctx, testEnv)
	if err != nil {
		span.SetStatus(codes.Error, err.Error())
		span.End()
		panic(err)
	}

	testEnv.Print()

	code := m.Run()
	if code != 0 {
		span.SetStatus(codes.Error, "m.Run() exited with non-zero code")
	}
	os.Exit(code)
}
