package handlers

import (
	"fmt"
	"net/http"
	"strings"

	"groupie-tracker/api"
	"groupie-tracker/models"
)

func HandleManualSearch(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	if r.Method != "GET" {
		http.Redirect(w, r, "/wrongmethod", http.StatusMethodNotAllowed)
		return
	}

	Artistname := r.FormValue("search")

	artists, err := api.FetchArtists()
	if err != nil {
		fmt.Println("Error fetching artists:", err)
		http.Redirect(w, r, "/500", http.StatusFound)
		return
	}

	var artistinfomation models.Artist
	found := false
	for _, musicartist := range artists {
		if tolowercase(musicartist.Name) == tolowercase(Artistname) {
			artistinfomation = musicartist
			found = true
			break
		}
	}
	if !found {
		http.Redirect(w, r, "/badrequest", http.StatusFound)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/artist?id=%v", artistinfomation.ID), http.StatusFound)
}

func tolowercase(m string) string {
	k := strings.ToLower(m)

	return k
}
