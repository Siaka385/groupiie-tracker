package handlers

import (
	"fmt"
	"net/http"
	"os"
	"text/template"
)

func ErrorRenderPage(w http.ResponseWriter, r *http.Request, ErrorStatusCode int, Errorpage string) {
	tmp, err := template.ParseFiles(Errorpage)
	if err != nil {
		http.Redirect(w, r, "/500", http.StatusFound)
		return
	}
	w.WriteHeader(ErrorStatusCode)
	err = tmp.Execute(w, nil)

	if err != nil {
		http.Redirect(w, r, "/500", http.StatusFound)
		return
	}
}

func InternalServerError(w http.ResponseWriter, r *http.Request) {
	tmp, err := template.ParseFiles("Errortemplate/error500.html")
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusInternalServerError)
	if err := tmp.Execute(w, nil); err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
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
	resp, err := http.Get("http://clients3.google.com/generate_204")
	if err != nil {
		return false
	}
	defer resp.Body.Close() // Ensure the response body is closed

	// Check if the response status code is 204
	return resp.StatusCode == http.StatusNoContent
}
