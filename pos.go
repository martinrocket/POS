//Package POS is the run time file for the POS application
package main

import (
	"fmt"

	ui "github.com/martinrocket/POS/POSUI"
	pos "github.com/martinrocket/POS/posData"
)

func main() {
	ui.ClearScreen()
	fmt.Println("POS")
	ui.BuildBox("mainMenu")
	ui.BuildBox("itemMenuItems")
	pos.ListItems()
	fmt.Println("done")

}
