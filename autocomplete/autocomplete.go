package autocomplete

import (
	"fmt"
	"strconv"

	"groupie-tracker/api"
	"groupie-tracker/models"
)

// Suggestion struct holds suggestion content for artists, members, creation dates, etc.
type Suggestion struct {
	Contents string `json:"contents"`
	Name     string `json:"name"`
	Id       []int  `json:"id"`
}

var (
	SuggestionItems []Suggestion    // Slice to store all suggestion items
	Artists         []models.Artist // Slice of Artist structs to store fetched artists
	err             error
	FetchingError   bool
)

// GenerateSuggestions fetches artist data and generates suggestions for autocomplete.
func GenerateSuggestions() {
	// Fetch artist data from the API.
	Artists, err = api.FetchArtists()
	if err != nil {
		fmt.Println("Error fetching artists:", err)
		FetchingError = true
		return
	}

	// Generate artist/band suggestions.
	var sugestArist Suggestion
	for i := 0; i < len(Artists); i++ {
		sugestArist.Contents = "artist/band"
		sugestArist.Name = Artists[i].Name
		sugestArist.Id = append(sugestArist.Id, Artists[i].ID)
		SuggestionItems = append(SuggestionItems, sugestArist)
		// Clear the ID slice for the next artist.
		sugestArist.Id = []int{}
	}

	// Generate suggestions for band members.
	var sugestBandMember Suggestion
	for i := 0; i < len(Artists); i++ {
		for k := 0; k < len(Artists[i].Members); k++ {
			// Ensure that artist name is not duplicated in members list.
			if Artists[i].Name != Artists[i].Members[k] {
				sugestBandMember.Contents = "member"
				sugestBandMember.Name = Artists[i].Members[k]
				sugestBandMember.Id = append(sugestBandMember.Id, Artists[i].ID)
				SuggestionItems = append(SuggestionItems, sugestBandMember)
				// Clear the ID slice for the next member.
				sugestBandMember.Id = []int{}
			}
		}
	}

	// Generate creation date suggestions.
	var CreationDates Suggestion
	var creates []string // To store unique creation dates.
	for i := 0; i < len(Artists); i++ {
		if check(strconv.Itoa(Artists[i].CreationDate), creates) {
			CreationDates.Contents = "creation-date"
			CreationDates.Name = strconv.Itoa(Artists[i].CreationDate)
			CreationDates.Id = append(CreationDates.Id, Artists[i].ID)
			SuggestionItems = append(SuggestionItems, CreationDates)
			// Clear the ID slice for the next creation date.
			CreationDates.Id = []int{}
			creates = append(creates, CreationDates.Name)
		} else {
			// Append the artist ID to an existing creation date suggestion.
			for k := len(SuggestionItems) - 2; k >= 0; k-- {
				if SuggestionItems[k].Contents == "member" {
					break
				}
				if SuggestionItems[k].Contents == "creation-date" {
					if SuggestionItems[k].Name == strconv.Itoa(Artists[i].CreationDate) {
						SuggestionItems[k].Id = append(SuggestionItems[k].Id, Artists[i].ID)
					}
				}
			}
		}
	}

	// Generate first album suggestions.
	creates = []string{} // Reset the creates array for first album suggestions.
	var FirstAlbum Suggestion
	for i := 0; i < len(Artists); i++ {
		if check(Artists[i].FirstAlbum, creates) {
			FirstAlbum.Contents = "First-Album"
			FirstAlbum.Name = Artists[i].FirstAlbum
			FirstAlbum.Id = append(FirstAlbum.Id, Artists[i].ID)
			SuggestionItems = append(SuggestionItems, FirstAlbum)
			// Clear the ID slice for the next first album.
			FirstAlbum.Id = []int{}
			creates = append(creates, FirstAlbum.Name)
		} else {
			// Append the artist ID to an existing first album suggestion.
			for k := len(SuggestionItems) - 2; k >= 0; k-- {
				if SuggestionItems[k].Contents == "creation-date" {
					break
				}
				if SuggestionItems[k].Contents == "First-Album" {
					if SuggestionItems[k].Name == Artists[i].FirstAlbum {
						SuggestionItems[k].Id = append(SuggestionItems[k].Id, Artists[i].ID)
					}
				}
			}
		}
	}

	// Fetch locations from the API.
	creates = []string{} // Reset the creates array for location suggestions.
	locations, errs := api.FetchLocations()
	if errs != nil {
		fmt.Println("Error fetching locations:", err)
		return
	}

	// Generate location suggestions.
	var locationSuggestion Suggestion
	for i := 0; i < len(locations); i++ {
		for k := 0; k < len(locations[i].Locations); k++ {
			if check(locations[i].Locations[k], creates) {
				locationSuggestion.Contents = "location"
				locationSuggestion.Name = locations[i].Locations[k]
				locationSuggestion.Id = append(locationSuggestion.Id, locations[i].ID)
				SuggestionItems = append(SuggestionItems, locationSuggestion)
				// Clear the ID slice for the next location.
				locationSuggestion.Id = []int{}
				creates = append(creates, locationSuggestion.Name)
			} else {
				// Append the location ID to an existing location suggestion.
				for v := len(SuggestionItems) - 2; v >= 0; v-- {
					if SuggestionItems[v].Contents == "First-Album" {
						break
					}
					if SuggestionItems[v].Contents == "location" {
						if SuggestionItems[v].Name == locations[i].Locations[k] {
							SuggestionItems[v].Id = append(SuggestionItems[v].Id, locations[i].ID)
						}
					}
				}
			}
		}
	}

	fmt.Println("Server is well up")
}

// check ensures the uniqueness of the provided string in the array.
func check(creation string, arr []string) bool {
	if len(arr) == 0 {
		return true
	}
	for i := 0; i < len(arr); i++ {
		if creation == arr[i] {
			return false
		}
	}
	return true
}
