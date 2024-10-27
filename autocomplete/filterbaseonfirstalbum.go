package autocomplete

import (
	"net/http"
	"text/template"

	"groupie-tracker/handlers"
	"groupie-tracker/models"
)

func FirstAlbums(albumDate string, w http.ResponseWriter, r *http.Request) {
	// Initialize a slice to store bands filtered by the first album release date
	var filteredBands []models.Artist

	// Loop through the list of artists to find those whose first album matches the given albumDate
	for i := 0; i < len(Artists); i++ {
		if Artists[i].FirstAlbum == albumDate {
			filteredBands = append(filteredBands, Artists[i]) // Add matching artists to the filtered list
		}
	}

	if len(filteredBands) == 0 {
		http.Redirect(w, r, "/badrequest", http.StatusFound)
		return
	}

	// Set the title and description dynamically based on the album release date
	titlePage := "Debut Band(s) of " + albumDate
	pageDescription := "Explore the bands that made their mark with debut albums released on " + albumDate + ". Each of these talented groups embarked on their musical journey, shaping the sounds of their time."

	// Create a Creationbands struct to hold the title, description, and filtered artist data
	band := Creationbands{
		Title:    titlePage,
		Pagedesc: pageDescription,
		Mydata:   filteredBands,
	}

	// Check if the "band_search.html" file exists
	isFilePresent, _ := handlers.Checkfile("./", "band_search.html")
	if !isFilePresent {
		// If the file is not present, redirect to a 404 page
		http.Redirect(w, r, "/404", http.StatusFound)
		return
	}

	// Parse the "band_search.html" template file
	tmp, _ := template.ParseFiles("band_search.html")

	// Execute the template and pass in the band data to render the page
	tmp.Execute(w, band)
}
