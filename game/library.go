package magicalduels

// Spell library

//////////// SPELL CARDS DEFINITION //////////////////

// Base Damage Spell

var Fireball = Spell{
	Name:          "Fireball",
	ElementDamage: []Element{Fire},
	Damage:        12,
	Cost: []Cost{
		Cost{Fire, 1},
	},
}

var IceDaggers = Spell{
	Name:          "Ice Daggers",
	ElementDamage: []Element{Water},
	Damage:        10,
	Cost: []Cost{
		Cost{Water, 1},
	},
}

var RockPunch = Spell{
	Name:          "Rock Punch",
	ElementDamage: []Element{Earth},
	Damage:        10,
	Cost: []Cost{
		Cost{Earth, 1},
	},
}

////////// Multi Elemental Spells /////

var VoidStrikeAction = Action{
	ActionType: BonusDamage,
	Damage:     2,
	Chance:     100,
}

var VoidStrike = Spell{
	Name:          "Void Strike",
	ElementDamage: []Element{Void, Earth},
	Damage:        10,
	Cost: []Cost{
		Cost{Void, 1},
		Cost{Earth, 2},
	},
	Action: VoidStrikeAction,
}

var HeatVortex = Spell{
	Name:          "Heat Vortex",
	ElementDamage: []Element{Fire, Air},
	Damage:        16,
	Cost: []Cost{
		Cost{Air, 1},
		Cost{Fire, 1},
	},
}

var MoltenBarrage = Spell{
	Name:          "Molten Barrage",
	ElementDamage: []Element{Fire, Earth},
	Damage:        30,
	Cost: []Cost{
		Cost{Earth, 2},
		Cost{Fire, 2},
	},
}

//////////// Healing spells ////////

var SweetWaterAction = Action{
	ActionType: Heal,
	Value:      8,
}

var SweetWater = Spell{

	Name: "Sweetwater",
	Cost: []Cost{
		Cost{Water, 1},
	},
	Action: SweetWaterAction,
}

var SacredRootAction = Action{
	ActionType: Heal,
	Value:      20,
}

var SacredRoot = Spell{

	Name: "Sacred Root",
	Cost: []Cost{
		Cost{Earth, 3},
	},
	Action: SacredRootAction,
}

///////////// Shield Spells //////

var SmallShieldAction = Action{
	ActionType: Shield,
	Value:      6,
	Uses:       1,
}

var SmallShield = Spell{
	Name: "Small Shield",
	Cost: []Cost{
		Cost{Void, 1},
	},
	Action: SmallShieldAction,
}

var SturdyShieldAction = Action{
	ActionType: Shield,
	Value:      5,
	Uses:       2,
}

var SturdyShield = Spell{
	Name: "Sturdy Shield",
	Cost: []Cost{
		Cost{Void, 2},
	},
	Action: SturdyShieldAction,
}

var MoltenArmorAction = Action{
	ActionType: ElementalShield,
	BlockElement: []Element{
		Fire,
		Earth,
	},
	Value: 20,
	Uses:  2,
}

var MoltenArmor = Spell{
	Name: "Molten Armor",
	Cost: []Cost{
		Cost{Earth, 2},
		Cost{Fire, 1},
	},
	Action: MoltenArmorAction,
}

var DivineProtectionAction = Action{
	ActionType: TrueShield,
	Uses:       1,
}

var DivineProtection = Spell{
	Name: "Divine Protection",
	Cost: []Cost{
		Cost{Void, 2},
	},
	Action: DivineProtectionAction,
}

////// Draw Cards //////

var ShadowGift = Spell{
	Name: "Shadow Gift",
	Self: true,
	Cost: []Cost{
		Cost{Void, 1},
	},
	Action: Action{
		ActionType: Draw,
		Value:      1,
		Chance:     100,
	},
}

var ShadowWager = Spell{
	Name: "Shadow Wager",
	Self: true,
	Cost: []Cost{
		Cost{Void, 2},
	},
	Action: Action{
		ActionType: Draw,
		Value:      2,
		Chance:     50,
	},
}

///// Debuff Spells /////

var NullStat = Spell{
	Name: "NullStat",
	Cost: []Cost{
		Cost{Void, 1},
	},
	Action: Action{
		ActionType: None,
		IsStatus:   true,
		StatusType: Debuff,
		StatusName: "VOIDED",
		Lifespan:   2,
		Chance:     100,
		Trigger:    TurnStart,
	},
}

var TestPoison = Spell{
	Name: "Poison",
	Cost: []Cost{
		Cost{Earth, 1},
	},
	Action: Action{
		ActionType:    DamagePerTurn,
		IsStatus:      true,
		StatusType:    Debuff,
		StatusName:    "Poisoned",
		DamagePerTurn: 5,
		Lifespan:      2,
		Chance:        100,
		Trigger:       TurnStart,
	},
}

var SpellVirus = Spell{
	Name: "Spell Virus",
	Cost: []Cost{
		Cost{Void, 1},
	},
	Action: Action{
		ActionType:    EffectDamage,
		IsStatus:      true,
		StatusType:    Debuff,
		StatusName:    "Spell Virus",
		DamagePerProc: 5,
		Uses:          2,
		Chance:        100,
		Trigger:       CastingSpell,
	},
}
