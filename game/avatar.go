package magicalduels

import (
	//"errors"
	"fmt"
)

type ActiveShield struct {
	Type         ActionType // The shield type
	Name         string
	Value        int
	Uses         int
	BlockElement []Element
	Active       bool
}

type Wizard struct {
	Name         string
	Health       int
	MaxHealth    int
	Hand         []Card
	Pool         []Card
	Buff, Debuff Spell
	ActiveShield
	Alive bool
}

func NewWizard(Name string, MaxHealth int) Wizard {

	wiz1 := Wizard{
		Name:         Name,
		Health:       MaxHealth,
		MaxHealth:    MaxHealth,
		Hand:         []Card{},
		Pool:         []Card{},
		Buff:         Spell{},
		Debuff:       Spell{},
		ActiveShield: ActiveShield{},
		Alive:        true,
	}

	return wiz1
}

func (w *Wizard) Stats() {
	fmt.Println(fmt.Sprintf("%s %d HP \n FIRE: %d", w.Name, w.Health, len(w.Pool)))

}

func (w *Wizard) AddCardToHand(card Card, amount int) {
	for i := 0; i < amount; i++ {
		if card.IsSpell() {
			w.Hand = append(w.Hand, card)
		}
	}
}

func (w *Wizard) AddCardToPool(card Card, amount int) {
	for i := 0; i < amount; i++ {
		if card.IsElement() {
			w.Pool = append(w.Pool, card)
		}
	}
	fmt.Println(w.Name, "gains", "+", amount, card.GetName())
}

func (w *Wizard) DrawFromPool() {
}

func (w *Wizard) Draw(d *Library, amount int) {

	verb := "card"
	drawn := 0

	for i := 0; i < amount; i++ {
		if len(d.Cards) > 0 {
			rand := random(0, len(d.Cards))
			card := d.Cards[rand]
			w.Hand = append(w.Hand, card)
			drawn++
			d.Cards = append(d.Cards[:rand], d.Cards[rand+1:]...)
		} else {
			// No more cards in the deck. Take 5HP damage
			fmt.Println("There are no more spells available. Avatar begins to wither.")
			w.TakeDamage(5)

		}

	}
	if drawn != 0 && drawn != 1 {
		verb = "cards"
		fmt.Println(w.Name, "drew", drawn, verb)
	} else if drawn == 1 {
		fmt.Println(w.Name, "drew one card")

	}
	//	return nil

}

func (w *Wizard) GetShield() (bool, ActiveShield) {
	//returns true if a shield is found, and the status

	// check if actionshield is not empy
	if !w.ActiveShield.Active {
		return false, ActiveShield{}
	} else {
		return true, w.ActiveShield

	}

}

func (w *Wizard) UseActiveShield() {
	if w.ActiveShield.Uses != 0 {
		w.ActiveShield.Uses--
	}
	if w.ActiveShield.Uses == 0 {
		w.RemoveShield()

	}

}
func CanElementalShieldBlock(shield ActiveShield, s Spell) bool {

	emap := GetElementCount(s.ElementDamage)
	smap := GetElementCount(shield.BlockElement)

	canBlockMax := len(s.ElementDamage)
	counter := 0

	for element, _ := range emap {
		if _, ok := smap[element]; ok {
			counter++
		}
	}

	if counter == canBlockMax {
		return true
	}
	return false
}

func (w *Wizard) CheckForShields(s Spell) int {

	damage := s.Damage
	hasShield, shield := w.GetShield()

	if hasShield {
		if shield.Type == TrueShield {
			fmt.Println(shield.Name, "prevents any damage from happening")
			//w.ResolveStatus(status)
			return 0
		} else if shield.Type == ElementalShield {
			if CanElementalShieldBlock(shield, s) {
				damage = damage - shield.Value
				fmt.Println(shield.Name, "reduced incoming damage by", shield.Value)
				w.UseActiveShield()
			}
		} else if shield.Type == Shield {
			damage = damage - shield.Value
			fmt.Println(shield.Name, "reduced incoming damage by", shield.Value)
			w.UseActiveShield()
		}
		if damage <= 0 {
			damage = 0
		}
		//w.ResolveStatus(status)
	}
	return damage
}

func (w *Wizard) TakeEffectDamage(s Spell) {
	// effects
	fmt.Println(s.Name, "triggers")
	damage := 0
	if s.Action.DamagePerTurn != 0 {
		damage = s.Action.DamagePerTurn
	} else if s.DamagePerProc != 0 {
		damage = s.Action.DamagePerProc
	}
	w.TakeDamage(damage)
}

func (w *Wizard) TakeBonusDamage(s Spell) {

	if RNG(s.Action.Chance) {
		fmt.Println(s.Action.Name, "suceeded. Bonus Damage will apply.")
		w.TakeDamage(s.Action.Damage)
	}

}

func (w *Wizard) TakeSpellDamage(s Spell) {

	damage := w.CheckForShields(s)
	w.TakeDamage(damage)

}

func (w *Wizard) TakeDamage(damage int) {
	w.Health = w.Health - damage
	if damage <= 0 {
		fmt.Println(w.Name, "recieves", "no damage")
	} else {
		fmt.Println(w.Name, "recieves", damage, "damage")
	}
	if w.Health <= 0 {
		w.Alive = false
		fmt.Println(w.Name, "dies")

	}

}

func (w *Wizard) RestoreHealth(value int) {
	w.Health = w.Health + value
	if w.Health > w.MaxHealth {
		w.Health = w.MaxHealth
	}
	fmt.Println(w.Name, "heals for", value, "health")
}

func (w *Wizard) RemoveShield() {
	name := w.ActiveShield.Name
	w.ActiveShield = ActiveShield{}
	fmt.Println(name, "breaks")
}

func (w *Wizard) AddShield(s Spell) {
	name := s.Action.Name
	if name == "" {
		name = s.Name
	}
	a := ActiveShield{
		Name:         name,
		Type:         s.ActionType,
		Value:        s.Action.Value,
		Uses:         s.Action.Uses,
		BlockElement: s.Action.BlockElement,
		Active:       true,
	}
	w.ActiveShield = a
	fmt.Println(w.Name, "gains", w.ActiveShield.Name)
}

func (w *Wizard) RemoveElement(element Element) {

	for i, card := range w.Pool {
		if element == Any || card == element {
			//fmt.Println("Removing:", card)
			w.Pool = append(w.Pool[:i], w.Pool[i+1:]...)
			break
		}

	}
}

func (w *Wizard) RemoveCost(cost []Cost) {

	for _, c := range cost {
		for i := 0; i < c.Amount; i++ {
			w.RemoveElement(c.Element)
		}

	}

}

func (w *Wizard) RemoveSpell(spell Spell) {

	for i, card := range w.Hand {
		if card.GetName() == spell.Name {
			//fmt.Println("Removing:", card)
			w.Hand = append(w.Hand[:i], w.Hand[i+1:]...)
			break
		}

	}

}

func (w *Wizard) HasSpell(s Spell) (bool, int) {

	count := 0
	for _, spell := range w.Hand {
		if spell.GetName() == s.Name {
			count++
		}
	}

	if count > 0 {
		return true, count
	} else {
		return false, count
	}

}
