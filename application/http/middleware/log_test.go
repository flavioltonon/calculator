package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"calculator/application/http/router"
	mocks "calculator/mocks/shared/logging"

	"github.com/stretchr/testify/mock"
)

func TestLog(t *testing.T) {
	logger := mocks.NewLogger(t)
	logger.On("Info", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Once()
	logger.On("Info", mock.Anything, mock.Anything, mock.Anything).Once()

	middleware := Log(logger)

	t.Run("Should log once when the request is received and again when a response is being returned", func(t *testing.T) {
		middleware(router.NopHandler).ServeHTTP(httptest.NewRecorder(), httptest.NewRequest(http.MethodPost, "/some-url", nil))
	})
}
