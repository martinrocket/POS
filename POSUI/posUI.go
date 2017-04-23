package POSUI

import (
	"fmt"
	"strings"

	"github.com/martinrocket/POS/Version"
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

func BuildBox() {
	boxWidth = 50
	boxHeight = 20
	fmt.Print(rTop, strings.Repeat(hLine, boxWidth), lTop, nL)
	fmt.Print(vLine, strings.Repeat(space, boxWidth-len(version)), version, vLine, nL)
	for i := 0; i < boxHeight; i++ {
		fmt.Print("...")
	}
	fmt.Println()
}
