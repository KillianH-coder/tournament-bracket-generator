package handlers

import (
	"html/template"
	"net/http"
	"slices"
)

// Parse all templates
var tmpl = template.Must(template.ParseGlob("./templates/*.html"))

// Struct for datas
type Datas struct {
	Names           []string
	Error           string
	GenerateBracket bool
	Rounds          [][][]string
}

var datas = &Datas{}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		action := r.FormValue("action")

		switch action {
		case "add":
			AddNameHandler(w, r)
		case "clear":
			ClearNamesHandler(w, r)
		case "generate":
			CreateBracket(w, r)
		case "delete":
			DeleteBracket(w, r)
		}
		return
	}

	tmpl.Execute(w, datas)
}

// Add players to the list
func AddNameHandler(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")

	// If name is empty
	if name == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		datas.Error = "Name cannot be empty"
		return
	}

	// If name have more than 10 chars
	if len(name) > 10 {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		datas.Error = "Name cannot be longer than 10 characters"
		return
	}

	// If name already exist
	if slices.Contains(datas.Names, name) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		datas.Error = "Name already exists"
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

// Create the bracket
func CreateBracket(w http.ResponseWriter, r *http.Request) {
	if datas.GenerateBracket {
		datas.Error = "delete bracket first before creating a new one"
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	datas.Rounds = [][][]string{}
	players := len(datas.Names)
	names := datas.Names
	firstRound := [][]string{}

	if players > 32 {
		datas.Error = "too much players"
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	if players < 2 {
		datas.Error = "not enough players"
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	datas.GenerateBracket = true

	switch {
	case len(names) <= 4:
		for len(names) < 4 {
			names = append(names, "empty")
		}
	case len(names) <= 8:
		for len(names) < 8 {
			names = append(names, "empty")
		}
	case len(names) <= 16:
		for len(names) < 16 {
			names = append(names, "empty")
		}
	case len(names) <= 32:
		for len(names) < 32 {
			names = append(names, "empty")
		}
	}

	// Add bracket first round
	for index := 0; index < len(names); index += 2 {
		firstRound = append(firstRound, []string{names[index], names[index+1]})
	}
	datas.Rounds = append(datas.Rounds, firstRound)

	// Add next rounds (add datas.Rounds[1, etc...])
	for matches := len(datas.Rounds[0]) / 2; matches >= 1; matches /= 2 {
		round := [][]string{}

		for i := 0; i < matches; i++ {
			round = append(round, []string{"", ""})
		}

		datas.Rounds = append(datas.Rounds, round)
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// Delete the bracket
func DeleteBracket(w http.ResponseWriter, r *http.Request) {
	datas.GenerateBracket = false
	datas.Rounds = [][][]string{}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func Winner(w http.ResponseWriter, r *http.Request) {
	// winner := r.FormValue("winner")

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
