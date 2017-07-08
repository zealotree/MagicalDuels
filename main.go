package main

import (
	"fmt"
	md "magicalduels/game"
)

func main() {
	// Create two players
	fmt.Println("********** BEGIN ************")
	d := new(md.Library)
	wiz1 := md.NewWizard("Merlin", 100)
	wiz2 := md.NewWizard("Gandalf", 100)

	g := md.NewGame(wiz1, wiz2, *d)

	g.Init()
	g.Run(true)

	fmt.Println("********** END ************")
}
