package handlers

import (
	"fmt"
	"net/http"
	"os"
	"text/template"
)

func InternalServerError(w http.ResponseWriter, r *http.Request) {
	isfilepresent, _ := Checkfile("./Errortemplate/", "error500.html")
	if !isfilepresent {
		http.Redirect(w, r, "/404", http.StatusFound)
		return
	}

	tmp, err := template.ParseFiles("Errortemplate/error500.html")
	if err != nil {
		http.Redirect(w, r, "/500", http.StatusFound)
		return
	}
	w.WriteHeader(http.StatusInternalServerError)
	if err := tmp.Execute(w, nil); err != nil {
		http.Redirect(w, r, "/500", http.StatusFound)
		return
	}
}

func Error404(w http.ResponseWriter, r *http.Request) {
	isfilepresent, _ := Checkfile("./Errortemplate/", "error500.html")
	if !isfilepresent {
		http.Error(w, "page not found", http.StatusNotFound)
		return
	}
	tmpl, err := template.ParseFiles("Errortemplate/error.html")
	if err != nil {
		http.Redirect(w, r, "/500", http.StatusFound)
		return
	}

	w.WriteHeader(http.StatusNotFound)
	err = tmpl.Execute(w, nil)

	if err != nil {
		http.Redirect(w, r, "/500", http.StatusFound)
		return
	}
}

func Wrongmethod(w http.ResponseWriter, r *http.Request) {
	isfilepresent, _ := Checkfile("./Errortemplate/", "wrongmethodused.html")

	if !isfilepresent {
		http.Redirect(w, r, "/404", http.StatusFound)
		return
	}

	tmp, err := template.ParseFiles("Errortemplate/wrongmethodused.html")
	if err != nil {
		http.Redirect(w, r, "/500", http.StatusFound)
		return
	}
	w.WriteHeader(http.StatusMethodNotAllowed)

	if err := tmp.Execute(w, nil); err != nil {
		http.Redirect(w, r, "/500", http.StatusFound)
		return
	}
}

func Nointernetconnection(w http.ResponseWriter, r *http.Request) {
	isfilepresent, _ := Checkfile("./Errortemplate/", "internetconnection.html")

	if !isfilepresent {
		http.Error(w, "File is missing", http.StatusNotFound)
		return
	}
	tmp, err := template.ParseFiles("Errortemplate/internetconnection.html")
	if err != nil {
		http.Error(w, "server errr", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusRequestTimeout)

	if err := tmp.Execute(w, nil); err != nil {
		http.Error(w, "server errr", http.StatusInternalServerError)
		return
	}
}

func ArtistNotFound(w http.ResponseWriter, r *http.Request) {
	isfilepresent, _ := Checkfile("./Errortemplate/", "Noaristfound.html")

	if !isfilepresent {
		http.Error(w, "File is missing", http.StatusNotFound)
		return
	}
	tmp, err := template.ParseFiles("Errortemplate/Noaristfound.html")
	if err != nil {
		http.Error(w, "server errr", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusBadRequest)

	if err := tmp.Execute(w, nil); err != nil {
		http.Error(w, "server errr", http.StatusInternalServerError)
		return
	}
}

func Checkfile(directort, filename string) (bool, error) {
	// Open the directory
	dir, err := os.Open(directort)
	if err != nil {
		return false, err
	}
	defer dir.Close()

	files, err := dir.Readdir(-1)
	if err != nil {
		fmt.Println(err)
		return false, err
	}

	// Check if the file exists in the directory
	for _, file := range files {
		if file.Name() == filename {
			return true, nil
		}
	}

	return false, nil
}

func CheckInternetConnectivity() (ok bool) {
	resp, error := http.Get("http://clients3.google.com/generate_204")
	if error != nil {
		return false
	}
	defer resp.Body.Close() // Ensure the response body is closed

	// Check if the response status code is 204
	return resp.StatusCode == http.StatusNoContent
}
