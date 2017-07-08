package magicalduels

import "testing"

var w1 = NewWizard("JoJo", 100)
var w2 = NewWizard("Anselm", 100)
var d = Library{}
var g = NewGame(w1, w2, d)

func TestNewGame(t *testing.T) {

	AssertEqual(t, g.GetCurrentTurn().Name, w1.Name)
	AssertEqual(t, g.GetNextTurn().Name, w2.Name)
	AssertEqual(t, g.TurnStarted, false)
}

func TestPoolDrawOnStartTurn(*testing.T) {
	// Wizard can choose 2 elements to place to his pool

}

func TestStartturn(t *testing.T) {
	g.StartTurn(true) // Ok!
	AssertEqual(t, g.TurnStarted, true)
}

func TestEndTurn(t *testing.T) {
}
