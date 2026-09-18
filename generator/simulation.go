package generator

import (
	"net/http"
)

func StartSimulation(w http.ResponseWriter, r *http.Request) {
	if !datas.StartSimulation {
		datas.Error = "simulation already started"
		http.Redirect(w, r, "/simulation", http.StatusSeeOther)
		return
	}

	datas.StartSimulation = true
}

func StopSimulation(w http.ResponseWriter, r *http.Request) {
	datas.StartSimulation = false

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func Winner(w http.ResponseWriter, r *http.Request) {
}
