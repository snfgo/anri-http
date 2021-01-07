package helpers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCorsHelper(t *testing.T) {
	assert := assert.New(t)

	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		allowedOrigin := w.Header().Get("Access-Control-Allow-Origin")
		allowedMethods := w.Header().Get("Access-Control-Allow-Methods")
		allowedHeaders := w.Header().Get("Access-Control-Allow-Headers")

		assert.Equal(allowedOrigin, CorsOrigin)
		assert.Equal(allowedMethods, CorsMethods)
		assert.Equal(allowedHeaders, CorsHeaders)
	})

	testHandler := Cors(nextHandler)

	req := httptest.NewRequest("GET", "http://testing", nil)

	testHandler.ServeHTTP(httptest.NewRecorder(), req)
}
