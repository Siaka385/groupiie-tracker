package api

import (
    "encoding/json"
    "fmt"
    "groupie-tracker/models"
    "io"
    "net/http"

)

// FetchAPIData is a helper function that sends an HTTP GET request to a given
// URL and returns the response body as a byte slice.
func FetchAPIData(url string) ([]byte, error) {
    response, err := http.Get(url)
    if err != nil {
        return nil, fmt.Errorf("failed to fetch data from API: %v", err)
    }
    defer response.Body.Close()

    // Read the response body
    body, err := io.ReadAll(response.Body)
    if err != nil {
        return nil, fmt.Errorf("failed to read response body: %v", err)
    }

    return body, nil
}

// FetchArtists fetches and unmarshals the artist data from the API.
func FetchArtists() ([]models.Artist, error) {
    url := "https://groupietrackers.herokuapp.com/api/artists"
    body, err := FetchAPIData(url)
    if err != nil {
        return nil, err
    }

    var artists []models.Artist
    err = json.Unmarshal(body, &artists)
    if err != nil {
        return nil, fmt.Errorf("failed to unmarshal artist data: %v", err)
    }

    return artists, nil
}

// FetchLocations fetches and unmarshals the location data from the API.
// Assuming the structure is a map of location objects
func FetchLocations() ([]models.Location, error) {
    url := "https://groupietrackers.herokuapp.com/api/locations"
    body, err := FetchAPIData(url)
    if err != nil {
        return nil, err
    }

    var response models.LocationResponse
    err = json.Unmarshal(body, &response)
    if err != nil {
        return nil, fmt.Errorf("failed to unmarshal location data: %v", err)
    }

    return response.Index, nil
}

// FetchDates fetches and unmarshals the date data from the API.
func FetchDates() ([]models.Date, error) {
    url := "https://groupietrackers.herokuapp.com/api/dates"
    body, err := FetchAPIData(url)
    if err != nil {
        return nil, err
    }

    // Define a struct to match the JSON response structure
    type DatesResponse struct {
        Index []models.Date `json:"index"`
    }

    var response DatesResponse
    err = json.Unmarshal(body, &response)
    if err != nil {
        return nil, fmt.Errorf("failed to unmarshal date data: %v", err)
    }

    return response.Index, nil
}
// FetchRelations fetches and unmarshals the relation data from the API.

func FetchRelations() ([]models.Relation, error) {
    url := "https://groupietrackers.herokuapp.com/api/relation"
    body, err := FetchAPIData(url)
    if err != nil {
        return nil, err
    }

    var response models.RelationsResponse
    err = json.Unmarshal(body, &response)
    if err != nil {
        return nil, fmt.Errorf("failed to unmarshal relation data: %v", err)
    }

    return response.Index, nil
}


