package handlers

import (
	"tournament/generator"
	"tournament/models"

	"html/template"
	"net/http"
)

// Parse all templates
var tmpl = template.Must(template.ParseGlob("./templates/*.html"))

var datas = models.Datas
var simulation = models.Simulation

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

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
	if r.Method == http.MethodPost {
		if simulation.SimulationStarted {
			datas.Error = "simulation already started"
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}

		if !datas.GenerateBracket {
			datas.Error = "generate bracket first"
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}

		action := r.FormValue("simulation")
		switch action {
		case "start":
			generator.StartSimulation(w, r)
			generator.DecideWinner(w, r)
		case "stop":
			generator.StopSimulation(w, r)
		case "home":
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}
		return
	}

	tmpl.ExecuteTemplate(w, "simulation.html", datas)
}
