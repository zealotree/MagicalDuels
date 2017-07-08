package magicalduels

import "testing"

var Merlo = NewWizard("JoJo", 100)
var Gandal = NewWizard("Gandal", 100)
var deck = Library{}

func TestTriggerAStatusbyCasting(t *testing.T) {
	var m = Merlo
	var c = Gandal
	d.AddCard(SpellVirus, 2)
	m.AddCardToHand(SpellVirus, 1)
	m.AddCardToHand(Fireball, 1)
	m.AddCardToPool(Void, 20)
	m.AddCardToPool(Fire, 20)
	m.Cast(SpellVirus, &m, &d)

	AssertEqual(t, m.Debuff.Name, SpellVirus.Name)
	m.Cast(Fireball, &c, &d)
	AssertEqual(t, m.Health, 95)
}

func TestInflictAPoison(t *testing.T) {
	var m = Merlo
	d.AddCard(TestPoison, 2)
	m.AddCardToHand(TestPoison, 1)
	m.AddCardToPool(Earth, 20)
	m.Cast(TestPoison, &m, &d)
	t.Log(m.Health)
	AssertEqual(t, m.Debuff.Lifespan, 2)
	AssertEqual(t, m.Debuff.Action.StatusName, TestPoison.Action.StatusName)
	ProcEffects(&m, TurnStart)
	AssertEqual(t, m.Debuff.Lifespan, 1)
	AssertEqual(t, m.Health, 95)
	ProcEffects(&m, TurnStart)
	AssertEqual(t, m.Health, 90)
	AssertEqual(t, m.Debuff.Lifespan, 0)
	AssertEqual(t, m.Debuff.Name, "") // Should be removed now
}

func TestInflictAStatus(t *testing.T) {
	var m = Merlo
	d.AddCard(NullStat, 2)
	m.AddCardToHand(NullStat, 1)
	m.AddCardToPool(Void, 20)
	m.Cast(NullStat, &m, &d)

	AssertEqual(t, m.Debuff.Action.StatusName, NullStat.Action.StatusName)

	t.Log(NullStat.InflictStatus(&m))
	AssertEqual(t, NullStat.InflictStatus(&m), false)
}
