package magicalduels

import "fmt"

func (w *Wizard) CanCast(s Spell) bool {

	// Check Wizard's Pool if spell can be casted
	//
	cost := s.Cost
	elements := s.ElementDamage

	if len(cost) == 0 {
		return true
	}

	// Counts elements in Pool
	costCount := GetElementCountFromCards(w.Pool, elements)

	var costs []Element
	for _, e := range cost {
		costs = append(costs, e.Element)
	}

	spellCost := GetElementCount(costs)
	minCount := len(spellCost)
	costCounter := 0
	for element, _ := range spellCost {
		if _, ok := costCount[element]; ok {
			costCounter++
		}
	}

	if costCounter >= minCount {
		return true
	}
	return false

}

func DoSpellChance(s Spell) bool {
	if s.Action.Chance != 0 { // Cards with 0 chance dont exist
		if RNG(s.Action.Chance) {
			return true
		} else {
			return false
		}

	} else {
		return true
	}
	return false
}

func ActionTypeSwitch(s Spell, target *Wizard, deck *Library, applyEffectDamage bool) {
	switch s.ActionType {

	case Heal:
		target.RestoreHealth(s.Value)

	case Draw:
		target.Draw(deck, s.Action.Value)

	case Shield:
		target.AddShield(s)

	case ElementalShield:
		target.AddShield(s)

	case TrueShield:
		target.AddShield(s)

	case BonusDamage:
		target.TakeBonusDamage(s)

	case EffectDamage:
		if applyEffectDamage {
			target.TakeEffectDamage(s)
		}

	case DamagePerTurn:
		if applyEffectDamage {
			target.TakeEffectDamage(s)
		}

	}
}

func (w *Wizard) Cast(s Spell, target *Wizard, deck *Library) bool {

	// Check if this spell is in the hand
	hasCard, _ := w.HasSpell(s)

	if !hasCard {
		fmt.Println(w.Name, "does not own", s.Name)
		return false
	}

	targetName := "at " + target.Name
	if target.Name == w.Name {
		targetName = "to himself"
	}
	if s.Self {
		target = w
		targetName = ""
	}

	// Check cost of spell
	if w.CanCast(s) && hasCard {

		apply := false
		fmt.Println(w.Name, "casts", s.Name, targetName)
		if !DoSpellChance(s) {
			fmt.Println(s.Name, "failed!")
			w.RemoveCost(s.Cost)
			w.RemoveSpell(s)
			return false
		}
		// Trigger any cast effects
		ProcEffects(w, CastingSpell)
		success := true

		if s.Damage != 0 {
			target.TakeSpellDamage(s)
		}

		if s.Action.IsStatus {
			s.InflictStatus(target)
		}

		ActionTypeSwitch(s, target, deck, apply)

		if success {
			w.RemoveCost(s.Cost)
			w.RemoveSpell(s)
		}
		return true

	} else {
		fmt.Println(w.Name, "tried to cast ", s.Name, targetName, "But there's not enough element")
		return false
	}
}
