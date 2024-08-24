package myfunc

import (
	"net/http"
	"os"
)

func StaticServer(w http.ResponseWriter, r *http.Request) {
	filePath := "." + r.URL.Path
	info, err := os.Stat(filePath)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	if info.IsDir() {
		http.Error(w, "Acces Forbiden", http.StatusForbidden)
		return
	}
	http.ServeFile(w, r, filePath)
}
