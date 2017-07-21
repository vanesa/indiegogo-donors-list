package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestIndiegogoEndpoint(t *testing.T) {
	// Create a request to pass to our handler.
	req := httptest.NewRequest("GET", "http://google.com", nil)

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()

	handler(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
