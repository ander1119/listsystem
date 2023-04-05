package utils

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRespondWithError(t *testing.T) {
	// create a fake HTTP response writer
	w := httptest.NewRecorder()

	// call the function with a sample error message
	RespondWithError(w, http.StatusBadRequest, "Invalid input")

	// check that the response has the expected status code
	if w.Code != http.StatusBadRequest {
		t.Errorf("Unexpected status code: got %d, want %d", w.Code, http.StatusBadRequest)
	}

	// decode the response body to a map
	var response map[string]string
	if err := json.NewDecoder(w.Body).Decode(&response); err != nil {
		t.Errorf("Error decoding response body: %v", err)
	}

	// check that the response body has the expected error message
	expectedMessage := "Invalid input"
	if response["error"] != expectedMessage {
		t.Errorf("Unexpected error message: got %s, want %s", response["error"], expectedMessage)
	}
}

func TestRespondWithJSON(t *testing.T) {
	// create a fake HTTP response writer
	w := httptest.NewRecorder()

	// define a sample payload
	type TestHead struct {
		ListKey string
		NextPageKey string
	}
	payload := TestHead{"123456789", "abcdefg"}

	// call the function with the sample payload
	respondWithJSON(w, http.StatusOK, payload)

	// check that the response has the expected status code
	if w.Code != http.StatusOK {
		t.Errorf("Unexpected status code: got %d, want %d", w.Code, http.StatusOK)
	}

	// decode the response body to a map
	var response TestHead
	if err := json.NewDecoder(w.Body).Decode(&response); err != nil {
		t.Errorf("Error decoding response body: %v", err)
	}

	// check that the response body has the expected payload
	expectedListKey := "123456789"
	if response.ListKey != expectedListKey {
		t.Errorf("Unexpected name: got %s, want %s", response.ListKey, expectedListKey)
	}
	expectedNextPageKey := "abcdefg"
	if response.NextPageKey != expectedNextPageKey {
		t.Errorf("Unexpected age: got %s, want %s", response.NextPageKey, expectedNextPageKey)
	}
}