package autocomplete

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"groupie-tracker/api"
	"groupie-tracker/models"
)

type SearchResultPage struct {
	TitleMessage         string
	Artist               []models.Artist
	Member               []BandMember
	Locations            [][]string
	CreationDates        []string
	FirstALbum           []string
	DisplayArtist        string
	DisplayMembers       string
	DisplayLocations     string
	DisplayCreationDates string
	DispayFirstAlbum     string
}
type BandMember struct {
	Name     string
	Bandname string
	Id       int
}

func SearchPageHandler(w http.ResponseWriter, r *http.Request) {
	time.Sleep((time.Second * 1) / 2)

	seachitem := r.FormValue("search")

	for i := 0; i < len(SuggestionItems); i++ {
		if strings.EqualFold(strings.TrimSpace(seachitem), strings.TrimSpace(SuggestionItems[i].Name)) {
			if SuggestionItems[i].Contents == "member" {

				http.Redirect(w, r, fmt.Sprintf("/serch?search=%v-bandmember-%v", SuggestionItems[i].Id[0], SuggestionItems[i].Name), http.StatusFound)
				return
			} else if SuggestionItems[i].Contents == "artist/band" {

				http.Redirect(w, r, fmt.Sprintf("/artist?id=%v", SuggestionItems[i].Id[0]), http.StatusFound)
				return
			} else {

				http.Redirect(w, r, fmt.Sprintf("/serch?search=%v", SuggestionItems[i].Name), http.StatusFound)
				return

			}
		}
	}

	var Storesuggestion []Suggestion
	jsonfile, err := os.Open("js/search.json")
	if err != nil {
		Error500(w, r)
		return
	}
	defer jsonfile.Close()

	content, err := io.ReadAll(jsonfile)
	if err != nil {
		Error500(w, r)
		return
	}

	err = json.Unmarshal(content, &Storesuggestion)
	if err != nil {
		Error500(w, r)
		return
	}

	if len(Storesuggestion) == 0 {
		http.Redirect(w, r, "/requestnotfound", http.StatusFound)
		return
	}

	myartist, err := api.FetchArtists()
	if err != nil {
		Error500(w, r)
		return
	}

	var searchPage SearchResultPage
	searchPage.TitleMessage = "Suggestions that match" + ` "` + seachitem + `"` + ``
	for i := 0; i < len(Storesuggestion); i++ {
		if Storesuggestion[i].Contents == "artist/band" {
			for k := 0; k < len(myartist); k++ {
				if Storesuggestion[i].Id[0] == myartist[k].ID {
					searchPage.Artist = append(searchPage.Artist, myartist[k])
					break
				}
			}
		} else if Storesuggestion[i].Contents == "member" {
			for k := 0; k < len(myartist); k++ {
				if Storesuggestion[i].Id[0] == myartist[k].ID {
					bandmember := BandMember{
						Name:     Storesuggestion[i].Name,
						Bandname: myartist[k].Name,
						Id:       myartist[k].ID,
					}
					searchPage.Member = append(searchPage.Member, bandmember)
					break
				}
			}
		} else if Storesuggestion[i].Contents == "First-Album" {
			searchPage.FirstALbum = append(searchPage.FirstALbum, Storesuggestion[i].Name)
		} else if Storesuggestion[i].Contents == "location" {
			searchPage.Locations = append(searchPage.Locations, strings.Split(Storesuggestion[i].Name, "-"))
		} else if SuggestionItems[i].Contents == "creation-date" {
			searchPage.CreationDates = append(searchPage.CreationDates, Storesuggestion[i].Name)
		}
	}

	if len(searchPage.Artist) == 0 {
		searchPage.DisplayArtist = "hider"
	} else {
		searchPage.DisplayArtist = "show"
	}

	if len(searchPage.Member) == 0 {
		searchPage.DisplayMembers = "hider"
	} else {
		searchPage.DisplayMembers = "show"
	}

	if len(searchPage.Locations) == 0 {
		searchPage.DisplayLocations = "hider"
	} else {
		searchPage.DisplayLocations = "show"
	}

	if len(searchPage.CreationDates) == 0 {
		searchPage.DisplayCreationDates = "hider"
	} else {
		searchPage.DisplayCreationDates = "show"
	}

	if len(searchPage.FirstALbum) == 0 {
		searchPage.DispayFirstAlbum = "hider"
	} else {
		searchPage.DispayFirstAlbum = "show"
	}

	tmp, err := template.ParseFiles("SearchResultPage.html")
	if err != nil {
		fmt.Println(err)
		Error500(w, r)
		return
	}

	err = tmp.Execute(w, searchPage)
	if err != nil {
		Error500(w, r)
		return
	}
}
