package magicalduels

import (
	//"fmt"
	"math/rand"
	"time"
)

func random(min, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return rand.Intn(max-min) + min
}

func RNG(chance int) bool {
	num := random(1, 100)
	if num < chance {
		return true
	}
	return false
}

// Two types of Cards: Elements, Spell
// Spells can only be casted by discarding Element Card

type Card interface {
	GetName() string
	IsElement() bool
	IsSpell() bool
}

type Library struct {
	Cards []Card
}

func (d *Library) AddCard(card Card, amount int) {
	for i := 0; i < amount; i++ {
		if card.IsSpell() {
			d.Cards = append(d.Cards, card)
		}
	}
}

// Adds Spells
func (d *Library) Init() {
	d.AddCard(Fireball, 5)
	d.AddCard(IceDaggers, 5)
	d.AddCard(RockPunch, 5)

	d.AddCard(ShadowWager, 5)

}

func GetElementCount(elements []Element) map[Element]int {
	out := make(map[Element]int)
	for _, element := range elements {
		out[element]++
	}
	return out
}

func GetElementCountFromCards(cards []Card, elements []Element) map[Element]int {
	// returns a map of elements counts in a given list of cards
	// map[Element]int

	out := make(map[Element]int)

	for _, card := range cards {
		if card.IsElement() {
			out[card.(Element)]++
		}
	}

	return out

}
