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
		http.Redirect(w, r, "/500", http.StatusFound)
		return
	}
	if info.IsDir() {
		tmp, err := template.ParseFiles("Errortemplate/accessforbidden.html")
		if err != nil {
			http.Redirect(w, r, "/500", http.StatusFound)
			return
		}
		w.WriteHeader(http.StatusForbidden)
		if err := tmp.Execute(w, nil); err != nil {
			http.Redirect(w, r, "/500", http.StatusFound)
			return
		}
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
	if r.Method != "GET" {
		http.Redirect(w, r, "/wrongmethod", http.StatusMethodNotAllowed)
		return
	}

	artists, err := api.FetchArtists()
	if err != nil {
		fmt.Println("Error fetching artists:", err)
		http.Redirect(w, r, "/500", http.StatusFound)
		return
	}

	artistInfo := Myartists{
		Mydata: artists,
	}

	if r.URL.Path == "/explore" {
		isfilepresent, _ := Checkfile("./", "explore.html")
		if !isfilepresent {
			http.Redirect(w, r, "/500", http.StatusFound)
			return
		}
		tmp, _ := template.ParseFiles("explore.html")

		tmp.Execute(w, artistInfo)
		return
	}
	isfilepresent, _ := Checkfile("./", "index.html")
	if !isfilepresent {
		http.Redirect(w, r, "/500", http.StatusFound)
		return
	}
	tmp, _ := template.ParseFiles("index.html")

	tmp.Execute(w, artistInfo)
}
