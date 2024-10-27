package autocomplete

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"groupie-tracker/models"
)

func setup() {
	// Create the directory for the JSON file if it doesn't exist
	os.MkdirAll("./js", os.ModePerm)

	// Mock suggestions for testing
	SuggestionItems = []Suggestion{
		{Name: "The Beatles", Contents: "artist/band", Id: []int{1}},
		{Name: "John Lennon", Contents: "member", Id: []int{2}},
		{Name: "USA", Contents: "location", Id: []int{1}},
		{Name: "UK", Contents: "location", Id: []int{2}},
	}

	// Mock artists data
	Artists = []models.Artist{
		{ID: 1, Name: "The Beatles", CreationDate: 1960, FirstAlbum: "22-03-1963"},
		{ID: 2, Name: "John Lennon", CreationDate: 1956, FirstAlbum: "12-10-1970"},
	}

	// Ensure the search.json file exists for the test
	os.WriteFile("./js/search.json", []byte("[]"), 0o644)
}

func TestHandleSearchSuggestions(t *testing.T) {
	setup()
	// Prepare request with a search key
	requestBody, _ := json.Marshal(RequestData{Key: "The"})
	req, err := http.NewRequest("POST", "/searchy", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}

	// Record the response
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HandleSearchSuggestions)
	handler.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusFound {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusFound)
	}

	// Check if redirect URL is correct
	expectedLocation := "http://localhost:8089/js/search.json"
	if rr.Header().Get("Location") != expectedLocation {
		t.Errorf("Handler returned wrong redirect location: got %v want %v", rr.Header().Get("Location"), expectedLocation)
	}
}

func TestCheckIfCreationdate(t *testing.T) {
	// Test valid creation date
	if !CheckIfCreationdate("1960") {
		t.Error("Expected true for valid creation date")
	}

	// Test invalid creation date
	if CheckIfCreationdate("invalid-date") {
		t.Error("Expected false for invalid creation date")
	}
}

func TestCheckFirstAlbum(t *testing.T) {
	// Test valid album release date
	if !CheckFirstAlbum("22-03-1963") {
		t.Error("Expected true for valid album release date")
	}

	// Test invalid album release date
	if CheckFirstAlbum("invalid-date") {
		t.Error("Expected false for invalid album release date")
	}
}

func TestHandleAutocompleteSelection(t *testing.T) {
	setup()

	// Test selecting by creation date
	req, err := http.NewRequest("GET", "/searchy?search=1960", nil)
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HandleAutocompleteSelection)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK && rr.Code != http.StatusFound {
		t.Errorf("Expected status 200 or 302 but got %v", rr.Code)
	}
}
