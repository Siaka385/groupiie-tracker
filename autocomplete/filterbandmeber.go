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
	if !HandleBandMemberBadRequest(path) {
		http.Redirect(w, r, "/badrequest", http.StatusFound)
		return
	}

	members := strings.Split(path, "-")

	id, _ := strconv.Atoi(strings.TrimSpace(members[0]))

	// Fetch the list of artists
	artists, err := api.FetchArtists()
	if err != nil {
		fmt.Println("Error fetching artists:", err)
		Error500(w, r)
		return
	}

	// Find the artist with the specified ID
	var artistInfo models.Artist
	ArtistFound := false
	for _, musicArtist := range artists {
		if musicArtist.ID == id {
			artistInfo = musicArtist
			ArtistFound = true
			break
		}
	}

	// If no artist found, redirect to 404
	if !ArtistFound {
		Error404(w, r)
		return
	}

	MemberPresent := false

	for i := 0; i < len(artistInfo.Members); i++ {
		if strings.EqualFold(strings.TrimSpace(artistInfo.Members[i]), strings.TrimSpace(members[2])) {
			MemberPresent = true
		}
	}
	if !MemberPresent {
		Error404(w, r)
		return
	}

	// Fetch and set the artist's timeline (relations)
	artistInfo.Relation = handlers.Timeline(id)
	if artistInfo.Relation == nil {
		Error500(w, r)
		return
	}
	artistInfo.Locate = handlers.Locations(id)
	if artistInfo.Locate == nil {
		Error500(w, r)
		return
	}
	artistInfo.Datess = handlers.Dates(id)
	if artistInfo.Datess == nil {
		Error500(w, r)
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
		Error500(w, r)
		return
	}

	// Execute the template (no redirects after this point)
	if err := tmpl.Execute(w, BandMemberInformation); err != nil {
		Error500(w, r)
	}
}

func HandleBandMemberBadRequest(m string) bool {
	myslice := strings.Split(m, "-")

	_, err := strconv.Atoi(myslice[0])
	if err != nil {
		return false
	}
	if myslice[1] != "bandmember" {
		return false
	}
	return true
}
