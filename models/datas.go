package models

// Struct for datas
type datas struct {
	Error           string
	Names           []string
	Rounds          [][][]string
	GenerateBracket bool
	*simulation
}

type simulation struct {
	SimulationRounds [][][]string
	Player1          string
	Player2          string
	Winner           string
	CurrentRound     int
	CurrentMatch     int
	Started          bool
	Finished         bool
}

var Simulation = &simulation{}
var Datas = &datas{simulation: Simulation}

// Same here but understandable
// var Datas = &datas{}
// var Simulation = &simulation{}

// Datas.simulation = Simulation
