package main_test

import (
	"net/http"
	"testing"
)

func TestGetOriginalURL(t *testing.T) {
	resp, err := http.Get("http://localhost:9000/v1/short/1")

	if http.StatusOK != resp.StatusCode {
		t.Errorf("Expected response code %d. Got %d\n", http.StatusOK, resp.StatusCode)
	}

	if err != nil {
		t.Error("Encountered an error:", err)
	}
}
