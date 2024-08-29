package handlers

import (
	"fmt"
	"net/http"
	"strings"

	"groupie-tracker/api"
	"groupie-tracker/models"
)

func SearchBar(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	Artistname := r.FormValue("search")

	artists, err := api.FetchArtists()
	if err != nil {
		fmt.Println("Error fetching artists:", err)
		return
	}

	var artistinfomation models.Artist

	for _, musicartist := range artists {
		if tolowercase(musicartist.Name) == tolowercase(Artistname) {
			artistinfomation = musicartist
			break
		}
	}

	http.Redirect(w, r, fmt.Sprintf("/artist?id=%v", artistinfomation.ID), http.StatusFound)
}

func tolowercase(m string) string {
	k := strings.ToLower(m)

	return k
}
