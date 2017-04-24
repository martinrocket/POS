package POSUI

import (
	"fmt"
	"strings"

	"github.com/martinrocket/POS/Version"
	"github.com/martinrocket/POS/posMenu"
)

const (
	space = " "
	hLine = "─"
	vLine = "│"
	rTop  = "┌"
	lTop  = "┐"
	rBot  = "└"
	lBot  = "┘"
	nL    = "\n"
)

var boxWidth int
var boxHeight int
var version = "ver " + Version.GetVersion()
var menu = PosMenu.GetMainMenu()

// BuildBox funtion is exported for use with main POS.go file.
func BuildBox() {
	boxWidth = 50
	boxHeight = 20
	fmt.Print(rTop, strings.Repeat(hLine, boxWidth), lTop, nL)
	fmt.Print(vLine, space, strings.Repeat(space, boxWidth-len(version)-2), version, space, vLine, nL)
	for i := 0; i < boxHeight; i++ {
		fmt.Print(vLine, menu.NewCheck.Name, strings.Repeat(space, boxWidth-len(menu.NewCheck.Name)-1), space, vLine, nL)
	}
	fmt.Println()
	fmt.Println(len(menu.NewCheck.Name), " ", menu.NewCheck.Name)

}
