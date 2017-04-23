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
var menu posMenu.PosFuncs

// BuildBox funtion is exported for use with main POS.go file.
func BuildBox() {
	boxWidth = 50
	boxHeight = 20
	fmt.Print(rTop, strings.Repeat(hLine, boxWidth), lTop, nL)
	fmt.Print(vLine, strings.Repeat(space, boxWidth-len(version)), version, vLine, nL)
	for i := 0; i < boxHeight; i++ {
		fmt.Print("...")
	}
	fmt.Println()

	menu.ClockIn = "yes"
	fmt.Println(menu)
}
