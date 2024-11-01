package autocomplete

import (
	"net/http"
	"text/template"

	"groupie-tracker/handlers"
	"groupie-tracker/models"
)

func Locations(location string, w http.ResponseWriter, r *http.Request) {
	// Initialize a slice to store bands filtered by location
	var filteredBands []models.Artist
	var ids []int // Slice to store artist IDs that match the location

	// Loop through the SuggestionItems to find the matching location
	for i := 0; i < len(SuggestionItems); i++ {
		if SuggestionItems[i].Contents == "location" { // Check if the suggestion is a location
			if SuggestionItems[i].Name == location { // If the location matches the search term
				// Collect all the artist IDs associated with the location
				for k := 0; k < len(SuggestionItems[i].Id); k++ {
					ids = append(ids, SuggestionItems[i].Id[k])
				}
				break // Stop searching once the location is found
			}
		}
	}

	if len(ids) == 0 {
		http.Redirect(w, r, "/notfound", http.StatusFound)
		return
	}

	// Loop through the artist IDs and match them with the corresponding artists
	for h := 0; h < len(ids); h++ {
		for j := 0; j < len(Artists); j++ {
			if Artists[j].ID == ids[h] { // If the artist ID matches, add the artist to the filtered list
				filteredBands = append(filteredBands, Artists[j])
			}
		}
	}

	// Set the title and description dynamically based on the location
	titlePage := "Bands That Played in " + location
	pageDescription := "Join us as we delve into the rich musical legacy of bands that have graced the stage in " + location + ". Discover the influential sounds, unforgettable performances, and the unique stories behind the artists who made their mark on this vibrant music scene."

	// Create a Creationbands struct to hold the title, description, and filtered artist data
	band := Creationbands{
		Title:    titlePage,
		Pagedesc: pageDescription,
		Mydata:   filteredBands,
	}

	// Check if the "band_search.html" file exists
	isFilePresent, _ := handlers.Checkfile("./", "band_search.html")
	if !isFilePresent {
		// If the file is not present, redirect to a 50 page
		Error500(w, r)
		return
	}

	// Parse the "band_search.html" template file
	tmp, _ := template.ParseFiles("band_search.html")

	// Execute the template and pass in the band data to render the page
	tmp.Execute(w, band)
}
