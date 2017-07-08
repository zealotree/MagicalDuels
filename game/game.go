package magicalduels

import "fmt"

type Player struct {
	Username string
	*Wizard
}

type GameObjects struct {
	TurnOrder []Wizard // 0 is Current, 1 is next.
	Library
	TurnStarted bool
	On          bool
}

type GameLoop interface {
	Init()
	GetCurrentTurn()
	GetNextTurn()
	GameOver()
	StartTurn()
	Do()
	Run()
}

func (g *GameObjects) Run(test bool) {
	for {
		if g.On {
			g.StartTurn(test)
			//g.EndTurn()
		} else {
			break
		}
	}
}

func (g *GameObjects) Do() {
}

func NewGame(w1 Wizard, w2 Wizard, d Library) GameObjects {
	game := GameObjects{
		[]Wizard{
			w1,
			w2,
		},
		d,
		false,
		true,
	}
	return game
}

func (g *GameObjects) Init() {
	g.Library.Init()
}

func (g *GameObjects) GameOver(w Wizard) {
	if !w.Alive {
		g.On = false
		fmt.Println(w.Name, "is dead.")
		fmt.Println("GameOVER")
	}

}

func (g *GameObjects) GetCurrentTurn() *Wizard {
	return &g.TurnOrder[0]
}

func (g *GameObjects) GetNextTurn() *Wizard {
	return &g.TurnOrder[1]
}

func (g *GameObjects) StartTurn(test bool) {

	w := g.GetCurrentTurn()
	if !w.Alive {
		g.GameOver(*w)
	} else {
		g.TurnStarted = true
		fmt.Println(w.Name, "begins the turn")
		ProcEffects(w, TurnStart)
		w.Draw(&g.Library, 1)
		if test {
			w.AddCardToPool(Void, 2)
			w.AddCardToPool(Fire, 2)
			w.AddCardToHand(SpellVirus, 1)
			w.AddCardToHand(Fireball, 1)
			w.Cast(SpellVirus, w, &g.Library)
			w.Cast(Fireball, w, &g.Library)

		} else {
			w.BigPoolRequest() // CLI: ask for the elements
		}
	}
}

func (g *GameObjects) EndTurn(test bool) {

	w := g.GetCurrentTurn()
	// Proc effects
	if !w.Alive {
		g.GameOver(*w)
	} else {
		g.TurnStarted = true
		fmt.Println(w.Name, "begins the turn")
		w.Draw(&g.Library, 1)
		if test {

		} else {
			w.BigPoolRequest() // CLI: ask for the elements
		}
	}
}
