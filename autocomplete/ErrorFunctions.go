package autocomplete

import "net/http"

func Error404(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/notfound", http.StatusFound)
}

func Error500(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/500", http.StatusFound)
}

func Error403(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/wrongmethod", http.StatusFound)
}
