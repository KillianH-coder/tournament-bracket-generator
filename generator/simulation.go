package generator

import (
	"errors"
	"tournament/models"
)

var simulation = models.Simulation

// Current state of the matches
func LoadMatch() {
	match := simulation.SimulationRounds[simulation.CurrentRound][simulation.CurrentMatch]
	simulation.Player1, simulation.Player2 = match[0], match[1]
}

// Make a clean copy of result in case we restart
func CopyRounds(rounds [][][]string) [][][]string {
	copy := make([][][]string, len(rounds))
	for i, round := range rounds {
		copy[i] = make([][]string, len(round))
		for j, match := range round {
			copy[i][j] = append([]string(nil), match...)
		}
	}
	return copy
}

// Initialize everything for the simulation
func StartSimulation(rounds [][][]string) {
	simulation.SimulationRounds = CopyRounds(rounds)
	simulation.Started = true
	simulation.Finished = false
	simulation.Winner = ""
	simulation.CurrentRound, simulation.CurrentMatch = 0, 0
	LoadMatch()
}

// Reset everything when we stop
func StopSimulation(rounds [][][]string) {
	simulation.Started = false
	simulation.Finished = false
	simulation.Winner = ""
	simulation.CurrentRound, simulation.CurrentMatch = 0, 0
	simulation.Player1, simulation.Player2 = "", ""
	simulation.SimulationRounds = CopyRounds(rounds)
}

// Where we choose the winner (with the buttons)
// Add the winner and advance in the tournament
func DecideWinner(winner string) error {
	s := simulation

	if !s.Started || s.Finished {
		return errors.New("simulation not running")
	}
	if winner != s.Player1 && winner != s.Player2 {
		return errors.New("invalid winner")
	}

	r, m := s.CurrentRound, s.CurrentMatch

	// Choose the winner when we are at the end
	if r == len(s.SimulationRounds)-1 {
		s.Finished = true
		s.Winner = winner
		return nil
	}

	// Put the winner in the right slot of the next round
	s.SimulationRounds[r+1][m/2][m%2] = winner

	// Advance in the tournament
	if m == len(s.SimulationRounds[r])-1 {
		s.CurrentRound++
		s.CurrentMatch = 0
	} else {
		s.CurrentMatch++
	}

	LoadMatch()
	return nil
}
