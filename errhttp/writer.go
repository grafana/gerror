package errhttp

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"reflect"

	"github.com/grafana/gerror"
)

var ErrNonGrafanaError = gerror.New(gerror.StatusInternal, "core.MalformedError")
var ErrWritingError = gerror.New(gerror.StatusInternal, "core.ErrorWritingError")

// ErrorOptions is a container for functional options passed to [Write].
type ErrorOptions struct {
	fallback *gerror.Error
	callback func(gerror.Error)
}

// Write writes an error to the provided [http.ResponseWriter] with the
// appropriate HTTP status and JSON payload from [gerror.Error].
// Write also logs the provided error to either the "request-errors"
// logger, or the logger provided as a functional option using
// [WithLogger].
// When passing errors that are not [errors.As] compatible with
// [gerror.Error], [ErrNonGrafanaError] will be used to create a
// generic 500 Internal Server Error payload by default, this is
// overrideable by providing [WithFallback] for a custom fallback
// error.
func Write(ctx context.Context, err error, w http.ResponseWriter, opts ...func(ErrorOptions) ErrorOptions) error {
	opt := ErrorOptions{}
	for _, o := range opts {
		opt = o(opt)
	}

	var gerr gerror.Error
	if !errors.As(err, &gerr) {
		gerr = fallbackOrInternalError(err, opt)
	}

	if opt.callback != nil {
		opt.callback(gerr)
	}

	pub := gerr.Public()
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(pub.StatusCode)
	err = json.NewEncoder(w).Encode(pub)
	if err != nil {
		return ErrWritingError.Errorf("error writing error: %w", err)
	}
	return nil
}

// WithFallback sets the default error returned to the user if the error
// sent to [Write] is not an [gerror.Error].
func WithFallback(opt ErrorOptions, fallback gerror.Error) ErrorOptions {
	opt.fallback = &fallback
	return opt
}

// WithCallback sets a function that sends a [gerror.Error] that contain
// information about the error that will be sent to the
// [http.ResponseWriter] from [Write].
func WithCallback(opt ErrorOptions, callback func(gerror.Error)) ErrorOptions {
	opt.callback = callback
	return opt
}

func fallbackOrInternalError(err error, opt ErrorOptions) gerror.Error {
	if opt.fallback != nil {
		fErr := *opt.fallback
		fErr.Underlying = err
		return fErr
	}

	return ErrNonGrafanaError.Errorf("unexpected error type [%s]: %w", reflect.TypeOf(err), err)
}
