package generator

import (
	"net/http"
	"tournament/models"
)

var simulation = models.Simulation

func StartSimulation(w http.ResponseWriter, r *http.Request) {
	if !simulation.SimulationStarted {
		datas.Error = "simulation already started"
		http.Redirect(w, r, "/simulation", http.StatusSeeOther)
		return
	}

	simulation.SimulationStarted = true
}

func StopSimulation(w http.ResponseWriter, r *http.Request) {
	simulation.SimulationStarted = false

	http.Redirect(w, r, "/simulation", http.StatusSeeOther)
}

func ChooseWinner(w http.ResponseWriter, r *http.Request) {
	// stop at the current match
}

func DecideWinner(w http.ResponseWriter, r *http.Request) {
	winner := r.FormValue("winner")

	if !simulation.SimulationStarted {
		datas.Error = "Simulation not started"
		return
	}
}
