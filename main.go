package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"groupie-tracker/autocomplete"
	"groupie-tracker/handlers"
)

func router(w http.ResponseWriter, r *http.Request) {
	// Check internet connectivity; if not connected, handle the "no internet" response.
	if !handlers.CheckInternetConnectivity() {
		handlers.ErrorRenderPage(w, r, http.StatusRequestTimeout, "Errortemplate/internetconnection.html")
		return
	}

	if autocomplete.FetchingError {
		// Serve the internal server error page for a path indicating a server error.
		handlers.InternalServerError(w, r)
		return
	}

	// Serve the homepage or explore page based on the URL path.
	if r.URL.Path == "/" || r.URL.Path == "/explore" {
		handlers.Homepage(w, r)
	} else if r.URL.Path == "/artist" {
		// Serve the artist information page when the path is "/artist".
		handlers.Artinfo(w, r)
	} else if r.URL.Path == "/500" {
		// Serve the internal server error page for a path indicating a server error.
		handlers.InternalServerError(w, r)
	} else if r.URL.Path == "/wrongmethod" {
		// Handle requests made using an incorrect HTTP method.
		handlers.ErrorRenderPage(w, r, http.StatusMethodNotAllowed, "Errortemplate/wrongmethodused.html")
	} else if r.URL.Path == "/about" {
		// Serve the "About Us" page when the path is "/about".
		handlers.Aboutus(w, r)
	} else if r.URL.Path == "/badrequest" {
		// Handle cases where the artist is not found (bad request).
		handlers.ErrorRenderPage(w, r, http.StatusBadRequest, "Errortemplate/Noaristfound.html")
	} else if r.URL.Path == "/serch" {
		// Handle artist autocomplete selection when a user clicks on a suggestion.
		autocomplete.HandleAutocompleteSelection(w, r)
	} else if r.URL.Path == "/searchy" {
		// Handle search suggestions for artists based on partial input from the user.
		autocomplete.HandleSearchSuggestions(w, r)
	} else if r.URL.Path == "/searchresult" {
		autocomplete.SearchPageHandler(w, r)
	} else {
		// Serve a 404 error page for unrecognized routes.
		handlers.ErrorRenderPage(w, r, http.StatusNotFound, "Errortemplate/error.html")
	}
}

func main() {
	if len(os.Args) != 1 {
		fmt.Println("Usage: go run .")
		os.Exit(1)
	}

	mux := http.NewServeMux()

	autocomplete.GenerateSuggestions()

	mux.HandleFunc("/css/", handlers.StaticServer)
	mux.HandleFunc("/fonts/", handlers.StaticServer)
	mux.HandleFunc("/images/", handlers.StaticServer)
	mux.HandleFunc("/js/", handlers.StaticServer)

	// Set up the HTTP server and route
	mux.HandleFunc("/", router)

	// Start the server on port 8080
	fmt.Println("Server is running on port 8089...")
	if err := http.ListenAndServe(":8089", mux); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
