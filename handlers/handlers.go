package handlers

import (
	"tournament/generator"
	"tournament/models"

	"html/template"
	"net/http"
)

// Parse all templates
var tmpl = template.Must(template.ParseGlob("./templates/*.html"))

// structs
var datas = models.Datas

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	// 404
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	// buttons
	if r.Method == http.MethodPost {
		action := r.FormValue("action")

		switch action {
		case "add":
			generator.AddNameHandler(w, r)
		case "clear":
			generator.ClearNamesHandler(w, r)
		case "generate":
			generator.CreateBracket(w, r)
		case "delete":
			generator.DeleteBracket(w, r)
		}
		return
	}

	tmpl.ExecuteTemplate(w, "index.html", datas)
}

func SimulationHandler(w http.ResponseWriter, r *http.Request) {
	// GET for the page
	if r.Method != http.MethodPost {
		if !datas.Started {
			datas.SimulationRounds = generator.CopyRounds(datas.Rounds)
		}
		tmpl.ExecuteTemplate(w, "simulation.html", datas)
		return
	}

	// buttons
	switch r.FormValue("simulation") {
	case "start":
		if datas.Started {
			datas.Error = "simulation already started"
		} else if !datas.GenerateBracket {
			datas.Error = "generate bracket first"
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		} else {
			generator.StartSimulation(datas.Rounds)
		}
	// If a winner exist
	case "winner":
		if err := generator.DecideWinner(r.FormValue("winner")); err != nil {
			datas.Error = err.Error()
		}
	case "stop":
		generator.StopSimulation(datas.Rounds)
	case "home":
		generator.StopSimulation(datas.Rounds)
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	http.Redirect(w, r, "/simulation", http.StatusSeeOther)
}
