package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"

	"groupie-tracker/api"
	"groupie-tracker/models"
)

func Artinfo(w http.ResponseWriter, r *http.Request) {
	// Ensure the method is GET
	if r.Method != http.MethodGet {
		http.Redirect(w, r, "/wrongmethod", http.StatusMethodNotAllowed)
		return
	}

	// Get the artist ID from the query string
	name := r.URL.Query().Get("id")
	id, err := strconv.Atoi(name)
	if err != nil {
		http.Redirect(w, r, "/badrequest", http.StatusFound)
		return
	}

	// Fetch the list of artists
	artists, err := api.FetchArtists()
	if err != nil {
		fmt.Println("Error fetching artists:", err)
		http.Redirect(w, r, "/500", http.StatusFound)
		return
	}

	// Find the artist with the specified ID
	var artistInfo models.Artist
	found := false
	for _, musicArtist := range artists {
		if musicArtist.ID == id {
			artistInfo = musicArtist
			found = true
			break
		}
	}

	// If no artist found, redirect to 404
	if !found {
		http.Redirect(w, r, "/404error", http.StatusFound)
		return
	}

	// Fetch and set the artist's timeline (relations)
	artistInfo.Relation = Timeline(id)
	if artistInfo.Relation == nil {
		http.Redirect(w, r, "/500", http.StatusFound)
		return
	}
	artistInfo.Locate = Locations(id)
	if artistInfo.Locate == nil {
		http.Redirect(w, r, "/500", http.StatusFound)
		return
	}
	artistInfo.Datess = Dates(id)
	if artistInfo.Datess == nil {
		http.Redirect(w, r, "/500", http.StatusFound)
		return
	}

	// Parse and execute the template with the artist information
	tmpl, err := template.ParseFiles("artistinformation.html")
	if err != nil {
		http.Redirect(w, r, "/500", http.StatusFound)
		return
	}

	// Execute the template (no redirects after this point)
	if err := tmpl.Execute(w, artistInfo); err != nil {
		http.Redirect(w, r, "/500", http.StatusFound)
	}
}

// Create a new artist info object from the artist info template and return it as a string in the format
func Locations(id int) []string {
	// Fetch the relations data
	relations, err := api.FetchLocations()
	if err != nil {
		fmt.Println("Error fetching relations:", err)
		return nil // added error handler for error handling here
	}
	// Find the relation for the specified artist ID
	var rel models.Location
	for _, relation := range relations {
		if relation.ID == id {
			rel = relation
			break
		}
	}

	return rel.Locations
}

func Dates(id int) []string {
	// Fetch the relations data
	relations, _ := api.FetchDates()
	var rel models.Date
	for _, relation := range relations {
		if relation.ID == id {
			rel = relation
			break
		}
	}
	sortdatess(rel.Dates)
	var removestar []string
	for i := 0; i < len(rel.Dates); i++ {
		removestar = append(removestar, strings.ReplaceAll(rel.Dates[i], "*", ""))
	}

	return removestar
}

func Timeline(id int) [][]string {
	// Fetch the relations data
	relations, err := api.FetchRelations()
	if err != nil {
		fmt.Println("Error fetching relations:", err)
		return nil
	}

	// Find the relation for the specified artist ID
	var rel models.Relation
	for _, relation := range relations {
		if relation.ID == id {
			rel = relation
			break
		}
	}

	// Construct the date-location pairs
	var motherSlice [][]string
	for key, values := range rel.DatesLocations {
		for _, cat := range values {
			motherSlice = append(motherSlice, []string{key, cat})
		}
	}

	// Sort the dates
	SortDate(motherSlice)
	return motherSlice
}

func SortDate(dates [][]string) {
	dateFormat := "02-01-2006" // Date format (day-month-year)
	sort.Slice(dates, func(i, j int) bool {
		// Parse the dates in index 1
		dateI, _ := time.Parse(dateFormat, dates[i][1])
		dateJ, _ := time.Parse(dateFormat, dates[j][1])
		// Sort by date
		return dateI.Before(dateJ)
	})
}

func sortdatess(dates []string) {
	datesformat := "02-01-2006"
	sort.Slice(dates, func(i, j int) bool {
		// Parse the dates in index 1
		dateI, _ := time.Parse(datesformat, dates[i])
		dateJ, _ := time.Parse(datesformat, dates[j])
		// Sort by date
		return dateI.Before(dateJ)
	})
}
