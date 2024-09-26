package autocomplete

import (
	"net/http"
	"strconv"
	"text/template"

	"groupie-tracker/handlers"
	"groupie-tracker/models"
)
type Creationbands struct {
	Title    string          // Title of the page
	Pagedesc string          // Description of the page
	Mydata   []models.Artist  // Slice of Artist structs filtered by creation date
}

// CreationDate generates a page displaying bands created in a specified year
func CreationDate(yearStr string, w http.ResponseWriter, r *http.Request) {
	// Convert the year string (yearStr) to an integer
	num, _ := strconv.Atoi(yearStr)

	// Filter the list of artists based on the creation date matching the year
	var filteredBands []models.Artist
	for i := 0; i < len(Artists); i++ {
		if Artists[i].CreationDate == num {
			filteredBands = append(filteredBands, Artists[i]) // Add matching artists to the filtered list
		}
	}

	// Set the title and page description dynamically based on the year
	titlePage := "Iconic Bands Created in " + yearStr
	pageDescription := "Dive into the vibrant world of music from " + yearStr + "! Explore the legendary bands that emerged this year and discover how their unique sounds and styles shaped the music landscape, leaving an indelible mark on generations to come."

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
