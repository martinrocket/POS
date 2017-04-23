//Package POS is the run time file for the POS application
package main

import (
	"fmt"

	ui "github.com/martinrocket/POS/POSUI"
)

func main() {
	fmt.Println("POS")
	ui.BuildBox()
	fmt.Println("done")

}
