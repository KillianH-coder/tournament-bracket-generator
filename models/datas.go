package models

// Struct for datas
type datas struct {
	Error           string
	Names           []string
	Rounds          [][][]string
	GenerateBracket bool
}

type simulation struct {
	SimulationRounds  [][][]string
	Player1           string
	Player2           string
	SimulationStarted bool
}

var Datas = &datas{}
var Simulation = &simulation{}
