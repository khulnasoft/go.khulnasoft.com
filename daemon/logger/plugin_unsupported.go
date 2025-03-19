//go:build !linux && !freebsd

package logger // import "go.khulnasoft.com/daemon/logger"

import (
	"errors"
	"io"
)

func openPluginStream(a *pluginAdapter) (io.WriteCloser, error) {
	return nil, errors.New("log plugin not supported")
}
