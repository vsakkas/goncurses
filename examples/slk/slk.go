// Demonstarates the use of the SLK Soft-Keys facilities.

package main

import (
	"log"

	gc "github.com/vsakkas/goncurses"
)

const SOFTKEYS = 1

func main() {
	gc.SlkInit(gc.SLK_323)
	scr, err := gc.Init()
	if err != nil {
		log.Fatal("INIT:", err)
	}
	defer gc.End()

	gc.StartColor()
	gc.InitPair(SOFTKEYS, gc.C_YELLOW, gc.C_BLUE)

	scr.Print("Type any key to exit...")
	labels := [...]string{"ONE", "TWO", "THREE", "FOUR", "FIVE", "SIX", "SEVEN",
		"EIGHT"}
	for i := range labels {
		gc.SlkSet(i+1, labels[i], gc.SLK_CENTER)
	}
	gc.SlkColor(SOFTKEYS)
	gc.SlkNoutRefresh()
	scr.Refresh()
	scr.GetChar()
}
