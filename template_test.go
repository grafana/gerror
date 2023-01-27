package gerror_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/grafana/gerror"
)

func TestTemplate(t *testing.T) {
	tmpl := gerror.New(gerror.StatusInternal, "template.sampleError").MustTemplate("[{{ .Public.user }}] got error: {{ .Error }}")
	err := tmpl.Build(gerror.TemplateData{
		Public: map[string]any{
			"user": "grot the bot",
		},
		Error: errors.New("oh noes"),
	})

	t.Run("Built error should return true when compared with templated error ", func(t *testing.T) {
		require.True(t, errors.Is(err, tmpl))
	})

	t.Run("Built error should return true when compared with templated error base ", func(t *testing.T) {
		require.True(t, errors.Is(err, tmpl.Base))
	})
}

func ExampleTemplate() {
	// Initialization, this is typically done on a package or global
	// level.
	var tmpl = gerror.New(gerror.StatusInternal, "template.sampleError").MustTemplate("[{{ .Public.user }}] got error: {{ .Error }}")

	// Construct an error based on the template.
	err := tmpl.Build(gerror.TemplateData{
		Public: map[string]any{
			"user": "grot the bot",
		},
		Error: errors.New("oh noes"),
	})

	fmt.Println(err.Error())

	// Output:
	// [template.sampleError] [grot the bot] got error: oh noes
}

func ExampleTemplate_public() {
	// Initialization, this is typically done on a package or global
	// level.
	var tmpl = gerror.New(gerror.StatusInternal, "template.sampleError").
		MustTemplate(
			"[{{ .Public.user }}] got error: {{ .Error }}",
			gerror.WithPublic("Oh, no, error for {{ .Public.user }}"),
		)

	// Construct an error based on the template.
	//nolint:errorlint
	err := tmpl.Build(gerror.TemplateData{
		Public: map[string]any{
			"user": "grot the bot",
		},
		Error: errors.New("oh noes"),
	}).(gerror.Error)

	fmt.Println(err.Error())
	fmt.Println(err.PublicMessage)

	// Output:
	// [template.sampleError] [grot the bot] got error: oh noes
	// Oh, no, error for grot the bot
}
