package magicalduels

// Spells
// Name: Name of SPell
// Damage: This spell does damage
// Cost: The cost of the elemental
// Action: This spell has a special action.
//  - Heal
//  - Shield
//  - Buff/Debuff
//  - Bonus Damage
// Self: This spell can only be casted to self, if false; anyone can be targetted.

// Action

type Cost struct {
	Element
	Amount int
}

type Spell struct {
	Name          string
	Damage        int
	ElementDamage []Element
	Cost          []Cost
	Action
	Self bool // If self, this spell only targets the Caster
}

func (s Spell) GetName() string {
	return s.Name
}

func (s Spell) IsElement() bool {
	return false
}

func (s Spell) IsSpell() bool {
	return true
}
