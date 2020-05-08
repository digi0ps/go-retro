package server

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPing(t *testing.T) {
	t.Run("Should be ok and return pong", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/ping", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(Ping)

		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Error("Ping handler returned wrong status")
		}

		if expected := "Pong\n"; rr.Body.String() != expected {
			t.Error("Ping handler return wrong body")
		}
	})
}
