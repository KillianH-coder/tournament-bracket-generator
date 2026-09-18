package generator

import (
	"tournament/models"

	"net/http"
	"slices"
)

var datas = models.Datas

// Add players to the list
func AddNameHandler(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")

	// If name is empty
	if name == "" {
		datas.Error = "Name cannot be empty"
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	// If name have more than 10 chars
	if len(name) > 10 {
		datas.Error = "Name cannot be longer than 10 characters"
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	// If name already exist
	if slices.Contains(datas.Names, name) {
		datas.Error = "Name already exists"
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	datas.Names = append(datas.Names, name)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// Clear the list
func ClearNamesHandler(w http.ResponseWriter, r *http.Request) {
	datas.Names = []string{}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
