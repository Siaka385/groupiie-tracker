package main

import (
	"fmt"
	"log"
	"net/http"

	"groupie-tracker/handlers"
	//"groupie-tracker/models"
	//"groupie-tracker/api"

)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/css/", handlers.StaticServer)
	mux.HandleFunc("/fonts/", handlers.StaticServer)
	mux.HandleFunc("/images/", handlers.StaticServer)
	mux.HandleFunc("/js/", handlers.StaticServer)

	// Set up the HTTP server and route
	mux.HandleFunc("/", handlers.Homepage)
	mux.HandleFunc("/artist", handlers.Artinfo)
	mux.HandleFunc("/search", handlers.SearchBar)
	// Start the server on port 8080
	fmt.Println("Server is running on port 8089...")
	if err := http.ListenAndServe(":8089", mux); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

// Fetch and print data from the API used for testing purposes

// 	locations, err := api.FetchLocations()
// 	if err != nil {
// 	    fmt.Println("Error fetching locations:", err)
// 	    return
// 	}
// 	fmt.Printf("Locations: %+v\n", locations) // Print locations to use the variable

// 	dates, err := api.FetchDates()
// 	if err != nil {
// 	    fmt.Println("Error fetching dates:", err)
// 	    return
// 	}
// 	fmt.Printf("Dates: %+v\n", dates) // Print dates to use the variable

// 	relations, err := api.FetchRelations()
// 	 if err != nil {
// 	   fmt.Println("Error fetching relations:", err)
// 	    return
// 	}
// 	fmt.Printf("Relations: %+v\n", relations) // Print relations to use the variable
// }
