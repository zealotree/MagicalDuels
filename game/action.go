package magicalduels

type Phase string

var TurnStart Phase = "Triggered upon starting a new turn"
var CastingSpell Phase = "Triggered upon Casting a spell"

type ActionType struct {
	Name        string
	Description string
}

type Action struct {
	ActionType
	Name          string
	Description   string
	Value         int       // Any Non Damage Value (i.e. Heal, Shield, Draw)
	Chance        int       // The chance of this Action occuring
	Damage        int       // Used damage dealing abilities
	Lifespan      int       // Used by abilities that relies on turns
	BlockElement  []Element // Use by ElementalShields
	GiveElement   []Element // Used bu BonusDamageElementalDrops
	IsStatus      bool
	StatusType    StatusType
	StatusName    string // Used to override debuff Name
	DamagePerTurn int    // How much damage it deals per turn
	DamagePerProc int    // How much damage it deals when triggered
	Uses          int    // How much time this Ability can be used
	Trigger       Phase
}

/////////// BASE ACTION Type DEFINITIONS ///////////////

var None = ActionType{
	Name:        "None",
	Description: "This Action does nothing! Purely for testing",
}

var Heal = ActionType{
	Name:        "Heal",
	Description: "Specify a value",
}

var Shield = ActionType{
	Name:        "Shield",
	Description: "Specify a value",
}

var ElementalShield = ActionType{
	Name:        "Elemental Shield",
	Description: "Specify a value and BlockElement",
}

var TrueShield = ActionType{
	Name:        "True Shield",
	Description: "Specify a value",
}

var Draw = ActionType{
	Name:        "Draw",
	Description: "Draw a spell from the deck",
}

var BonusDamage = ActionType{
	Name:        "Bonus Damage",
	Description: "Specify a damage and chance",
}

var EffectDamage = ActionType{
	Name:        "Deals Effect Damage",
	Description: "Specify a damage and chance",
}

var DamagePerTurn = ActionType{
	Name:        "Damage Per Turn",
	Description: "Specify DamagePerTurn, Lifespan, and chance",
}
