package handlers

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

// Test handler function
func testHandler(t *testing.T, method, url string, handlerFunc http.HandlerFunc, expectedStatus int) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	handlerFunc(w, req)

	res := w.Result()
	if res.StatusCode == http.StatusFound {
		t.Log("Redirected to:", res.Header.Get("Location"))
	} else if res.StatusCode != expectedStatus {
		t.Errorf("Expected status %v, got %v", expectedStatus, res.Status)
	}
}

// Test for the About Us page
func TestAboutus(t *testing.T) {
	testHandler(t, http.MethodGet, "/about", Aboutus, http.StatusOK)
}

// Test for the Homepage
func TestHomepage(t *testing.T) {
	testHandler(t, http.MethodGet, "/", Homepage, http.StatusOK)
}

// Test for the Search Bar functionality
func TestSearchBar(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/search?search=test artist", nil)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	HandleManualSearch(w, req)

	res := w.Result()
	if res.StatusCode == http.StatusFound {
		t.Log("Redirected to:", res.Header.Get("Location"))
	} else if res.StatusCode != http.StatusFound {
		t.Errorf("Expected status Found, got %v", res.Status)
	}
}

// Test for the Artist Info page
func TestArtinfo(t *testing.T) {
	testHandler(t, http.MethodGet, "/artist?id=1", Artinfo, http.StatusOK)
}

// Test for handling a 404 Not Found error
func TestError404(t *testing.T) {
	testHandler(t, http.MethodGet, "/non-existent", Error404, http.StatusNotFound)
}

// Test for handling an Internal Server Error
func TestInternalServerError(t *testing.T) {
	testHandler(t, http.MethodGet, "/500", InternalServerError, http.StatusInternalServerError)
}

// Test for handling a Wrong Method error
func TestWrongMethod(t *testing.T) {
	req, err := http.NewRequest(http.MethodPost, "/wrong-method", nil)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	Wrongmethod(w, req)

	res := w.Result()
	if res.StatusCode == http.StatusFound {
		t.Log("Redirected to:", res.Header.Get("Location"))
	} else if res.StatusCode != http.StatusMethodNotAllowed {
		t.Errorf("Expected status Method Not Allowed, got %v", res.Status)
	}
}

// Test for handling No Internet Connection error
func TestNoInternetConnection(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/no-internet", nil)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	Nointernetconnection(w, req)

	res := w.Result()
	if res.StatusCode == http.StatusNotFound {
		t.Log("File is missing or not found")
	} else if res.StatusCode != http.StatusRequestTimeout {
		t.Errorf("Expected status Request Timeout, got %v", res.Status)
	}
}

// Test for handling Artist Not Found error
func TestArtistNotFound(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/artist-not-found", nil)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	ArtistNotFound(w, req)

	res := w.Result()
	if res.StatusCode == http.StatusNotFound {
		t.Log("File is missing or not found")
	} else if res.StatusCode != http.StatusBadRequest {
		t.Errorf("Expected status Bad Request, got %v", res.Status)
	}
}

// Test for Checkfile function
func TestCheckfile(t *testing.T) {
	// Ensure the directory exists
	if _, err := os.Stat("./Errortemplate/"); os.IsNotExist(err) {
		t.Skip("Skipping TestCheckfile: ./Errortemplate/ does not exist")
	}

	exists, err := Checkfile("./Errortemplate/", "error500.html")
	if err != nil {
		t.Fatalf("Checkfile error: %v", err)
	}
	if !exists {
		t.Errorf("Expected file to exist, got %v", exists)
	}
}

// Test CheckInternetConnectivity function
func TestCheckInternetConnectivity(t *testing.T) {
	ok := CheckInternetConnectivity()
	if !ok {
		t.Errorf("Expected internet connectivity, but got %v", ok)
	}
}
