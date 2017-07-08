package magicalduels

import "testing"

func AssertEqual(t *testing.T, a interface{}, b interface{}) {
	if a != b {
		t.Fatalf("Values are not equal. Got %v. Expected %v.", a, b)
	}

}

func TestDrawSpell(t *testing.T) {

	d := new(Library)
	d.AddCard(Fireball, 3)
	wiz1 := NewWizard("Merlin", 100)
	wiz1.AddCardToHand(ShadowGift, 1)
	wiz1.AddCardToPool(Void, 1)
	wiz1.Cast(ShadowGift, &wiz1, d)

	AssertEqual(t, len(wiz1.Hand), 1)
}

func TestCardOwnership(t *testing.T) {

	wiz1 := NewWizard("Merlin", 100)
	wiz1.AddCardToHand(MoltenBarrage, 1)
	hasCard, _ := wiz1.HasSpell(MoltenBarrage)
	AssertEqual(t, hasCard, true)
}

func TestMoltenArmorSuccessAgainstMoltenBarrage(t *testing.T) {

	d := new(Library)

	wiz1 := NewWizard("Merlin", 50)
	wiz2 := NewWizard("Gandalf", 50)

	wiz1.AddCardToHand(MoltenBarrage, 1)
	wiz1.AddCardToPool(Earth, 2)
	wiz1.AddCardToPool(Fire, 2)

	wiz2.AddCardToHand(MoltenArmor, 1)
	wiz2.AddCardToPool(Earth, 2)
	wiz2.AddCardToPool(Fire, 1)

	wiz2.Cast(MoltenArmor, &wiz2, d)
	wiz1.Cast(MoltenBarrage, &wiz2, d)

	// This should go through without reduction since Molten Armor cannot block a Fire/Earth attack

	AssertEqual(t, wiz2.Health, 40)

}

func TestMoltenArmorFailAgainstHeatVortex(t *testing.T) {

	d := new(Library)

	wiz1 := NewWizard("Merlin", 50)
	wiz2 := NewWizard("Gandalf", 50)

	wiz1.AddCardToHand(HeatVortex, 1)
	wiz1.AddCardToPool(Air, 1)
	wiz1.AddCardToPool(Fire, 1)

	wiz2.AddCardToHand(MoltenArmor, 1)
	wiz2.AddCardToPool(Earth, 2)
	wiz2.AddCardToPool(Fire, 1)

	wiz2.Cast(MoltenArmor, &wiz2, d)
	wiz1.Cast(HeatVortex, &wiz2, d)

	// This should go through without reduction since Molten Armor cannot block a Fire/Earth attack

	AssertEqual(t, wiz2.Health, 34)

}

func TestTrueShield(t *testing.T) {

	d := new(Library)

	wiz1 := NewWizard("Merlin", 50)
	wiz2 := NewWizard("Gandalf", 50)

	wiz1.AddCardToHand(RockPunch, 1)
	wiz1.AddCardToPool(Earth, 1)

	wiz2.AddCardToHand(DivineProtection, 1)
	wiz2.AddCardToPool(Void, 3)

	wiz2.Cast(DivineProtection, &wiz2, d)
	wiz1.Cast(RockPunch, &wiz2, d)

	AssertEqual(t, wiz2.Health, 50)
}

func TestMultiElementShieldAgainstMonoElementalSpells(t *testing.T) {
	wiz1 := NewWizard("Merlin", 50)
	wiz2 := NewWizard("Gandalf", 50)
	d := new(Library)
	wiz1.AddCardToHand(MoltenArmor, 1)
	wiz1.AddCardToPool(Earth, 2)
	wiz1.AddCardToPool(Fire, 1)

	AssertEqual(t, len(wiz1.Pool), 3)
	wiz1.Cast(MoltenArmor, &wiz1, d)
	AssertEqual(t, wiz1.ActiveShield.Name, "Molten Armor")
	AssertEqual(t, len(wiz1.Pool), 0)

	wiz2.AddCardToHand(Fireball, 1)
	wiz2.AddCardToHand(IceDaggers, 1)
	wiz2.AddCardToPool(Fire, 1)
	wiz2.AddCardToPool(Water, 1)
	wiz2.AddCardToPool(Earth, 1)

	wiz2.Cast(Fireball, &wiz1, d)
	AssertEqual(t, wiz1.Health, 50)

	wiz2.Cast(IceDaggers, &wiz1, d)
	AssertEqual(t, wiz1.Health, 40)

	wiz2.Cast(RockPunch, &wiz1, d)
	AssertEqual(t, wiz1.Health, 40)
}

