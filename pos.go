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

	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("fizz buzz")
		} else if i%3 == 0 {
			fmt.Println("fizz")
		} else if i%5 == 0 {
			fmt.Println("buzz")
		} else {
			fmt.Println(i)
		}

	}

}
