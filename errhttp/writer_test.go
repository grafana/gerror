package errhttp

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/grafana/gerror"
)

func TestWrite(t *testing.T) {
	ctx := context.Background()
	const msgID = "test.thisIsExpected"
	base := gerror.New(gerror.StatusTimeout, msgID)
	handler := func(writer http.ResponseWriter, request *http.Request) error {
		return Write(ctx, base.Errorf("got expected error"), writer)
	}

	req := httptest.NewRequest("GET", "http://localhost:3000/fake", nil)
	recorder := httptest.NewRecorder()

	ret := handler(recorder, req)

	assert.Equal(t, http.StatusGatewayTimeout, recorder.Code)
	assert.JSONEq(t, `{"message": "Timeout", "messageId": "test.thisIsExpected", "statusCode": 504}`, recorder.Body.String())
	assert.NoError(t, ret)
}
