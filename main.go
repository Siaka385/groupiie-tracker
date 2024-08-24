package main

import (
	"encoding/json"
	"fmt"
	"groupiie-tracker/myfunc"
	"html/template"
	"io"
	"log"
	"net/http"
)

// Artist struct represents the data model for each artist
type Artist struct {
	Id           int      `json:"id"`
	Name         string   `json:"NAME"`
	Image        string   `json:"image"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationdate"`
	FirstAlbum   string   `json:"firstalbum"`
}

// Myartists struct contains a slice of Artist structs
type Myartists struct {
	Mydata []Artist
}

// Handle function serves the parsed artist data as JSON
func Handle(w http.ResponseWriter, r *http.Request) {
	response, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		http.Error(w, "Failed to fetch data", http.StatusInternalServerError)
		return
	}
	defer response.Body.Close()

	data, _ := io.ReadAll(response.Body)
	var myartist []Artist

	jsonErr := json.Unmarshal(data, &myartist)
	if jsonErr != nil {
		http.Error(w, "Failed to parse JSON", http.StatusInternalServerError)
		return
	}

	// mytrial := Myartists{
	// 	Mydata: myartist,
	// }

	tmp, _ := template.ParseFiles("index.html")

	tmp.Execute(w, nil)
}

func main() {

	http.HandleFunc("/css/", myfunc.StaticServer)
	http.HandleFunc("/fonts/", myfunc.StaticServer)
	http.HandleFunc("/images/", myfunc.StaticServer)
	http.HandleFunc("/js/", myfunc.StaticServer)

	// Set up the HTTP server and route
	http.HandleFunc("/", Handle)

	// Start the server on port 8080
	fmt.Println("Server is running on port 8080...")
	if err := http.ListenAndServe(":8089", nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
