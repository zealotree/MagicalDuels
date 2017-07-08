package magicalduels

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

type Element struct {
	Name string
}

func (e Element) GetName() string {
	return e.Name
}

func (e Element) IsElement() bool {
	return true
}

func (e Element) IsSpell() bool {
	return false
}

func PrintElements() string {
	return fmt.Sprint("(1) Fire\n(2) Water\n(3) Air\n(4) Earth\n(5) Void\n")
}

func ElementsSelector(text string) (Element, error) {
	if strings.Compare("1", text) == 0 {
		return Fire, nil
	}
	errors := errors.New("This element is not real!")
	return Element{}, errors
}

func (w *Wizard) BigPoolRequest() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("[%s] Add First Element To Your Pool\n", w.Name)
	fmt.Print(PrintElements())
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)
	e, err := ElementsSelector(text)
	if err != nil {
		panic(err)
	}
	w.AddCardToPool(e, 1)
	reader = bufio.NewReader(os.Stdin)
	fmt.Printf("[%s] Add Second Element To Your Pool\n", w.Name)
	fmt.Print(PrintElements())
	text, _ = reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)
	e, err = ElementsSelector(text)
	if err != nil {
		panic(err)
	}
	w.AddCardToPool(e, 1)

}

var Water = Element{"Water"}
var Fire = Element{"Fire"}
var Air = Element{"Air"}
var Earth = Element{"Earth"}
var Void = Element{"Void"}
var Any = Element{"Any"}
