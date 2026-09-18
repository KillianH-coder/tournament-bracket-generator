package models

// Struct for datas
type datas struct {
	Error           string
	Names           []string
	Rounds          [][][]string
	GenerateBracket bool
	StartSimulation bool
}

var Datas = &datas{}
