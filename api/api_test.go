package api

import (
	"bytes"
	"io"
	"net/http"
	"testing"
)

// wrapper that accepts httpGet for mocking
func fetchAPIDataWithCustomGet(httpGet func(url string) (*http.Response, error), url string) ([]byte, error) {
	resp, err := httpGet(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func TestFetchAPIDataSuccess(t *testing.T) {
	url := "http://example.com"
	expectedBody := []byte("response body")

	// Mock the http.Get function
	httpGet := func(url string) (*http.Response, error) {
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(bytes.NewReader(expectedBody)),
		}, nil
	}

	// Use the custom version for testing
	body, err := fetchAPIDataWithCustomGet(httpGet, url)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if !bytes.Equal(body, expectedBody) {
		t.Fatalf("expected %s, got %s", expectedBody, body)
	}
}

// Fetches artist data successfully from the API
func TestFetchArtistsSuccess(t *testing.T) {
	expectedBody := 1

	artists, err := FetchArtists()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if artists[0].ID != expectedBody {
		t.Fatalf("unexpected artist names: %v", artists)
	}
}

func TestLocationSuccess(t *testing.T) {
	expectedBody := "north_carolina-usa"

	response, err := FetchLocations()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if response[0].Locations[0] != expectedBody {
		t.Fatalf("unexpected artist names: %v", response)
	}
}

func TestDateSuccess(t *testing.T) {
	expectedBody := "*23-08-2019"

	response, err := FetchDates()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if response[0].Dates[0] != expectedBody {
		t.Fatalf("unexpected artist names: %v", response)
	}
}

func TestRelationSuccess(t *testing.T) {
	// Slice of keys directly
	keys := []string{
		"dunedin-new_zealand",
		"georgia-usa",
		"los_angeles-usa",
		"nagoya-japan",
		"north_carolina-usa",
		"osaka-japan",
		"penrose-new_zealand",
		"saitama-japan",
	}

	response, err := FetchRelations()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	for i := 0; i < len(keys); i++ {
		if _, ok := response[0].DatesLocations[keys[i]]; !ok {
			t.Fatalf("unexpected artist names: %v", response)
		}
	}
}
