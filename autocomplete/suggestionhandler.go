package autocomplete

import (
	"net/http"
	"strconv"
	"strings"
	"time"
)

func HandleAutocompleteSelection(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Query().Get("search")

	if CheckIfCreationdate(path) {
		CreationDate(path, w, r)
	} else if CheckFirstAlbum(path) {
		FirstAlbums(path, w, r)
	} else if checkifbandmember(path) {
		MemberDisplay(w, r, path)
	} else {
		Locations(path, w, r)
	}
}

func CheckIfCreationdate(date string) bool {
	_, err := strconv.Atoi(strings.TrimSpace(date))

	return err == nil
}

func checkifbandmember(member string) bool {
	return strings.Contains(member, "-bandmember-")
}

func CheckFirstAlbum(date string) bool {
	dateformat := "02-01-2006"
	_, err := time.Parse(dateformat, date)

	return err == nil
}
