package registry // import "go.khulnasoft.com/registry"

import (
	"net/url"

	"github.com/khulnasoft/distribution/registry/api/errcode"
	"go.khulnasoft.com/errdefs"
	"github.com/pkg/errors"
)

func translateV2AuthError(err error) error {
	switch e := err.(type) {
	case *url.Error:
		switch e2 := e.Err.(type) {
		case errcode.Error:
			switch e2.Code {
			case errcode.ErrorCodeUnauthorized:
				return errdefs.Unauthorized(err)
			}
		}
	}

	return err
}

func invalidParam(err error) error {
	return errdefs.InvalidParameter(err)
}

func invalidParamf(format string, args ...interface{}) error {
	return errdefs.InvalidParameter(errors.Errorf(format, args...))
}

func invalidParamWrapf(err error, format string, args ...interface{}) error {
	return errdefs.InvalidParameter(errors.Wrapf(err, format, args...))
}
