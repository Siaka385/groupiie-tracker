package handlers

import (
	"html/template"
	"net/http"
)

// Handle function serves the parsed artist data as JSON
func Aboutus(w http.ResponseWriter, r *http.Request) {
	isfilepresent, _ := Checkfile("./", "about.html")

	if !isfilepresent {
		http.Redirect(w, r, "/404", http.StatusFound)
		return

	}

	tmp, _ := template.ParseFiles("about.html")

	tmp.Execute(w, nil)
}
