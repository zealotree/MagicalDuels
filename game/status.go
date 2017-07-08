package magicalduels

import "fmt"

type StatusType string

var Buff StatusType = "A status that strengtens a Wizard"
var Debuff StatusType = "A status that weakens a Wizard"

// This is used by Spells that inflict status
type Status struct {
	Name       string // Name of Status
	ActionType        // The Status data from spell
	StatusType        // Taken from ActionType.
}

func CheckStatus(target Wizard) (bool, bool) {
	// Returns true if the status exist
	// Status can only be applied once
	buff := false
	debuff := false
	if target.Buff.Name != "" {
		buff = true
	}
	if target.Debuff.Name != "" {
		debuff = true
	}
	return buff, debuff

}

func (w *Wizard) AddStatus(s Spell) {

	if s.Action.StatusType == Debuff {
		w.Debuff = s
	}

}

func ResolveStatus(w *Wizard, s StatusType, p Phase) {
	// Check the lifespan of a
	if s == Debuff {
		if w.Debuff.Name != "" {
			if w.Debuff.Lifespan != 0 && p == TurnStart {
				fmt.Println("Lifespan:", w.Debuff.Lifespan)
				fmt.Println("Decreasing lifespan")
				w.Debuff.Lifespan--
				fmt.Println("Lifespan:", w.Debuff.Lifespan)
				if w.Debuff.Lifespan == 0 {
					w.Debuff = Spell{}
				}
			}

			if w.Debuff.Uses != 0 {
				fmt.Println("Decreasing Uses")
				w.Debuff.Uses--
				if w.Debuff.Uses == 0 {
					o := w.Debuff.Name
					w.Debuff = Spell{}
					fmt.Println(o, "expires")
				}
			}
		}
	}

}

func ProcEffects(w *Wizard, p Phase) {

	var d Library
	fmt.Println("Checking effects", p)
	if w.Debuff.Name != "" {
		if w.Debuff.Trigger == p {
			fmt.Println("Trigger")
			ActionTypeSwitch(w.Debuff, w, &d, true)
			ResolveStatus(w, w.Debuff.StatusType, p)
		}
	}

}

func (s *Spell) InflictStatus(target *Wizard) bool {
	if !s.Action.IsStatus {
		fmt.Println("This spell does not apply a status!")
		return false
	}
	statustype := s.Action.StatusType
	buff, debuff := CheckStatus(*target)
	fmt.Println(debuff)
	if debuff == false && statustype == Debuff {
		target.AddStatus(*s)
		fmt.Println(s.Name, "was inflicted!")
		return true
	} else if debuff == true && statustype == Debuff {
		fmt.Println(s.Name, "failed.", target.Name, "already has a debuff!")
		return false
	}
	if !buff && statustype == Buff {
		fmt.Println(s.Name, "was applied!")
		return true
	}
	return false
}
