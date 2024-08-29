package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"os"

	"groupie-tracker/api"
	"groupie-tracker/models"
)

func StaticServer(w http.ResponseWriter, r *http.Request) {
	filePath := "." + r.URL.Path
	info, err := os.Stat(filePath)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	if info.IsDir() {
		http.Error(w, "Acces Forbiden", http.StatusForbidden)
		return
	}
	http.ServeFile(w, r, filePath)
}

// Myartists struct contains a slice of Artist structs
type Myartists struct {
	Mydata []models.Artist
}

// Handle function serves the parsed artist data as JSON
func Homepage(w http.ResponseWriter, r *http.Request) {
	artists, err := api.FetchArtists()
	if err != nil {
		fmt.Println("Error fetching artists:", err)
		return
	}

	artistInfo := Myartists{
		Mydata: artists,
	}

	tmp, _ := template.ParseFiles("index.html")

	tmp.Execute(w, artistInfo)
}
