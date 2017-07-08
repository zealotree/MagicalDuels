package magicalduels

import "testing"

func TestNumOfElementInPool(t *testing.T) {

	wiz1 := NewWizard("Merlin", 100)
	wiz1.AddCardToPool(Fire, 30)
	wiz1.AddCardToPool(Water, 30)
	wiz1.AddCardToPool(Air, 30)
	wiz1.AddCardToPool(Earth, 30)
	wiz1.AddCardToPool(Void, 30)

	m := GetElementCountFromCards(wiz1.Pool, []Element{Fire})
	AssertEqual(t, m[Fire], 30)
	AssertEqual(t, m[Water], 30)
	AssertEqual(t, m[Air], 30)
	AssertEqual(t, m[Earth], 30)
	AssertEqual(t, m[Void], 30)

}
