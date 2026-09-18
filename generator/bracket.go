package generator

import (
	"net/http"
)

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
