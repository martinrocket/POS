// Package POSUI builds the menus, interfaces and shows the results
// of the POS application
package POSUI

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"sort"
	"strings"

	"github.com/martinrocket/POS/posData"
	"github.com/martinrocket/POS/posDateTime"
	"github.com/martinrocket/POS/posMenu"
	"github.com/martinrocket/POS/version"
)

// Lines, Spaces and New Lines for building a menu box.
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
var ver = "ver " + version.GetVersion()
var sysDate = posDateTime.PosDate()
var menu = PosMenu.GetMainMenu()

func mainMenuItems() map[int]string {
	myMenu := map[int]string{
		0: " n) " + menu.NewCheck.Name,
		1: " r) " + menu.RecallCheck.Name,
		2: " i) " + menu.ClockIn.Name,
		3: " o) " + menu.ClockOut.Name,
	}

	return myMenu

}

func itemMenuItems() map[int]string {
	m := make(map[int]string)
	var v = posData.ListItems()
	for i := range v {
		fmt.Println(v[i])
	}

	return m

}

// BuildBox funtion is exported for use with main POS.go file.
func BuildBox(m string) {
	boxWidth = 50

	var menuList map[int]string
	switch m {
	case "mainMenu":
		menuList = mainMenuItems()
	case "itemMenuItems":
		menuList = itemMenuItems()
	}

	// Sorting my map above for not random output
	var keys []int
	for k := range menuList {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	// Build the box with the menu selected above
	fmt.Print(rTop, strings.Repeat(hLine, boxWidth), lTop, nL)
	fmt.Print(vLine, space, strings.Repeat(space, boxWidth-len(ver)-2), ver, space, vLine, nL)
	for _, k := range keys {
		fmt.Print(vLine, menuList[k], strings.Repeat(space, boxWidth-len(menuList[k])-1), space, vLine, nL)
	}
	fmt.Print(vLine, space, strings.Repeat(space, boxWidth-len(sysDate)-2), sysDate, space, vLine, nL)
	fmt.Print(rBot, strings.Repeat(hLine, boxWidth), lBot, nL)

}

// ClearScreen exports a command to clear the screen based on the Operating System of the computer
func ClearScreen() {
	if runtime.GOOS == "darwin" || runtime.GOOS == "linux" {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}

	if runtime.GOOS == "windows" {
		cmd := exec.Command("cls") // Clears for Windows OS
		cmd.Stdout = os.Stdout
		cmd.Run()
	}

}
