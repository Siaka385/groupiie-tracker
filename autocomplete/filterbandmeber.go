package autocomplete

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"text/template"

	"groupie-tracker/api"
	"groupie-tracker/handlers"
	"groupie-tracker/models"
)

type Bandinfo struct {
	Bandmember   string     `json:"bandmember"`
	ID           int        `json:"id"`           // Unique identifier for the artist.
	Name         string     `json:"name"`         // Name of the artist or band.
	Image        string     `json:"image"`        // URL of the artist or band's image.
	Members      []string   `json:"members"`      // List of band members.
	CreationDate int        `json:"creationDate"` // Year the band was formed.
	FirstAlbum   string     `json:"firstAlbum"`   // Date of the first album release.
	Relation     [][]string `json:"relation"`     // relation
	Locate       []string   `json:"locate"`       // locattion
	Datess       []string   `json:"datess"`       // dates
}

func MemberDisplay(w http.ResponseWriter, r *http.Request, path string) {
	members := strings.Split(path, "-")
	id, _ := strconv.Atoi(strings.TrimSpace(members[0]))

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
	artistInfo.Relation = handlers.Timeline(id)
	if artistInfo.Relation == nil {
		http.Redirect(w, r, "/500", http.StatusFound)
		return
	}
	artistInfo.Locate = handlers.Locations(id)
	if artistInfo.Locate == nil {
		http.Redirect(w, r, "/500", http.StatusFound)
		return
	}
	artistInfo.Datess = handlers.Dates(id)
	if artistInfo.Datess == nil {
		http.Redirect(w, r, "/500", http.StatusFound)
		return
	}

	BandMemberInformation := Bandinfo{
		Bandmember:   members[2],
		ID:           artistInfo.ID,
		Name:         artistInfo.Name,
		Image:        artistInfo.Image,
		Members:      artistInfo.Members,
		CreationDate: artistInfo.CreationDate,
		FirstAlbum:   artistInfo.FirstAlbum,
		Relation:     artistInfo.Relation,
		Locate:       artistInfo.Locate,
		Datess:       artistInfo.Datess,
	}

	// Parse and execute the template with the artist information
	tmpl, err := template.ParseFiles("bandmemberpage.html")
	if err != nil {
		http.Redirect(w, r, "/500", http.StatusFound)
		return
	}

	// Execute the template (no redirects after this point)
	if err := tmpl.Execute(w, BandMemberInformation); err != nil {
		http.Redirect(w, r, "/500", http.StatusFound)
	}
}
