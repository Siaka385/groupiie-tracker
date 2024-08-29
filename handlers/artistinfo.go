package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"groupie-tracker/api"
	"groupie-tracker/models"
)

func Artinfo(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("id")
	artists, err := api.FetchArtists()
	if err != nil {
		fmt.Println("Error fetching artists:", err)
		return
	}

	id, _ := strconv.Atoi(name)

	var artistinfomation models.Artist

	for _, musicartist := range artists {
		if musicartist.ID == id {
			artistinfomation = musicartist
			break 
		}
	}

	tmp, _ := template.ParseFiles("listing-page.html")
	fmt.Println(artistinfomation.Members[0])
	tmp.Execute(w, artistinfomation)
}
