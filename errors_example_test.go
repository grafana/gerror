package gerror_test

import (
	"context"
	"errors"
	"fmt"
	"path"
	"strings"

	"github.com/grafana/gerror"
)

var (
	// define the set of errors which should be presented using the
	// same error message for the frontend statically within the
	// package.

	errAbsPath     = gerror.New(gerror.StatusBadRequest, "shorturl.absolutePath")
	errInvalidPath = gerror.New(gerror.StatusBadRequest, "shorturl.invalidPath")
	errUnexpected  = gerror.New(gerror.StatusInternal, "shorturl.unexpected")
)

func Example() {
	var e gerror.Error

	_, err := CreateShortURL("abc/../def")
	errors.As(err, &e)
	fmt.Println(e.Reason.HTTPStatus(), e.MessageID)
	fmt.Println(e.Error())

	// Output:
	// 400 shorturl.invalidPath
	// [shorturl.invalidPath] path mustn't contain '..': 'abc/../def'
}

// CreateShortURL runs a few validations and returns
// 'https://example.org/s/tretton' if they all pass. It's not a very
// useful function, but it shows errors in a semi-realistic function.
func CreateShortURL(longURL string) (string, error) {
	if path.IsAbs(longURL) {
		return "", errAbsPath.Errorf("unexpected absolute path")
	}
	if strings.Contains(longURL, "../") {
		return "", errInvalidPath.Errorf("path mustn't contain '..': '%s'", longURL)
	}
	if strings.Contains(longURL, "@") {
		return "", errInvalidPath.Errorf("cannot shorten email addresses")
	}

	shortURL, err := createShortURL(context.Background(), longURL)
	if err != nil {
		return "", errUnexpected.Errorf("failed to create short URL: %w", err)
	}

	return shortURL, nil
}

func createShortURL(_ context.Context, _ string) (string, error) {
	return "https://example.org/s/tretton", nil
}
