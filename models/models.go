package models

// Artist struct represents the information about an artist or band.
type Artist struct {
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

// Location represents a single entry in the index array from the API response.
type Location struct {
	ID        int      `json:"id"`
	Locations []string `json:"locations"`
	Dates     string   `json:"dates"`
}

// LocationResponse represents the top-level structure of the API response.
type LocationResponse struct {
	Index []Location `json:"index"`
}

// Date struct represents the concert dates for an artist.
type Date struct {
	ID    int      `json:"id"`    // Unique identifier for the date data.
	Dates []string `json:"dates"` // List of concert dates.
}

// Relation struct links artists to their concert dates and locations.
// Relation represents each entry in the "index" array of the relations API response.
type Relation struct {
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

// RelationsResponse represents the top-level structure containing the array of relations.
type RelationsResponse struct {
	Index []Relation `json:"index"`
}