func TestCostAndPoolAfterCast(t *testing.T) {

	wiz1 := NewWizard("Merlin", 100)

	d := new(Library)
	wiz1.AddCardToHand(SweetWater, 1)
	wiz1.AddCardToHand(SacredRoot, 1)
	wiz1.AddCardToPool(Water, 3)

	AssertEqual(t, len(wiz1.Pool), 3)
	AssertEqual(t, len(wiz1.Hand), 2)

	wiz1.Cast(SweetWater, &wiz1, d)
	AssertEqual(t, len(wiz1.Pool), 2)
	AssertEqual(t, len(wiz1.Hand), 1)

	wiz1.AddCardToPool(Earth, 3)
	AssertEqual(t, len(wiz1.Pool), 5)

	wiz1.Cast(SacredRoot, &wiz1, d)
	AssertEqual(t, len(wiz1.Pool), 2)

}

func TestHealSpell(t *testing.T) {

	wiz1 := NewWizard("Merlin", 100)
	wiz1.Health = 50

	AssertEqual(t, wiz1.Health, 50)
	t.Log(wiz1.Health)
	d := new(Library)
	wiz1.AddCardToHand(SweetWater, 1)
	wiz1.AddCardToPool(Water, 3)

	wiz1.Cast(SweetWater, &wiz1, d)
	AssertEqual(t, wiz1.Health, 58)

}

func TestHealOnMaxHealth(t *testing.T) {

	wiz1 := NewWizard("Merlin", 100)

	d := new(Library)
	wiz1.AddCardToHand(SacredRoot, 1)
	wiz1.AddCardToPool(Earth, 3)

	wiz1.Cast(SacredRoot, &wiz1, d)
	// Heals should not exceed Max Health
	AssertEqual(t, wiz1.Health, wiz1.MaxHealth)

}

func TestFireballSpell(t *testing.T) {

	wiz1 := NewWizard("Merlin", 100)
	wiz2 := NewWizard("Gandalf", 100)

	d := new(Library)
	wiz1.AddCardToHand(Fireball, 1)
	wiz1.AddCardToPool(Fire, 3)

	t.Log(wiz1.Cast(Fireball, &wiz2, d))
	AssertEqual(t, wiz2.Health, 88)

}

func TestActiveShield(t *testing.T) {

	wiz1 := NewWizard("Merlin", 100)

	d := new(Library)
	wiz1.AddCardToHand(SmallShield, 1)
	wiz1.AddCardToPool(Void, 1)

	wiz1.Cast(SmallShield, &wiz1, d)
	AssertEqual(t, wiz1.ActiveShield.Value, 6)

}

func TestActiveShieldDamageReduction(t *testing.T) {

	wiz1 := NewWizard("Merlin", 100)
	wiz2 := NewWizard("Gandalf", 100)

	d := new(Library)
	wiz1.AddCardToHand(SmallShield, 1)
	wiz1.AddCardToPool(Void, 1)
	wiz1.Cast(SmallShield, &wiz1, d)

	wiz2.AddCardToHand(Fireball, 1)
	wiz2.AddCardToPool(Fire, 1)
	wiz2.Cast(Fireball, &wiz1, d)

	// Blocks damage 6

	AssertEqual(t, wiz1.Health, 94)

}

func TestActiveShieldWithMultipleUses(t *testing.T) {

	wiz1 := NewWizard("Merlin", 100)
	wiz2 := NewWizard("Gandalf", 100)

	d := new(Library)
	wiz1.AddCardToHand(SturdyShield, 1)
	wiz1.AddCardToPool(Void, 3)
	wiz1.Cast(SturdyShield, &wiz1, d)

	wiz2.AddCardToHand(Fireball, 3)
	wiz2.AddCardToPool(Fire, 2)

	AssertEqual(t, wiz1.ActiveShield.Uses, 2)
	// Sturdy Shield uses up 1
	wiz2.Cast(Fireball, &wiz1, d)

	AssertEqual(t, wiz1.ActiveShield.Value, 5)
	AssertEqual(t, wiz1.ActiveShield.Uses, 1)

	// Sturdy Shield uses up 1
	// Sturdy Shield breaks
	wiz2.Cast(Fireball, &wiz1, d)

	AssertEqual(t, wiz1.ActiveShield.Uses, 0)
	AssertEqual(t, wiz1.ActiveShield.Name, "")
	AssertEqual(t, wiz1.ActiveShield.Value, 0)

}
