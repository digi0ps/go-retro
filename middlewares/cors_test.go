package middlewares

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCorsMiddleware(t *testing.T) {
	assert := assert.New(t)

	t.Run("should execute the passed handler func", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/api/board", nil)
		recorder := httptest.NewRecorder()

		handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, "hello world")
		})
		corsHandler := CorsMiddleware(handler)

		corsHandler.ServeHTTP(recorder, req)
		assert.Equal("hello world", recorder.Body.String())
	})

	t.Run("should send back CORS headers", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodPut, "/api/board", nil)
		recorder := httptest.NewRecorder()

		handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, "hello world 2")
		})
		corsHandler := CorsMiddleware(handler)

		corsHandler.ServeHTTP(recorder, req)
		assert.Equal("hello world 2", recorder.Body.String())

		assert.Equal("*", recorder.Header().Get("Access-Control-Allow-Origin"))
		assert.Equal("GET, PUT, OPTIONS", recorder.Header().Get("Access-Control-Allow-Methods"))
	})

	t.Run("should send back CORS headers but not execute handler for OPTIONS", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodOptions, "/api/board", nil)
		recorder := httptest.NewRecorder()

		isCalled := false
		handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			isCalled = true
		})
		corsHandler := CorsMiddleware(handler)

		corsHandler.ServeHTTP(recorder, req)

		assert.Equal(false, isCalled)
		assert.Equal("*", recorder.Header().Get("Access-Control-Allow-Origin"))
		assert.Equal("GET, PUT, OPTIONS", recorder.Header().Get("Access-Control-Allow-Methods"))
	})
}
