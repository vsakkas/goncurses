// This example demonstrates reading a string from input, rather than a
// single character.

package main

import (
	"log"

	gc "github.com/vsakkas/goncurses"
)

func main() {
	stdscr, err := gc.Init()
	if err != nil {
		log.Fatal("init:", err)
	}
	defer gc.End()

	msg := "Enter a string: "
	row, col := stdscr.MaxYX()
	row, col = (row/2)-1, (col-len(msg))/2
	stdscr.MovePrint(row, col, msg)

	/* GetString will only retrieve the specified number of characters. Any
	attempts by the user to enter more characters will elicit an audible
	beep */
	var str string
	str, err = stdscr.GetString(10)
	if err != nil {
		stdscr.MovePrint(row+1, col, "GetString Error:", err)
	} else {
		stdscr.MovePrintf(row+1, col, "You entered: %s", str)
	}

	stdscr.Refresh()
	stdscr.GetChar()
}
