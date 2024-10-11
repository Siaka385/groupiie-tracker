package autocomplete

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"sort"
	"strings"
)

type RequestData struct {
	Key string `json:"key"` // Struct to hold incoming search key data from the request body
}

// HandleSearchSuggestions processes incoming POST requests, validates, and redirects to the search results
func HandleSearchSuggestions(w http.ResponseWriter, r *http.Request) {
	// Only allow POST method for search suggestions
	if r.Method != "POST" {
		http.Error(w, "wrong method", http.StatusMethodNotAllowed) // Return error if method is not POST
		return
	}

	// Read the request body to retrieve the search key
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Unable to read request body", http.StatusBadRequest) // Handle error in reading request body
		return
	}

	// Unmarshal the body content into RequestData struct
	var item RequestData
	err = json.Unmarshal(body, &item)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest) // Handle error if the request body contains invalid JSON
		return
	}

	// Call function to generate search suggestions based on the search key
	GenerateSearchSuggestions(item.Key)

	// Redirect the client to the search results file (JSON)
	http.Redirect(w, r, "http://localhost:8089/js/search.json", http.StatusFound)
}

// GenerateSearchSuggestions generates search suggestions based on the user input and writes the suggestions to a file
func GenerateSearchSuggestions(m string) {
	var autocompleteitems []Suggestion

	// Loop through the suggestions and find those that match the search key
	for i := 0; i < len(SuggestionItems); i++ {
		if strings.Contains(strings.ToLower(SuggestionItems[i].Name), strings.ToLower(m)) {
			autocompleteitems = append(autocompleteitems, SuggestionItems[i])
		}
	}

	sort.Slice(autocompleteitems, func(i, j int) bool {
		return autocompleteitems[i].Name < autocompleteitems[j].Name
	})

	// Marshal the matching suggestions into JSON format
	jsondata, err := json.MarshalIndent(autocompleteitems, "", " ")
	if err != nil {
		fmt.Println("failed to marshal JSON:")
	}

	// Write the JSON data into the search.json file
	err = os.WriteFile("./js/search.json", jsondata, 0o644)
	if err != nil {
		fmt.Println("failed to save JSON to file:")
	}
}
